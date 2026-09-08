import { ExercisePanel } from '@/components/exercise/exercise-panel';
import type { Exercise } from '@/types/exercise';

interface Props {
  params: { exerciseId: string };
}

async function getExercise(id: string): Promise<Exercise> {
  // TODO: exercisesApi.byId(id)
  return {
    id,
    slug: 'sum-two-numbers',
    title: 'Tính tổng hai số',
    description: 'Viết hàm nhận vào hai số nguyên và trả về tổng của chúng.',
    starterCode: 'def sum_two(a, b):\n    # TODO\n    pass\n',
    language: 'python',
    difficulty: 'easy',
    visibleTestCases: [{ id: '1', input: '1, 2', expectedOutput: '3' }],
  };
}

export default async function ExercisePage({ params }: Props) {
  const exercise = await getExercise(params.exerciseId);
  return (
    <div className="container py-10">
      <ExercisePanel exercise={exercise} />
    </div>
  );
}
