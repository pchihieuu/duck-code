export function CodeBlock({
  children,
  language,
}: {
  children: string;
  language?: string;
}) {
  return (
    <pre className="my-4 overflow-x-auto rounded-lg bg-zinc-900 p-4 text-sm text-zinc-100">
      <code data-language={language}>{children}</code>
    </pre>
  );
}

export const CodeExample = CodeBlock;
