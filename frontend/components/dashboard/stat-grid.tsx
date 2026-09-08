import { BookOpen, Sparkles, Zap, Medal } from 'lucide-react';
import { StatCard } from './stat-card';

export function StatGrid({
  lessonsDone,
  wordsLearned,
  weeklyXp,
  badges,
}: {
  lessonsDone: number;
  wordsLearned: number;
  weeklyXp: number;
  badges: number;
}) {
  return (
    <div className="grid grid-cols-2 gap-3">
      <StatCard icon={BookOpen} label="Bài học đã học" value={lessonsDone} />
      <StatCard icon={Sparkles} label="Từ vựng đã học" value={wordsLearned} />
      <StatCard icon={Zap} label="XP tuần" value={weeklyXp} />
      <StatCard icon={Medal} label="Huy hiệu" value={badges} />
    </div>
  );
}
