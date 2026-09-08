import Link from 'next/link';
import { ArrowRight } from 'lucide-react';
import { Button } from '@/components/ui/button';
import { ROUTES } from '@/lib/constants';

export function HeroBanner() {
  return (
    <div className="relative overflow-hidden rounded-2xl border bg-gradient-to-br from-accent via-secondary to-background p-8 md:p-10">
      <div className="max-w-md">
        <h1 className="text-3xl font-extrabold leading-tight tracking-tight md:text-4xl">
          Nâng cao mỗi ngày
          <br />
          <span className="text-primary">Tiến bộ không ngừng!</span>
        </h1>
        <p className="mt-4 text-sm text-muted-foreground md:text-base">
          Luyện code mỗi ngày giúp bạn mở ra những cơ hội nghề nghiệp mới.
        </p>
        <Button asChild size="lg" className="mt-6 rounded-full">
          <Link href={ROUTES.courses}>
            Bắt đầu học ngay
            <ArrowRight className="ml-2 h-4 w-4" />
          </Link>
        </Button>
      </div>
    </div>
  );
}
