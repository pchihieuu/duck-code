import { Library, type LucideIcon } from 'lucide-react';
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';

interface ResourceTile {
  icon: LucideIcon;
  value: string;
  label: string;
  iconBg: string;
  iconColor: string;
}

export function LearningResourcesCard({ tiles }: { tiles: ResourceTile[] }) {
  return (
    <Card>
      <CardHeader className="pb-3">
        <CardTitle className="flex items-center gap-2 text-base">
          <span className="flex h-7 w-7 items-center justify-center rounded-full bg-accent">
            <Library className="h-4 w-4 text-primary" />
          </span>
          Kho học liệu
        </CardTitle>
      </CardHeader>
      <CardContent className="grid grid-cols-2 gap-3">
        {tiles.map((tile) => (
          <div key={tile.label} className="flex items-center gap-3 rounded-xl border p-3">
            <span
              className="flex h-9 w-9 shrink-0 items-center justify-center rounded-lg"
              style={{ backgroundColor: tile.iconBg }}
            >
              <tile.icon className="h-[18px] w-[18px]" style={{ color: tile.iconColor }} />
            </span>
            <div>
              <p className="text-lg font-bold leading-tight">{tile.value}</p>
              <p className="text-xs text-muted-foreground">{tile.label}</p>
            </div>
          </div>
        ))}
      </CardContent>
    </Card>
  );
}
