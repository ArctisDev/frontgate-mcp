import Link from "next/link";

export function Navigation() {
  return (
    <header className="fixed top-0 left-0 right-0 z-50 border-b border-white/[0.06] bg-navy-950/80 backdrop-blur-xl">
      <nav className="mx-auto flex h-16 max-w-6xl items-center justify-between px-6">
        <Link href="/" className="flex items-center gap-2.5">
          <div className="flex h-8 w-8 items-center justify-center rounded-lg bg-blue-500/15">
            <svg
              width="18"
              height="18"
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
          <span className="text-[15px] font-semibold tracking-tight text-slate-200">
            Arctis
          </span>
        </Link>

        <div className="hidden items-center gap-8 md:flex">
          <NavLink href="#benefits">Benefits</NavLink>
          <NavLink href="#how-it-works">How it works</NavLink>
          <NavLink href="#features">Features</NavLink>
          <NavLink href="#pricing">Pricing</NavLink>
        </div>

        <div className="flex items-center gap-3">
          <Link
            href="#"
            className="hidden text-[13px] font-medium text-slate-400 transition-colors hover:text-slate-200 sm:block"
          >
            Sign in
          </Link>
          <Link
            href="#"
            className="flex h-9 items-center rounded-lg bg-blue-500 px-4 text-[13px] font-medium text-white transition-all hover:bg-blue-400"
          >
            Start deploying
          </Link>
        </div>
      </nav>
    </header>
  );
}

function NavLink({
  href,
  children,
}: {
  href: string;
  children: React.ReactNode;
}) {
  return (
    <Link
      href={href}
      className="text-[13px] font-medium text-slate-400 transition-colors hover:text-slate-200"
    >
      {children}
    </Link>
  );
}
