import { Button } from "@/components/ui/button";
import { cn } from "@/lib/utils";
import { type LucideIcon } from "lucide-react";

interface CTAButtonProps {
  children: React.ReactNode;
  variant?: "primary" | "outline";
  icon?: LucideIcon;
  className?: string;
}

export function CTAButton({
  children,
  variant = "primary",
  icon: Icon,
  className,
}: CTAButtonProps) {
  const baseStyles =
    "h-12 gap-2 px-8 text-sm font-semibold transition-all";

  if (variant === "primary") {
    return (
      <Button
        size="lg"
        className={cn(
          baseStyles,
          "bg-arctis-blue text-primary-foreground hover:bg-arctis-blue-light",
          className
        )}
      >
        {children}
        {Icon && <Icon className="h-4 w-4" />}
      </Button>
    );
  }

  return (
    <Button
      variant="outline"
      size="lg"
      className={cn(
        baseStyles,
        "border-arctis-border bg-transparent text-foreground hover:bg-arctis-surface-raised",
        className
      )}
    >
      {children}
      {Icon && <Icon className="h-4 w-4" />}
    </Button>
  );
}
