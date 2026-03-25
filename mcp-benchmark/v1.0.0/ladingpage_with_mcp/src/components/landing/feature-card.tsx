import { cn } from "@/lib/utils";
import { type LucideIcon } from "lucide-react";

interface FeatureCardProps {
  icon: LucideIcon;
  title: string;
  description: string;
  className?: string;
}

export function FeatureCard({
  icon: Icon,
  title,
  description,
  className,
}: FeatureCardProps) {
  return (
    <div
      className={cn(
        "group rounded-xl border border-arctis-border bg-background p-6 transition-all hover:border-arctis-blue/30 hover:bg-arctis-surface-raised",
        className
      )}
    >
      <div className="mb-4 inline-flex h-9 w-9 items-center justify-center rounded-lg border border-arctis-border bg-arctis-surface">
        <Icon className="h-[18px] w-[18px] text-arctis-blue" />
      </div>
      <h3 className="mb-1.5 text-[15px] font-semibold">{title}</h3>
      <p className="text-sm leading-relaxed text-muted-foreground">
        {description}
      </p>
    </div>
  );
}
