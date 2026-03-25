import { cn } from "@/lib/utils";

interface SectionHeaderProps {
  label: string;
  title: string;
  description?: string;
  align?: "left" | "center";
  className?: string;
}

export function SectionHeader({
  label,
  title,
  description,
  align = "left",
  className,
}: SectionHeaderProps) {
  return (
    <div
      className={cn(
        align === "center" && "text-center",
        align === "left" && "max-w-2xl",
        className
      )}
    >
      <p className="mb-3 text-xs font-semibold uppercase tracking-widest text-arctis-blue">
        {label}
      </p>
      <h2 className="text-3xl font-bold tracking-tight md:text-4xl">{title}</h2>
      {description && (
        <p
          className={cn(
            "mt-4 text-lg text-muted-foreground",
            align === "center" && "mx-auto max-w-xl"
          )}
        >
          {description}
        </p>
      )}
    </div>
  );
}
