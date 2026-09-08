export interface Course {
  id: string;
  slug: string;
  title: string;
  description: string;
  coverImage?: string;
  level: 'beginner' | 'intermediate' | 'advanced';
  tags: string[];
  chapters: Chapter[];
  totalLessons: number;
  totalExercises: number;
  estimatedHours: number;
}

export interface Chapter {
  id: string;
  slug: string;
  title: string;
  order: number;
  lessons: LessonSummary[];
}

export interface LessonSummary {
  id: string;
  slug: string;
  title: string;
  order: number;
  type: 'lesson' | 'exercise' | 'quiz';
  isCompleted?: boolean;
}
