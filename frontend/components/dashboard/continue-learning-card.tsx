import Link from 'next/link';
import { Play } from 'lucide-react';
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';
import { Progress } from '@/components/ui/progress';

export function ContinueLearningCard({
  lessonTitle,
  courseTitle,
  progressPercent,
  href,
}: {
  lessonTitle: string;
  courseTitle: string;
  progressPercent: number;
  href: string;
}) {
  return (
    <Card>
      <CardHeader className="pb-3">
        <CardTitle className="flex items-center gap-2 text-base">
          <Play className="h-4 w-4 fill-primary text-primary" />
          Tiếp tục học
        </CardTitle>
      </CardHeader>
      <CardContent>
        <Link href={href} className="flex items-center gap-4">
          <span className="flex h-16 w-16 shrink-0 items-center justify-center rounded-xl bg-accent text-2xl">
            🐼
          </span>
          <div className="min-w-0 flex-1">
            <p className="truncate font-semibold">{lessonTitle}</p>
            <p className="truncate text-sm text-muted-foreground">{courseTitle}</p>
            <div className="mt-2 flex items-center gap-2">
              <Progress value={progressPercent} className="h-1.5" />
              <span className="shrink-0 text-xs text-muted-foreground">{progressPercent}%</span>
            </div>
          </div>
        </Link>
      </CardContent>
    </Card>
  );
}
