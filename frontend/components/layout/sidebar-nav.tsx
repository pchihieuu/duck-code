// 'use client';

// import * as React from 'react';
// import Link from 'next/link';
// import { usePathname } from 'next/navigation';
// import * as Collapsible from '@radix-ui/react-collapsible';
// import {
//   LayoutDashboard,
//   Smartphone,
//   GraduationCap,
//   Headphones,
//   Trophy,
//   Award,
//   User,
//   Settings,
//   type LucideIcon,
// } from 'lucide-react';

// export type IconName = keyof typeof ICONS;
// import { cn } from '@/lib/utils';

// export const ICONS = {
//   dashboard: LayoutDashboard,
//   smartphone: Smartphone,
//   graduationCap: GraduationCap,
//   headphones: Headphones,
//   trophy: Trophy,
//   award: Award,
//   user: User,
//   settings: Settings,
// } as const;


// export interface NavLeaf {
//   label: string;
//   href: string;
// }

// export interface NavGroup {
//   label: string;
//   icon: LucideIcon;
//   items: NavLeaf[];
//   defaultOpen?: boolean;
// }

// export interface NavLink {
//   label: string;
//   href: string;
//   icon: IconName;
// }

// // Một mục đơn (không có submenu) — ví dụ "Dashboard".
// export function SidebarLink({ label, href, icon: Icon }: NavLink) {
//   const Icon = ICONS[icon]; 
//   const pathname = usePathname();
//   const isActive = pathname === href;

//   return (
//     <Link
//       href={href}
//       className={cn(
//         'flex items-center gap-3 rounded-xl px-3 py-2.5 text-sm font-semibold transition-colors',
//         isActive
//           ? 'bg-accent text-accent-foreground'
//           : 'text-muted-foreground hover:bg-accent/60 hover:text-foreground'
//       )}
//     >
//       <Icon className="h-[18px] w-[18px]" />
//       {label}
//     </Link>
//   );
// }

// // Một nhóm có thể xổ ra (collapsible) chứa danh sách link con —
// // giống "Khóa học" / "Kỹ năng" trong ảnh tham khảo.
// export function SidebarGroup({ label, icon: Icon, items, defaultOpen }: NavGroup) {
//   const pathname = usePathname();
//   const hasActiveChild = items.some((item) => pathname === item.href);
//   const [open, setOpen] = React.useState(defaultOpen ?? hasActiveChild);

//   return (
//     <Collapsible.Root open={open} onOpenChange={setOpen}>
//       <Collapsible.Trigger asChild>
//         <button
//           className={cn(
//             'flex w-full items-center gap-3 rounded-xl px-3 py-2.5 text-sm font-semibold text-foreground/90 transition-colors hover:bg-accent/60'
//           )}
//         >
//           <span className="flex h-8 w-8 shrink-0 items-center justify-center rounded-full bg-accent">
//             <Icon className="h-[18px] w-[18px] text-primary" />
//           </span>
//           <span className="flex-1 text-left">{label}</span>
//           <ChevronDown
//             className={cn('h-4 w-4 text-muted-foreground transition-transform', open && 'rotate-180')}
//           />
//         </button>
//       </Collapsible.Trigger>
//       <Collapsible.Content className="overflow-hidden data-[state=closed]:animate-none">
//         <ul className="ml-[22px] mt-1 space-y-0.5 border-l pl-4">
//           {items.map((item) => {
//             const isActive = pathname === item.href;
//             return (
//               <li key={item.href}>
//                 <Link
//                   href={item.href}
//                   className={cn(
//                     'block rounded-lg px-2.5 py-2 text-sm transition-colors',
//                     isActive
//                       ? 'font-semibold text-primary'
//                       : 'text-muted-foreground hover:text-foreground'
//                   )}
//                 >
//                   {item.label}
//                 </Link>
//               </li>
//             );
//           })}
//         </ul>
//       </Collapsible.Content>
//     </Collapsible.Root>
//   );
// }

