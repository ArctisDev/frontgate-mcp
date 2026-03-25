import { type ReactNode } from "react";

interface CardBaseProps {
  children: ReactNode;
  className?: string;
  hover?: boolean;
}

export function CardBase({
  children,
  className,
  hover = true,
}: CardBaseProps) {
  const hoverStyles = hover
    ? "transition-all hover:border-white/[0.1] hover:bg-white/[0.04]"
    : "";

  return (
    <div
      className={`rounded-2xl border border-white/[0.06] bg-white/[0.02] p-7 ${hoverStyles} ${className ?? ""}`}
    >
      {children}
    </div>
  );
}
