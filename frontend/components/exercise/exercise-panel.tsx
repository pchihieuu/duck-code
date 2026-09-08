'use client';

import * as React from 'react';
import { Button } from '@/components/ui/button';
import { CodeEditor } from './code-editor';
import { exercisesApi } from '@/lib/api/exercises';
import type { Exercise, RunResult, SubmissionResult } from '@/types/exercise';

// Run: chạy trong browser runner (nhanh, test case hiển thị).
// Submit: gửi backend -> queue -> code runner -> kết quả đầy đủ (ẩn test case).
export function ExercisePanel({ exercise }: { exercise: Exercise }) {
  const [code, setCode] = React.useState(exercise.starterCode);
  const [runResult, setRunResult] = React.useState<RunResult | null>(null);
  const [submission, setSubmission] = React.useState<SubmissionResult | null>(null);
  const [isRunning, setIsRunning] = React.useState(false);
  const [isSubmitting, setIsSubmitting] = React.useState(false);

  async function handleRun() {
    setIsRunning(true);
    try {
      const result = await exercisesApi.run(exercise.id, code);
      setRunResult(result);
    } finally {
      setIsRunning(false);
    }
  }

  async function handleSubmit() {
    setIsSubmitting(true);
    try {
      const result = await exercisesApi.submit(exercise.id, code);
      setSubmission(result);
    } finally {
      setIsSubmitting(false);
    }
  }

  return (
    <div className="grid gap-4 lg:grid-cols-2">
      <div>
        <h1 className="mb-2 text-xl font-semibold">{exercise.title}</h1>
        <p className="text-sm text-muted-foreground">{exercise.description}</p>
      </div>
      <div className="space-y-3">
        <CodeEditor value={code} language={exercise.language} onChange={setCode} />
        <div className="flex gap-2">
          <Button variant="outline" onClick={handleRun} disabled={isRunning}>
            {isRunning ? 'Đang chạy…' : 'Run'}
          </Button>
          <Button onClick={handleSubmit} disabled={isSubmitting}>
            {isSubmitting ? 'Đang nộp…' : 'Submit'}
          </Button>
        </div>
        {runResult && (
          <pre className="rounded-md bg-zinc-900 p-3 text-xs text-zinc-100">
            {runResult.stdout || runResult.stderr}
          </pre>
        )}
        {submission && (
          <p className="text-sm">
            Kết quả: {submission.passedCount}/{submission.totalCount} test case
            {submission.xpEarned ? ` · +${submission.xpEarned} XP` : ''}
          </p>
        )}
      </div>
    </div>
  );
}
