import { Handshake, ExternalLink, ChevronRight } from 'lucide-react';
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';

interface Mentor {
  name: string;
  detail: string;
}

export function PartnersCard({ mentors }: { mentors: Mentor[] }) {
  return (
    <Card>
      <CardHeader className="pb-3">
        <CardTitle className="flex items-center gap-2 text-base">
          <span className="flex h-7 w-7 items-center justify-center rounded-full bg-emerald-100">
            <Handshake className="h-4 w-4 text-emerald-600" />
          </span>
          Mentor &amp; Đối tác
        </CardTitle>
      </CardHeader>
      <CardContent className="space-y-3">
        {mentors.map((mentor) => (
          <div key={mentor.name} className="flex items-start justify-between rounded-xl border p-3">
            <div>
              <p className="flex items-center gap-1.5 text-sm font-semibold">
                {mentor.name}
                <ExternalLink className="h-3 w-3 text-muted-foreground" />
              </p>
              <p className="text-xs text-muted-foreground">{mentor.detail}</p>
            </div>
            <span className="shrink-0 rounded-full bg-emerald-50 px-2 py-0.5 text-[11px] font-semibold text-emerald-700">
              ĐỐI TÁC
            </span>
          </div>
        ))}
        <button className="flex w-full items-center justify-between rounded-xl border px-3 py-2.5 text-sm font-semibold">
          Trở thành đối tác
          <ChevronRight className="h-4 w-4 text-muted-foreground" />
        </button>
      </CardContent>
    </Card>
  );
}
