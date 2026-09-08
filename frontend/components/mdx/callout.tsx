import { type ReactNode } from 'react';
import { Info, AlertTriangle, Lightbulb } from 'lucide-react';
import { cn } from '@/lib/utils';

const variants = {
  note: { icon: Info, className: 'border-blue-200 bg-blue-50 text-blue-900' },
  tip: { icon: Lightbulb, className: 'border-green-200 bg-green-50 text-green-900' },
  warning: { icon: AlertTriangle, className: 'border-amber-200 bg-amber-50 text-amber-900' },
};

export function Callout({
  variant = 'note',
  children,
}: {
  variant?: keyof typeof variants;
  children: ReactNode;
}) {
  const { icon: Icon, className } = variants[variant];
  return (
    <div className={cn('my-4 flex gap-3 rounded-lg border p-4 text-sm', className)}>
      <Icon className="h-5 w-5 shrink-0" />
      <div>{children}</div>
    </div>
  );
}

export const Tip = (props: { children: ReactNode }) => <Callout variant="tip" {...props} />;
export const Warning = (props: { children: ReactNode }) => <Callout variant="warning" {...props} />;
export const Note = (props: { children: ReactNode }) => <Callout variant="note" {...props} />;
