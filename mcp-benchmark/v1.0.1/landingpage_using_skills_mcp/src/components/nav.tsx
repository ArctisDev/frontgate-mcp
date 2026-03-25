import Link from "next/link";

const navLinks = [
  { label: "Product", href: "#features" },
  { label: "Pricing", href: "#pricing" },
  { label: "Docs", href: "#docs" },
  { label: "Blog", href: "#blog" },
];

export function Nav() {
  return (
    <nav className="fixed top-0 left-0 right-0 z-50 border-b border-border/50 bg-bg-base/80 backdrop-blur-xl">
      <div className="section-container flex h-16 items-center justify-between">
        <Link href="/" className="flex items-center gap-2.5 group">
          <ArctisLogo />
          <span className="text-base font-semibold tracking-tight text-text-primary">
            Arctis
          </span>
        </Link>

        <div className="hidden md:flex items-center gap-8">
          {navLinks.map((link) => (
            <Link
              key={link.label}
              href={link.href}
              className="text-sm text-text-secondary hover:text-text-primary transition-colors duration-150"
            >
              {link.label}
            </Link>
          ))}
        </div>

        <div className="flex items-center gap-3">
          <Link
            href="#login"
            className="hidden sm:block text-sm text-text-secondary hover:text-text-primary transition-colors duration-150"
          >
            Sign in
          </Link>
          <Link
            href="#start"
            className="inline-flex h-9 items-center rounded-lg bg-accent px-4 text-sm font-medium text-white hover:bg-accent-hover transition-colors duration-150"
          >
            Start deploying
          </Link>
        </div>
      </div>
    </nav>
  );
}

function ArctisLogo() {
  return (
    <svg
      width="24"
      height="24"
      viewBox="0 0 24 24"
      fill="none"
      className="text-accent"
      aria-hidden="true"
    >
      <path
        d="M12 2L3 7v10l9 5 9-5V7l-9-5z"
        stroke="currentColor"
        strokeWidth="1.5"
        strokeLinejoin="round"
      />
      <path
        d="M12 22V12M12 12L3 7M12 12l9-5"
        stroke="currentColor"
        strokeWidth="1.5"
        strokeLinejoin="round"
      />
    </svg>
  );
}
