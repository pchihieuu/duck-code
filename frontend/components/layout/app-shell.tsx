import { Navbar } from './navbar';
import { Sidebar } from './sidebar';

// Layout dùng chung: Navbar full-width + Sidebar 2 cấp + vùng nội dung.
// Dùng cho trang chủ (public, dạng "dashboard preview") và toàn bộ khu vực
// học viên (student). Tách riêng để không lặp code giữa app/page.tsx và
// app/(student)/layout.tsx.
export function AppShell({ children }: { children: React.ReactNode }) {
  return (
    <div className="flex min-h-screen flex-col">
      <Navbar />
      <div className="flex flex-1">
        <Sidebar />
        <main className="flex-1 bg-muted/40 p-4 md:p-8">{children}</main>
      </div>
    </div>
  );
}
