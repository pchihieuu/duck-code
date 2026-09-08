export interface CourseProgress {
  courseId: string;
  completedLessons: number;
  totalLessons: number;
  percent: number;
  lastLessonSlug?: string;
}

export interface Achievement {
  id: string;
  title: string;
  description: string;
  icon: string;
  unlockedAt?: string;
}
