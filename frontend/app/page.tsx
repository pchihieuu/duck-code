import { BookOpen, ListChecks, ClipboardCheck, Layers } from 'lucide-react';
import { AppShell } from '@/components/layout/app-shell';
import { HeroBanner } from '@/components/dashboard/hero-banner';
import { StreakCtaCard } from '@/components/dashboard/streak-cta-card';
import { VideoLessonsCard } from '@/components/dashboard/video-lessons-card';
import { LearningResourcesCard } from '@/components/dashboard/learning-resources-card';
import { PartnersCard } from '@/components/dashboard/partners-card';
import { RecommendedLessons } from '@/components/dashboard/recommended-lessons';

// Trang chủ: hiển thị dạng "dashboard preview" kể cả khi chưa đăng nhập,
// theo đúng bố cục tham khảo (Navbar + Sidebar + hero + streak + các card).
// TODO: khi có auth, thay dữ liệu mock dưới đây bằng dữ liệu thật của user
// (nếu đã đăng nhập) hoặc dữ liệu tổng quan công khai (nếu chưa đăng nhập).
export default function HomePage() {
  const streak = { days: 0, activeDayIndex: 0 };

  const resourceTiles = [
    { icon: BookOpen, value: '1.253', label: 'Bài học', iconBg: '#fde8ea', iconColor: '#e11d48' },
    { icon: ListChecks, value: '10.996', label: 'Bài tập', iconBg: '#ede9fe', iconColor: '#7c3aed' },
    { icon: ClipboardCheck, value: '450+', label: 'Đề luyện tập', iconBg: '#fff2e0', iconColor: '#ea580c' },
    { icon: Layers, value: '9', label: 'Cấp độ', iconBg: '#e6f9f0', iconColor: '#059669' },
  ];

  const mentors = [
    { name: 'Mentor: Nam Trần', detail: 'Backend Engineer · 0933165673' },
    { name: 'Mentor: Quí Mai', detail: 'Dạy lập trình online · 0344682135' },
  ];

  return (
    <AppShell>
      <div className="space-y-6">
        <div className="grid gap-6 lg:grid-cols-3">
          <div className="lg:col-span-2">
            <HeroBanner />
          </div>
          <StreakCtaCard streakDays={streak.days} activeDayIndex={streak.activeDayIndex} />
        </div>

        <div className="grid gap-6 lg:grid-cols-3">
          <VideoLessonsCard />
          <LearningResourcesCard tiles={resourceTiles} />
          <PartnersCard mentors={mentors} />
        </div>

        <RecommendedLessons />
      </div>
    </AppShell>
  );
}
