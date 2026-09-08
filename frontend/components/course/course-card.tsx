import Link from 'next/link';
import { Card, CardContent, CardHeader, CardTitle, CardDescription } from '@/components/ui/card';
import { ROUTES } from '@/lib/constants';
import type { Course } from '@/types/course';

export function CourseCard({ course }: { course: Course }) {
  return (
    <Link href={ROUTES.course(course.slug)}>
      <Card className="h-full transition-shadow hover:shadow-md">
        <CardHeader>
          <CardTitle>{course.title}</CardTitle>
          <CardDescription>{course.description}</CardDescription>
        </CardHeader>
        <CardContent className="flex flex-wrap gap-2 text-xs text-muted-foreground">
          <span>{course.totalLessons} bài học</span>
          <span>·</span>
          <span>{course.totalExercises} bài tập</span>
          <span>·</span>
          <span>{course.level}</span>
        </CardContent>
      </Card>
    </Link>
  );
}
