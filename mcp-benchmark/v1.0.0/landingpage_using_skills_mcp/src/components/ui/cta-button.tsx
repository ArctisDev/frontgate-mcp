import Link from "next/link";
import { type ComponentProps } from "react";

type ButtonVariant = "primary" | "secondary";

interface CTAButtonProps extends ComponentProps<typeof Link> {
  variant?: ButtonVariant;
}

export function CTAButton({
  variant = "primary",
  className,
  children,
  ...props
}: CTAButtonProps) {
  const base =
    "flex h-12 w-full items-center justify-center rounded-xl px-8 text-[15px] font-medium transition-all sm:w-auto";

  const variants: Record<ButtonVariant, string> = {
    primary:
      "bg-blue-500 text-white shadow-lg shadow-blue-500/20 hover:bg-blue-400 hover:shadow-blue-400/25",
    secondary:
      "border border-white/[0.1] bg-white/[0.03] text-slate-300 hover:border-white/[0.15] hover:bg-white/[0.06]",
  };

  return (
    <Link className={`${base} ${variants[variant]} ${className ?? ""}`} {...props}>
      {children}
    </Link>
  );
}
