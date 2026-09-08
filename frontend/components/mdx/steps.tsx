import { type ReactNode, Children } from 'react';

export function Steps({ children }: { children: ReactNode }) {
  const items = Children.toArray(children);
  return (
    <ol className="my-4 space-y-4">
      {items.map((item, i) => (
        <li key={i} className="flex gap-3">
          <span className="flex h-6 w-6 shrink-0 items-center justify-center rounded-full bg-primary text-xs text-primary-foreground">
            {i + 1}
          </span>
          <div className="flex-1">{item}</div>
        </li>
      ))}
    </ol>
  );
}

export function Definition({ term, children }: { term: string; children: ReactNode }) {
  return (
    <div className="my-4 rounded-lg border-l-4 border-primary bg-muted/50 p-4">
      <dt className="font-semibold">{term}</dt>
      <dd className="mt-1 text-sm text-muted-foreground">{children}</dd>
    </div>
  );
}
