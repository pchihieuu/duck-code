import { apiClient } from './client';
import type { Course } from '@/types/course';

export const coursesApi = {
  list: () => apiClient.get<Course[]>('/courses'),
  bySlug: (slug: string) => apiClient.get<Course>(`/courses/${slug}`),
};
