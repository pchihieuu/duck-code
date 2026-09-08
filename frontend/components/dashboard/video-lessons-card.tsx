import Link from 'next/link';
import { PlayCircle, ChevronRight } from 'lucide-react';
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';

interface VideoItem {
  id: string;
  title: string;
  gradient: string; // tailwind gradient classes làm ảnh nền tạm thời
}

const videos: VideoItem[] = [
  { id: '1', title: 'Debug cơ bản trong VS Code', gradient: 'from-rose-200 to-orange-100' },
  { id: '2', title: 'Một ngày làm việc của dev', gradient: 'from-amber-100 to-rose-100' },
];

export function VideoLessonsCard() {
  return (
    <Card>
      <CardHeader className="flex flex-row items-center justify-between pb-3">
        <CardTitle className="flex items-center gap-2 text-base">
          <span className="flex h-7 w-7 items-center justify-center rounded-full bg-green-100">
            <PlayCircle className="h-4 w-4 text-green-600" />
          </span>
          Bài giảng video
          <span className="rounded-full bg-green-100 px-2 py-0.5 text-[11px] font-semibold text-green-700">
            MIỄN PHÍ
          </span>
        </CardTitle>
        <Link href="#" className="flex items-center text-xs font-semibold text-primary">
          Tất cả video <ChevronRight className="h-3.5 w-3.5" />
        </Link>
      </CardHeader>
      <CardContent className="grid grid-cols-2 gap-3">
        {videos.map((video) => (
          <div key={video.id} className="space-y-2">
            <div
              className={`relative flex aspect-video items-center justify-center rounded-xl bg-gradient-to-br ${video.gradient}`}
            >
              <span className="flex h-10 w-10 items-center justify-center rounded-full bg-white/90 shadow">
                <PlayCircle className="h-6 w-6 text-primary" />
              </span>
            </div>
            <p className="line-clamp-2 text-sm font-medium">{video.title}</p>
          </div>
        ))}
      </CardContent>
    </Card>
  );
}
