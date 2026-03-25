import Link from "next/link";

const links = [
  { label: "Linha", href: "#colecao" },
  { label: "Sistema", href: "#materiais" },
  { label: "Manifesto", href: "#manifesto" },
  { label: "Sessões", href: "#cta" },
];

export function Header() {
  return (
    <header className="sticky top-0 z-50 border-b border-white/8 bg-[rgba(3,5,10,0.82)] backdrop-blur-xl">
      <div className="section-shell flex min-h-[78px] items-center justify-between gap-6 py-4">
        <div className="flex items-center gap-4">
          <Link href="/" className="text-sm font-semibold uppercase tracking-[0.34em]">
            POLAR//ARC
          </Link>
          <span className="hidden font-mono text-xs uppercase tracking-[0.34em] text-fog md:inline">
            Technical editorial outerwear
          </span>
        </div>
        <div className="flex items-center gap-6">
          <nav className="hidden items-center gap-5 font-mono text-xs uppercase tracking-[0.3em] text-fog md:flex">
            {links.map((link) => (
              <Link key={link.label} href={link.href} className="transition-colors hover:text-white">
                {link.label}
              </Link>
            ))}
          </nav>
          <Link href="#cta" className="button-secondary px-5 py-0 text-xs">
            Appointment
          </Link>
        </div>
      </div>
    </header>
  );
}
