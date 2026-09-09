// Package mdxvalidate does a best-effort structural check of MDX content
// before it's persisted — catching the mistakes admins actually make
// (unclosed code fence, stray "{", unclosed JSX-like tag) without shelling
// out to a real MDX compiler.
//
// This is deliberately NOT a real MDX/JSX parser. A real one needs a JS
// runtime (@mdx-js/mdx) — running that as a Node subprocess from a Go
// service is a legitimate option, but adds an operational dependency that
// isn't worth it here: lesson content is admin-only input (already a
// trusted actor via middleware.RequireRole("admin")), so the goal is
// "catch obvious typos before they 500 the lesson page for students," not
// "defend against adversarial input." If that trust assumption ever
// changes (e.g. instructors who aren't full admins start authoring
// content), revisit this and add a real compile step.
//
// Plain prose/Markdown with no JSX tags or "{...}" expressions always
// passes trivially — this only flags structural mistakes, never rejects
// content just for being plain text.
//
// Known false positives (documented, not silently wrong): a comparison
// like "a<b" written outside a fenced code block (e.g. an admin pastes a
// bare line of Python without wrapping it in ``` fences) can be misread as
// an opening tag <b>. Mitigation is procedural, not code: always wrap code
// samples in fenced code blocks — which is already the recommended way to
// show example code in a lesson.
package mdxvalidate

import (
	"bufio"
	"fmt"
	"regexp"
	"strings"
)

var (
	// Matches both opening (<Foo ...>) and closing (</Foo>) tags in one
	// pass, so we can walk them in document order for correct nesting.
	tagRe     = regexp.MustCompile(`</?[A-Za-z][A-Za-z0-9]*\b[^<>]*>`)
	tagNameRe = regexp.MustCompile(`^</?([A-Za-z][A-Za-z0-9]*)`)

	// Inline code spans (`x`) are literal text in MDX too, same as fenced
	// blocks — strip them before checking braces/tags.
	inlineCodeRe = regexp.MustCompile("`[^`\n]*`")
)

// Validate returns a descriptive error if content has a structural MDX
// mistake, or nil if it looks fine (including: content has no MDX/JSX
// syntax at all, i.e. plain text/Markdown).
func Validate(content string) error {
	stripped, err := stripFencedCode(content)
	if err != nil {
		return err
	}
	stripped = inlineCodeRe.ReplaceAllString(stripped, "")

	if err := checkBraceBalance(stripped); err != nil {
		return err
	}
	if err := checkTagBalance(stripped); err != nil {
		return err
	}
	return nil
}

// stripFencedCode removes the *body* of ``` / ~~~ fenced blocks (keeping
// everything outside them) so example code (which legitimately contains
// unbalanced-looking braces, e.g. Python dict literals split across lines,
// or "<" used as a comparison operator) never triggers a false positive.
// Returns an error if a fence is opened but never closed — that itself is
// the single most common admin mistake worth catching.
func stripFencedCode(content string) (string, error) {
	var out strings.Builder
	scanner := bufio.NewScanner(strings.NewReader(content))
	scanner.Buffer(make([]byte, 0, 64*1024), 10*1024*1024) // allow long lines

	inFence := false
	fenceMarker := ""
	fenceStartLine := 0
	lineNo := 0

	for scanner.Scan() {
		lineNo++
		line := scanner.Text()
		trimmed := strings.TrimSpace(line)

		switch {
		case !inFence && (strings.HasPrefix(trimmed, "```") || strings.HasPrefix(trimmed, "~~~")):
			inFence = true
			fenceMarker = trimmed[:3]
			fenceStartLine = lineNo
			continue // drop the fence-open line itself
		case inFence && strings.HasPrefix(trimmed, fenceMarker):
			inFence = false
			continue // drop the fence-close line itself
		case inFence:
			continue // drop code sample content entirely
		default:
			out.WriteString(line)
			out.WriteByte('\n')
		}
	}
	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("mdxvalidate: failed reading content: %w", err)
	}
	if inFence {
		return "", fmt.Errorf("mdxvalidate: unclosed code fence starting at line %d — every ``` or ~~~ needs a matching closing fence", fenceStartLine)
	}
	return out.String(), nil
}

// checkBraceBalance validates "{...}" expression syntax (the MDX-specific
// part on top of plain Markdown). Plain text with no braces at all is
// trivially balanced (depth stays 0) and always passes.
func checkBraceBalance(s string) error {
	depth := 0
	line := 1
	for _, r := range s {
		switch r {
		case '\n':
			line++
		case '{':
			depth++
		case '}':
			depth--
			if depth < 0 {
				return fmt.Errorf("mdxvalidate: unmatched '}' near line %d — a closing brace with no matching '{' before it", line)
			}
		}
	}
	if depth > 0 {
		return fmt.Errorf("mdxvalidate: %d unclosed '{' expression(s) — every '{' needs a matching '}'", depth)
	}
	return nil
}

// checkTagBalance validates JSX-like tag nesting. Plain text with no "<...>"
// tags at all is trivially balanced (empty stack) and always passes.
func checkTagBalance(s string) error {
	var stack []string

	for _, tok := range tagRe.FindAllString(s, -1) {
		closing := strings.HasPrefix(tok, "</")
		selfClosing := strings.HasSuffix(tok, "/>")
		name := tagName(tok)

		switch {
		case selfClosing:
			continue // <Foo /> needs no matching close, nothing to push/pop
		case closing:
			if len(stack) == 0 {
				return fmt.Errorf("mdxvalidate: closing tag </%s> has no matching opening tag", name)
			}
			top := stack[len(stack)-1]
			if top != name {
				return fmt.Errorf("mdxvalidate: closing tag </%s> doesn't match currently open tag <%s>", name, top)
			}
			stack = stack[:len(stack)-1]
		default:
			stack = append(stack, name)
		}
	}

	if len(stack) > 0 {
		return fmt.Errorf("mdxvalidate: %d unclosed tag(s), innermost still open is <%s>", len(stack), stack[len(stack)-1])
	}
	return nil
}

func tagName(tag string) string {
	m := tagNameRe.FindStringSubmatch(tag)
	if len(m) < 2 {
		return tag
	}
	return m[1]
}