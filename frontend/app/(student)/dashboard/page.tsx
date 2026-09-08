import { TrendingUp } from 'lucide-react';
import { HeroBanner } from '@/components/dashboard/hero-banner';
import { StreakCard } from '@/components/dashboard/streak-card';
import { ContinueLearningCard } from '@/components/dashboard/continue-learning-card';
import { StatGrid } from '@/components/dashboard/stat-grid';
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';
import { ROUTES } from '@/lib/constants';

// Server Component: fetch dữ liệu ban đầu trên server.
export default async function DashboardPage() {
  // TODO: thay bằng gọi API thật (getServerSession + apiClient)
  const streak = { days: 0, activeDayIndex: 1, best: 0 };
  const stats = { lessonsDone: 0, wordsLearned: 0, weeklyXp: 0, badges: 0 };
  const continueLesson = {
    title: 'Bài 1: Làm quen với Python',
    courseTitle: 'Python Basic · Chương 1',
    progressPercent: 0,
    href: ROUTES.lesson('python-basic', 'gioi-thieu'),
  };

  return (
    <div className="grid gap-6 lg:grid-cols-3">
      <div className="space-y-6 lg:col-span-2">
        <HeroBanner />

        <div className="grid gap-4 sm:grid-cols-2">
          <ContinueLearningCard
            lessonTitle={continueLesson.title}
            courseTitle={continueLesson.courseTitle}
            progressPercent={continueLesson.progressPercent}
            href={continueLesson.href}
          />

          <Card>
            <CardHeader className="pb-3">
              <CardTitle className="flex items-center gap-2 text-base">
                <TrendingUp className="h-4 w-4 text-primary" />
                Thống kê học tập
              </CardTitle>
            </CardHeader>
            <CardContent>
              <StatGrid {...stats} />
            </CardContent>
          </Card>
        </div>
      </div>

      <div className="space-y-6">
        <StreakCard streakDays={streak.days} activeDayIndex={streak.activeDayIndex} bestStreak={streak.best} />
      </div>
    </div>
  );
}
