import Link from "next/link";

export function Footer() {
  return (
    <footer className="border-t border-white/[0.04] py-12">
      <div className="mx-auto max-w-6xl px-6">
        <div className="flex flex-col items-center justify-between gap-6 md:flex-row">
          {/* Logo & copy */}
          <div className="flex items-center gap-2.5">
            <div className="flex h-7 w-7 items-center justify-center rounded-lg bg-blue-500/15">
              <svg
                width="14"
                height="14"
                viewBox="0 0 24 24"
                fill="none"
                stroke="currentColor"
                strokeWidth="2"
                strokeLinecap="round"
                strokeLinejoin="round"
                className="text-blue-400"
              >
                <path d="M13 2L3 14h9l-1 8 10-12h-9l1-8z" />
              </svg>
            </div>
            <span className="text-[13px] text-slate-500">
              © 2026 Arctis Deploy, Inc.
            </span>
          </div>

          {/* Links */}
          <div className="flex items-center gap-6">
            <FooterLink href="#">Docs</FooterLink>
            <FooterLink href="#">Status</FooterLink>
            <FooterLink href="#">Blog</FooterLink>
            <FooterLink href="#">GitHub</FooterLink>
            <FooterLink href="#">Privacy</FooterLink>
            <FooterLink href="#">Terms</FooterLink>
          </div>
        </div>
      </div>
    </footer>
  );
}

function FooterLink({
  href,
  children,
}: {
  href: string;
  children: React.ReactNode;
}) {
  return (
    <Link
      href={href}
      className="text-[13px] text-slate-500 transition-colors hover:text-slate-300"
    >
      {children}
    </Link>
  );
}
