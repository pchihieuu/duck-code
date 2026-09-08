import Link from 'next/link';
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';
import { Progress } from '@/components/ui/progress';
import { ROUTES } from '@/lib/constants';
import type { CourseProgress } from '@/types/progress';

export function CourseProgressCard({
  title,
  slug,
  progress,
}: {
  title: string;
  slug: string;
  progress: CourseProgress;
}) {
  return (
    <Link href={ROUTES.course(slug)}>
      <Card className="transition-shadow hover:shadow-md">
        <CardHeader>
          <CardTitle className="text-base">{title}</CardTitle>
        </CardHeader>
        <CardContent>
          <Progress value={progress.percent} />
          <p className="mt-2 text-xs text-muted-foreground">
            {progress.completedLessons}/{progress.totalLessons} bài học
          </p>
        </CardContent>
      </Card>
    </Link>
  );
}
