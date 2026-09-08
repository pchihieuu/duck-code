export interface Quiz {
  id: string;
  slug: string;
  title: string;
  questions: QuizQuestion[];
  timeLimitSeconds?: number;
}

export interface QuizQuestion {
  id: string;
  prompt: string;
  options: { id: string; label: string }[];
  multiple?: boolean;
}

export interface QuizResult {
  quizId: string;
  score: number;
  total: number;
  xpEarned: number;
}
