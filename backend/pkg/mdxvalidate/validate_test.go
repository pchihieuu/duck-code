package mdxvalidate

import "testing"

func TestValidate(t *testing.T) {
	cases := []struct {
		name    string
		content string
		wantErr bool
	}{
		{
			name:    "empty content",
			content: "",
			wantErr: false,
		},
		{
			name:    "plain prose, no MDX syntax at all",
			content: "Đây là một bài học bình thường.\n\nCó vài đoạn văn, không có gì đặc biệt.",
			wantErr: false,
		},
		{
			name:    "plain markdown headings/lists/links",
			content: "# Tiêu đề\n\n- item 1\n- item 2\n\n[link](https://example.com)",
			wantErr: false,
		},
		{
			name:    "fenced python code with braces and comparisons — must not false-positive",
			content: "Ví dụ:\n\n```python\nd = {\"a\": 1}\nif a < b and c > d:\n    print(d)\n```\n\nGiải thích ở đây.",
			wantErr: false,
		},
		{
			name:    "inline code span with angle-bracket-ish content",
			content: "Dùng `a < b` để so sánh.",
			wantErr: false,
		},
		{
			name:    "valid self-closing JSX tag",
			content: "Xem hình: <Image src=\"x.png\" />",
			wantErr: false,
		},
		{
			name:    "valid matched JSX tag",
			content: "<Callout type=\"info\">Chú ý điều này.</Callout>",
			wantErr: false,
		},
		{
			name:    "nested matched JSX tags",
			content: "<Callout><Bold>chú ý</Bold> điều này.</Callout>",
			wantErr: false,
		},
		{
			name:    "valid balanced expression braces",
			content: "Kết quả là {1 + 1} thôi.",
			wantErr: false,
		},
		{
			name:    "unclosed code fence",
			content: "```python\nprint(\"hi\")\n",
			wantErr: true,
		},
		{
			name:    "unmatched opening brace",
			content: "Kết quả là {1 + 1 thôi.",
			wantErr: true,
		},
		{
			name:    "unmatched closing brace",
			content: "Kết quả là 1 + 1} thôi.",
			wantErr: true,
		},
		{
			name:    "unclosed JSX tag",
			content: "<Callout>Chú ý điều này.",
			wantErr: true,
		},
		{
			name:    "mismatched JSX tag",
			content: "<Callout>Chú ý</Warning>",
			wantErr: true,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			err := Validate(tc.content)
			if tc.wantErr && err == nil {
				t.Errorf("expected an error, got nil")
			}
			if !tc.wantErr && err != nil {
				t.Errorf("expected no error, got: %v", err)
			}
		})
	}
}