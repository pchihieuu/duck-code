import Link from 'next/link';
import { Button } from '@/components/ui/button';
import { ROUTES } from '@/lib/constants';

export function ExerciseEmbed({ exerciseId, title }: { exerciseId: string; title: string }) {
  return (
    <div className="my-4 flex items-center justify-between rounded-lg border p-4">
      <span className="font-medium">{title}</span>
      <Button asChild size="sm">
        <Link href={ROUTES.exercise(exerciseId)}>Làm bài tập</Link>
      </Button>
    </div>
  );
}

export function QuizEmbed({ quizId, title }: { quizId: string; title: string }) {
  return (
    <div className="my-4 flex items-center justify-between rounded-lg border p-4">
      <span className="font-medium">{title}</span>
      <Button asChild size="sm" variant="outline">
        <Link href={ROUTES.quiz(quizId)}>Làm quiz</Link>
      </Button>
    </div>
  );
}
