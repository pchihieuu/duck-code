// import {
//   LayoutDashboard,
//   Smartphone,
//   GraduationCap,
//   Headphones,
//   Trophy,
//   Award,
//   User,
//   Settings,
// } from 'lucide-react';
// import { SidebarLink, SidebarGroup, SidebarSectionLabel } from './sidebar-nav';
// import { ROUTES } from '@/lib/constants';

// export function Sidebar() {
//   return (
//     <aside className="hidden w-72 shrink-0 overflow-y-auto border-r bg-background px-3 py-4 md:block">
//       <nav className="flex flex-col gap-0.5">
//         <SidebarLink label="Dashboard" href={ROUTES.dashboard} icon={LayoutDashboard} />
//         <SidebarLink label="Tải ứng dụng" href="#" icon={Smartphone} />

//         <SidebarSectionLabel>Học tập</SidebarSectionLabel>
//         <SidebarGroup
//           label="Khoá học"
//           icon={GraduationCap}
//           defaultOpen
//           items={[
//             { label: 'Người mới bắt đầu', href: ROUTES.courses },
//             { label: 'Python cơ bản', href: ROUTES.course('python-basic') },
//             { label: 'Web cơ bản', href: '#' },
//             { label: 'Bài tập', href: '#' },
//             { label: 'Luyện thi chứng chỉ', href: '#' },
//           ]}
//         />
//         <SidebarGroup
//           label="Kỹ năng"
//           icon={Headphones}
//           items={[
//             { label: 'Giải thuật & Cấu trúc dữ liệu', href: '#' },
//             { label: 'Đọc hiểu tài liệu kỹ thuật', href: '#' },
//             { label: 'Viết code sạch', href: '#' },
//             { label: 'Phỏng vấn kỹ thuật', href: '#' },
//           ]}
//         />

//         <SidebarSectionLabel>Tiến trình</SidebarSectionLabel>
//         <SidebarLink label="Bảng xếp hạng" href={ROUTES.leaderboard} icon={Trophy} />
//         <SidebarLink label="Thành tích" href={ROUTES.achievements} icon={Award} />
//         <SidebarLink label="Hồ sơ" href={ROUTES.profile} icon={User} />
//         <SidebarLink label="Cài đặt" href={ROUTES.settings} icon={Settings} />
//       </nav>
//     </aside>
//   );
// }


import { SidebarLink, SidebarGroup, SidebarSectionLabel } from './sidebar-nav';
import { ROUTES } from '@/lib/constants';

export function Sidebar() {
  return (
    <aside className="hidden w-72 shrink-0 overflow-y-auto border-r bg-background px-3 py-4 md:block">
      <nav className="flex flex-col gap-0.5">
        <SidebarLink label="Dashboard" href={ROUTES.dashboard} icon="dashboard" />
        <SidebarLink label="Tải ứng dụng" href="#" icon="smartphone" />

        <SidebarSectionLabel>Học tập</SidebarSectionLabel>
        <SidebarGroup
          label="Khoá học"
          icon="graduationCap"
          defaultOpen
          items={[
            { label: 'Người mới bắt đầu', href: ROUTES.courses },
            { label: 'Python cơ bản', href: ROUTES.course('python-basic') },
            { label: 'Web cơ bản', href: '#' },
            { label: 'Bài tập', href: '#' },
            { label: 'Luyện thi chứng chỉ', href: '#' },
          ]}
        />
        <SidebarGroup
          label="Kỹ năng"
          icon="headphones"
          items={[
            { label: 'Giải thuật & Cấu trúc dữ liệu', href: '#' },
            { label: 'Đọc hiểu tài liệu kỹ thuật', href: '#' },
            { label: 'Viết code sạch', href: '#' },
            { label: 'Phỏng vấn kỹ thuật', href: '#' },
          ]}
        />

        <SidebarSectionLabel>Tiến trình</SidebarSectionLabel>
        <SidebarLink label="Bảng xếp hạng" href={ROUTES.leaderboard} icon="trophy" />
        <SidebarLink label="Thành tích" href={ROUTES.achievements} icon="award" />
        <SidebarLink label="Hồ sơ" href={ROUTES.profile} icon="user" />
        <SidebarLink label="Cài đặt" href={ROUTES.settings} icon="settings" />
      </nav>
    </aside>
  );
}