import { Flame } from 'lucide-react';
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';
import { Button } from '@/components/ui/button';
import { cn } from '@/lib/utils';

const DAYS = ['T2', 'T3', 'T4', 'T5', 'T6', 'T7', 'CN'];

// Biến thể "kêu gọi hành động": dùng ở trang chủ / khi chưa có streak nào,
// khác với StreakCard (biến thể điểm danh + kỷ lục dùng ở dashboard nội bộ).
export function StreakCtaCard({
  streakDays,
  activeDayIndex,
}: {
  streakDays: number;
  activeDayIndex: number; // 0 = T2 ... 6 = CN
}) {
  return (
    <Card>
      <CardHeader className="pb-3">
        <CardTitle className="flex items-center gap-2 text-base">
          <Flame className="h-5 w-5 text-orange-500" />
          Chuỗi ngày học
        </CardTitle>
      </CardHeader>
      <CardContent className="space-y-5">
        <div className="flex items-center gap-3">
          <span className="flex h-11 w-11 items-center justify-center rounded-full bg-orange-100 text-2xl">
            🔥
          </span>
          <div>
            <p className="text-xl font-bold">
              {streakDays} <span className="text-sm font-normal text-muted-foreground">ngày liên tiếp</span>
            </p>
            <p className="text-xs text-muted-foreground">
              Học mỗi ngày để giữ chuỗi streak và ghi nhớ lâu hơn.
            </p>
          </div>
        </div>

        <div className="flex justify-between">
          {DAYS.map((day, i) => (
            <div key={day} className="flex flex-col items-center gap-1.5">
              <span
                className={cn(
                  'text-[11px]',
                  i === activeDayIndex ? 'font-semibold text-primary' : 'text-muted-foreground'
                )}
              >
                {day}
              </span>
              <span
                className={cn(
                  'flex h-8 w-8 items-center justify-center rounded-full border-2',
                  i === activeDayIndex ? 'border-primary' : 'border-muted'
                )}
              />
            </div>
          ))}
        </div>

        <Button className="w-full rounded-full" size="lg">
          <Flame className="mr-2 h-4 w-4" />
          Bắt đầu chuỗi của bạn
        </Button>
      </CardContent>
    </Card>
  );
}
