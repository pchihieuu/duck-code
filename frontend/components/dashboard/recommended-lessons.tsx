import Link from 'next/link';
import { Sparkles, ChevronRight, PlayCircle, type LucideIcon } from 'lucide-react';

interface RecommendedLesson {
  id: string;
  category: string;
  categoryIcon: LucideIcon;
  title: string;
  subtitle: string;
  level: string;
  gradient: string;
}

const lessons: RecommendedLesson[] = [
  {
    id: '1',
    category: 'Cú pháp',
    categoryIcon: Sparkles,
    title: 'Biến & kiểu dữ liệu',
    subtitle: 'Học theo ví dụ thực tế',
    level: 'Cơ bản',
    gradient: 'from-rose-200 to-red-100',
  },
  {
    id: '2',
    category: 'Cấu trúc dữ liệu',
    categoryIcon: Sparkles,
    title: 'Mảng & danh sách liên kết',
    subtitle: 'So sánh ưu/nhược điểm',
    level: 'Cơ bản',
    gradient: 'from-amber-100 to-rose-100',
  },
  {
    id: '3',
    category: 'Gỡ lỗi',
    categoryIcon: Sparkles,
    title: 'Đọc hiểu thông báo lỗi',
    subtitle: 'Luyện phản xạ debug',
    level: 'Cơ bản',
    gradient: 'from-orange-100 to-amber-100',
  },
  {
    id: '4',
    category: 'Giải thuật',
    categoryIcon: Sparkles,
    title: 'Tư duy giải thuật cơ bản',
    subtitle: 'Chuẩn bị phỏng vấn kỹ thuật',
    level: 'Cơ bản',
    gradient: 'from-rose-100 to-orange-100',
  },
];

export function RecommendedLessons() {
  return (
    <section>
      <div className="mb-4 flex items-center justify-between">
        <h2 className="flex items-center gap-2 text-lg font-semibold">
          <Sparkles className="h-5 w-5 text-primary" />
          Bài học đề xuất cho bạn
        </h2>
        <Link href="#" className="flex items-center text-sm font-semibold text-primary">
          Xem tất cả <ChevronRight className="h-4 w-4" />
        </Link>
      </div>

      <div className="grid grid-cols-2 gap-4 lg:grid-cols-4">
        {lessons.map((lesson) => (
          <Link
            key={lesson.id}
            href="#"
            className="overflow-hidden rounded-2xl border bg-card transition-shadow hover:shadow-md"
          >
            <div className={`relative flex aspect-square items-center justify-center bg-gradient-to-br ${lesson.gradient}`}>
              <span className="absolute left-3 top-3 flex items-center gap-1 rounded-full bg-white/90 px-2.5 py-1 text-[11px] font-semibold text-foreground">
                <lesson.categoryIcon className="h-3 w-3 text-primary" />
                {lesson.category}
              </span>
              <span className="flex h-11 w-11 items-center justify-center rounded-full bg-primary text-primary-foreground shadow">
                <PlayCircle className="h-6 w-6" />
              </span>
            </div>
            <div className="p-3">
              <p className="truncate font-semibold">{lesson.title}</p>
              <div className="mt-1 flex items-center justify-between">
                <span className="truncate text-xs text-muted-foreground">{lesson.subtitle}</span>
                <span className="shrink-0 rounded-full bg-muted px-2 py-0.5 text-[11px] font-medium text-muted-foreground">
                  {lesson.level}
                </span>
              </div>
            </div>
          </Link>
        ))}
      </div>
    </section>
  );
}
