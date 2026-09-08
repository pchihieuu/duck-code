import { CourseCard } from '@/components/course/course-card';
import type { Course } from '@/types/course';

export const metadata = { title: 'Khoá học' };

async function getCourses(): Promise<Course[]> {
  // TODO: thay bằng coursesApi.list()
  return [
    {
      id: '1',
      slug: 'python-basic',
      title: 'Python Basic',
      description: 'Nhập môn Python: biến, vòng lặp, hàm, cấu trúc dữ liệu.',
      level: 'beginner',
      tags: ['python'],
      chapters: [],
      totalLessons: 20,
      totalExercises: 15,
      estimatedHours: 12,
    },
  ];
}

export default async function CoursesPage() {
  const courses = await getCourses();
  return (
    <div className="container py-10">
      <h1 className="mb-6 text-2xl font-semibold">Khoá học</h1>
      <div className="grid gap-4 sm:grid-cols-2 lg:grid-cols-3">
        {courses.map((course) => (
          <CourseCard key={course.id} course={course} />
        ))}
      </div>
    </div>
  );
}
