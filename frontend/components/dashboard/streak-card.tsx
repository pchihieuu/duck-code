import { Flame, ChevronRight, Trophy } from 'lucide-react';
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';
import { cn } from '@/lib/utils';

const DAYS = ['T2', 'T3', 'T4', 'T5', 'T6', 'T7', 'CN'];

export function StreakCard({
  streakDays,
  activeDayIndex,
  bestStreak,
}: {
  streakDays: number;
  activeDayIndex: number; // 0 = T2 ... 6 = CN
  bestStreak: number;
}) {
  return (
    <Card>
      <CardHeader className="pb-3">
        <CardTitle className="flex items-center gap-2 text-base">
          <Flame className="h-5 w-5 text-orange-500" />
          Streak của bạn
        </CardTitle>
      </CardHeader>
      <CardContent className="space-y-4">
        <div className="flex items-center gap-3">
          <span className="flex h-11 w-11 items-center justify-center rounded-full bg-orange-100 text-2xl">
            🔥
          </span>
          <div>
            <p className="text-xl font-bold">
              {streakDays} <span className="text-sm font-normal text-muted-foreground">ngày liên tiếp</span>
            </p>
            <p className="text-xs font-medium text-orange-600">Cố gắng quá! 🔥</p>
          </div>
        </div>

        <div className="flex justify-between">
          {DAYS.map((day, i) => (
            <div key={day} className="flex flex-col items-center gap-1.5">
              <span className="text-[11px] text-muted-foreground">{day}</span>
              <span
                className={cn(
                  'flex h-8 w-8 items-center justify-center rounded-full border-2',
                  i === activeDayIndex ? 'border-primary' : 'border-muted'
                )}
              />
            </div>
          ))}
        </div>

        <button className="flex w-full items-center justify-between rounded-xl bg-accent px-4 py-3 text-left">
          <div>
            <p className="text-sm font-semibold text-accent-foreground">Điểm danh hôm nay</p>
            <p className="text-xs text-primary">Giữ chuỗi ngày học · +5 XP</p>
          </div>
          <ChevronRight className="h-4 w-4 text-muted-foreground" />
        </button>

        <button className="flex w-full items-center justify-between rounded-xl border px-4 py-3 text-left">
          <div className="flex items-center gap-2 text-sm font-medium">
            <Trophy className="h-4 w-4 text-amber-500" />
            Kỷ lục: {bestStreak} ngày
          </div>
          <ChevronRight className="h-4 w-4 text-muted-foreground" />
        </button>
      </CardContent>
    </Card>
  );
}
