export interface Lesson {
  id: string;
  slug: string;
  courseSlug: string;
  chapterSlug: string;
  title: string;
  content: string; // compiled MDX
  order: number;
  nextLessonSlug?: string;
  prevLessonSlug?: string;
}