// export function SidebarSectionLabel({ children }: { children: React.ReactNode }) {
//   return (
//     <p className="px-3 pb-1 pt-4 text-xs font-semibold uppercase tracking-wider text-muted-foreground/70">
//       {children}
//     </p>
//   );
// }


'use client';

import * as React from 'react';
import Link from 'next/link';
import { usePathname } from 'next/navigation';
import * as Collapsible from '@radix-ui/react-collapsible';
import {
  LayoutDashboard,
  Smartphone,
  GraduationCap,
  Headphones,
  Trophy,
  Award,
  User,
  Settings,
  ChevronDown,
} from 'lucide-react';
import { cn } from '@/lib/utils';

export const ICONS = {
  dashboard: LayoutDashboard,
  smartphone: Smartphone,
  graduationCap: GraduationCap,
  headphones: Headphones,
  trophy: Trophy,
  award: Award,
  user: User,
  settings: Settings,
} as const;

export type IconName = keyof typeof ICONS;

export interface NavLeaf {
  label: string;
  href: string;
}

export interface NavGroup {
  label: string;
  icon: IconName;
  items: NavLeaf[];
  defaultOpen?: boolean;
}

export interface NavLink {
  label: string;
  href: string;
  icon: IconName;
}

// Một mục đơn (không có submenu) — ví dụ "Dashboard".
export function SidebarLink({ label, href, icon }: NavLink) {
  const Icon = ICONS[icon];
  const pathname = usePathname();
  const isActive = pathname === href;

  return (
    <Link
      href={href}
      className={cn(
        'flex items-center gap-3 rounded-xl px-3 py-2.5 text-sm font-semibold transition-colors',
        isActive
          ? 'bg-accent text-accent-foreground'
          : 'text-muted-foreground hover:bg-accent/60 hover:text-foreground'
      )}
    >
      <Icon className="h-[18px] w-[18px]" />
      {label}
    </Link>
  );
}

// Một nhóm có thể xổ ra (collapsible) chứa danh sách link con —
// giống "Khóa học" / "Kỹ năng" trong ảnh tham khảo.
export function SidebarGroup({ label, icon, items, defaultOpen }: NavGroup) {
  const Icon = ICONS[icon];
  const pathname = usePathname();
  const hasActiveChild = items.some((item) => pathname === item.href);
  const [open, setOpen] = React.useState(defaultOpen ?? hasActiveChild);

  return (
    <Collapsible.Root open={open} onOpenChange={setOpen}>
      <Collapsible.Trigger asChild>
        <button
          className={cn(
            'flex w-full items-center gap-3 rounded-xl px-3 py-2.5 text-sm font-semibold text-foreground/90 transition-colors hover:bg-accent/60'
          )}
        >
          <span className="flex h-8 w-8 shrink-0 items-center justify-center rounded-full bg-accent">
            <Icon className="h-[18px] w-[18px] text-primary" />
          </span>
          <span className="flex-1 text-left">{label}</span>
          <ChevronDown
            className={cn('h-4 w-4 text-muted-foreground transition-transform', open && 'rotate-180')}
          />
        </button>
      </Collapsible.Trigger>
      <Collapsible.Content className="overflow-hidden data-[state=closed]:animate-none">
        <ul className="ml-[22px] mt-1 space-y-0.5 border-l pl-4">
          {items.map((item) => {
            const isActive = pathname === item.href;
            return (
              <li key={item.href}>
                <Link
                  href={item.href}
                  className={cn(
                    'block rounded-lg px-2.5 py-2 text-sm transition-colors',
                    isActive
                      ? 'font-semibold text-primary'
                      : 'text-muted-foreground hover:text-foreground'
                  )}
                >
                  {item.label}
                </Link>
              </li>
            );
          })}
        </ul>
      </Collapsible.Content>
    </Collapsible.Root>
  );
}

export function SidebarSectionLabel({ children }: { children: React.ReactNode }) {
  return (
    <p className="px-3 pb-1 pt-4 text-xs font-semibold uppercase tracking-wider text-muted-foreground/70">
      {children}
    </p>
  );
}