export interface Exercise {
  id: string;
  slug: string;
  title: string;
  description: string;
  starterCode: string;
  language: 'python' | 'javascript' | 'typescript';
  difficulty: 'easy' | 'medium' | 'hard';
  visibleTestCases: TestCase[];
}

export interface TestCase {
  id: string;
  input: string;
  expectedOutput: string;
}

export interface RunResult {
  stdout: string;
  stderr: string;
  success: boolean;
}

export interface SubmissionResult {
  id: string;
  status: 'pending' | 'running' | 'passed' | 'failed' | 'error';
  passedCount: number;
  totalCount: number;
  xpEarned?: number;
}
