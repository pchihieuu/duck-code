import { apiClient } from './client';
import type { Exercise, RunResult, SubmissionResult } from '@/types/exercise';

export const exercisesApi = {
  byId: (id: string) => apiClient.get<Exercise>(`/exercises/${id}`),
  run: (id: string, code: string) =>
    apiClient.post<RunResult>(`/exercises/${id}/run`, { code }),
  submit: (id: string, code: string) =>
    apiClient.post<SubmissionResult>(`/exercises/${id}/submit`, { code }),
};
