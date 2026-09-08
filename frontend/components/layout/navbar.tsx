import Link from 'next/link';
import { Facebook, Type, Sun } from 'lucide-react';
import { Button } from '@/components/ui/button';
import { Logo } from './logo';
import { ROUTES } from '@/lib/constants';

function IconButton({
  children,
  label,
  disabled,
}: {
  children: React.ReactNode;
  label: string;
  disabled?: boolean;
}) {
  return (
    <button
      type="button"
      aria-label={label}
      disabled={disabled}
      className="hidden h-9 w-9 items-center justify-center rounded-full text-muted-foreground transition-colors hover:bg-accent hover:text-accent-foreground disabled:cursor-not-allowed disabled:text-muted-foreground/40 disabled:hover:bg-transparent sm:flex"
    >
      {children}
    </button>
  );
}

export function Navbar() {
  return (
    <header className="sticky top-0 z-40 flex h-[72px] items-center justify-between border-b bg-background px-4 md:px-6">
      <div className="flex items-center gap-4">
        <Logo />
        <span className="hidden h-8 w-px bg-border md:block" />
      </div>

      <div className="flex items-center gap-1.5 md:gap-3">
        <a
          href="https://facebook.com"
          target="_blank"
          rel="noreferrer"
          className="hidden h-9 w-9 items-center justify-center rounded-full text-muted-foreground hover:bg-accent hover:text-accent-foreground sm:flex"
          aria-label="Facebook"
        >
          <Facebook className="h-5 w-5" />
        </a>
        {/* Toggle cỡ chữ — placeholder, sẽ bật khi có tính năng accessibility */}
        <IconButton label="Cỡ chữ" disabled>
          <Type className="h-5 w-5" />
        </IconButton>
        <button
          className="hidden h-8 w-10 items-center justify-center rounded-md border text-base leading-none sm:flex"
          aria-label="Ngôn ngữ: Tiếng Việt"
          title="Tiếng Việt"
        >
          🇻🇳
        </button>
        <IconButton label="Đổi giao diện sáng/tối">
          <Sun className="h-5 w-5" />
        </IconButton>

        <Button variant="outline" asChild className="rounded-full">
          <Link href={ROUTES.login}>Đăng nhập</Link>
        </Button>
        <Button asChild className="rounded-full">
          <Link href={ROUTES.register}>Đăng ký</Link>
        </Button>
      </div>
    </header>
  );
}
