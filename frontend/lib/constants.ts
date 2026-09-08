export const ROUTES = {
  home: '/',
  courses: '/courses',
  course: (slug: string) => `/courses/${slug}`,
  lesson: (courseSlug: string, lessonSlug: string) =>
    `/courses/${courseSlug}/${lessonSlug}`,
  exercise: (id: string) => `/exercises/${id}`,
  quiz: (id: string) => `/quizzes/${id}`,
  dashboard: '/dashboard',
  progress: '/progress',
  achievements: '/achievements',
  leaderboard: '/leaderboard',
  profile: '/profile',
  settings: '/settings',
  login: '/login',
  register: '/register',
  admin: '/admin',
} as const;
