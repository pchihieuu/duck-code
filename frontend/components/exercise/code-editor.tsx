'use client';

import dynamic from 'next/dynamic';

// Lazy-load Monaco: nặng, chỉ tải khi vào trang exercise.
const Monaco = dynamic(() => import('@monaco-editor/react'), {
  ssr: false,
  loading: () => (
    <div className="flex h-full items-center justify-center text-sm text-muted-foreground">
      Đang tải trình soạn thảo…
    </div>
  ),
});

export function CodeEditor({
  value,
  language,
  onChange,
}: {
  value: string;
  language: string;
  onChange: (value: string) => void;
}) {
  return (
    <div className="h-[420px] overflow-hidden rounded-lg border">
      <Monaco
        height="100%"
        language={language}
        value={value}
        theme="vs-dark"
        onChange={(v) => onChange(v ?? '')}
        options={{ minimap: { enabled: false }, fontSize: 14 }}
      />
    </div>
  );
}
