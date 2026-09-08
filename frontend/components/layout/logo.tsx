import Link from 'next/link';
import { Code2 } from 'lucide-react';

export function Logo() {
  return (
    <Link href="/" className="flex items-center gap-3">
      <span className="flex h-11 w-11 shrink-0 items-center justify-center rounded-full bg-primary text-primary-foreground shadow-sm">
        <Code2 className="h-6 w-6" />
      </span>
      <span className="flex flex-col leading-tight">
        <span className="text-lg font-bold tracking-tight">Learning Platform</span>
        <span className="text-xs text-muted-foreground">Học lập trình dễ dàng</span>
      </span>
    </Link>
  );
}
