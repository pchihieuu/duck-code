import Link from 'next/link';
import { notFound } from 'next/navigation';
import { ROUTES } from '@/lib/constants';

interface Props {
  params: { courseSlug: string };
}

async function getCourse(slug: string) {
  // TODO: coursesApi.bySlug(slug)
  if (slug !== 'python-basic') return null;
  return {
    title: 'Python Basic',
    description: 'Nhập môn Python cho người mới bắt đầu.',
    chapters: [
      {
        title: 'Chương 1: Làm quen với Python',
        lessons: [{ slug: 'gioi-thieu', title: 'Giới thiệu Python' }],
      },
    ],
  };
}

export default async function CourseDetailPage({ params }: Props) {
  const course = await getCourse(params.courseSlug);
  if (!course) notFound();

  return (
    <div className="container py-10">
      <h1 className="text-2xl font-semibold">{course.title}</h1>
      <p className="mt-2 text-muted-foreground">{course.description}</p>

      <div className="mt-8 space-y-6">
        {course.chapters.map((chapter) => (
          <div key={chapter.title}>
            <h2 className="mb-2 font-medium">{chapter.title}</h2>
            <ul className="space-y-1">
              {chapter.lessons.map((lesson) => (
                <li key={lesson.slug}>
                  <Link
                    className="text-sm text-primary hover:underline"
                    href={ROUTES.lesson(params.courseSlug, lesson.slug)}
                  >
                    {lesson.title}
                  </Link>
                </li>
              ))}
            </ul>
          </div>
        ))}
      </div>
    </div>
  );
}
