'use client';

import * as React from 'react';
import Link from 'next/link';
import * as Dialog from '@radix-ui/react-dialog';
import { X } from 'lucide-react';
import { ROUTES } from '@/lib/constants';

export function MobileNav() {
  const [open, setOpen] = React.useState(false);

  return (
    <Dialog.Root open={open} onOpenChange={setOpen}>
      <Dialog.Trigger asChild>
        <button className="md:hidden">Menu</button>
      </Dialog.Trigger>
      <Dialog.Portal>
        <Dialog.Overlay className="fixed inset-0 bg-black/50" />
        <Dialog.Content className="fixed inset-y-0 left-0 w-64 bg-background p-4">
          <Dialog.Close asChild>
            <button className="mb-4">
              <X className="h-5 w-5" />
            </button>
          </Dialog.Close>
          <nav className="flex flex-col gap-2">
            <Link href={ROUTES.dashboard}>Dashboard</Link>
            <Link href={ROUTES.courses}>Khoá học</Link>
            <Link href={ROUTES.leaderboard}>Bảng xếp hạng</Link>
            <Link href={ROUTES.profile}>Hồ sơ</Link>
          </nav>
        </Dialog.Content>
      </Dialog.Portal>
    </Dialog.Root>
  );
}
