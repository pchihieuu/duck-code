import type { Metadata } from 'next';
import { Be_Vietnam_Pro } from 'next/font/google';
import { ThemeProvider } from '@/components/shared/theme-provider';
import { cn } from '@/lib/utils';
import './globals.css';

// Be Vietnam Pro: font chính của toàn bộ giao diện, expose qua CSS variable
// để dùng trong tailwind.config.ts (fontFamily.sans).
const beVietnamPro = Be_Vietnam_Pro({
  subsets: ['latin', 'vietnamese'],
  weight: ['400', '500', '600', '700', '800'],
  variable: '--font-be-vietnam-pro',
  display: 'swap',
});

export const metadata: Metadata = {
  title: {
    default: 'Learning Platform',
    template: '%s | Learning Platform',
  },
  description: 'Nền tảng học lập trình tương tác: khoá học, bài tập, quiz và theo dõi tiến độ.',
};

export default function RootLayout({ children }: { children: React.ReactNode }) {
  return (
    <html lang="vi" suppressHydrationWarning className={beVietnamPro.variable}>
      <body className={cn('font-sans antialiased')}>
        <ThemeProvider attribute="class" defaultTheme="light" enableSystem={false}>
          {children}
        </ThemeProvider>
      </body>
    </html>
  );
}
