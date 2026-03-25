import Link from "next/link";

const footerLinks = {
  Product: ["Features", "Pricing", "Changelog", "Roadmap"],
  Developers: ["Documentation", "API Reference", "CLI", "Status"],
  Company: ["About", "Blog", "Careers", "Contact"],
  Legal: ["Privacy", "Terms", "Security", "DPA"],
};

export function Footer() {
  return (
    <footer className="border-t border-border/50 bg-bg-base">
      <div className="section-container py-12 md:py-16">
        <div className="grid grid-cols-2 md:grid-cols-5 gap-8">
          {/* Brand column */}
          <div className="col-span-2 md:col-span-1">
            <div className="flex items-center gap-2 mb-4">
              <svg
                width="20"
                height="20"
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
              <span className="text-sm font-semibold text-text-primary">
                Arctis Deploy
              </span>
            </div>
            <p className="text-xs text-text-tertiary leading-relaxed max-w-[200px]">
              Infrastructure on your terms. Built for teams that value control,
              transparency, and sustainable growth.
            </p>
          </div>

          {/* Link columns */}
          {Object.entries(footerLinks).map(([group, links]) => (
            <div key={group}>
              <p className="text-xs font-semibold text-text-secondary uppercase tracking-wider mb-3">
                {group}
              </p>
              <ul className="space-y-2.5">
                {links.map((link) => (
                  <li key={link}>
                    <Link
                      href={`#${link.toLowerCase()}`}
                      className="text-sm text-text-tertiary hover:text-text-secondary transition-colors duration-150"
                    >
                      {link}
                    </Link>
                  </li>
                ))}
              </ul>
            </div>
          ))}
        </div>

        <div className="gradient-divider mt-12 mb-6" />

        <div className="flex flex-col sm:flex-row items-center justify-between gap-4">
          <p className="text-xs text-text-muted">
            &copy; {new Date().getFullYear()} Arctis Deploy. All rights
            reserved.
          </p>
          <div className="flex items-center gap-5">
            {["GitHub", "Twitter", "Discord"].map((social) => (
              <Link
                key={social}
                href={`#${social.toLowerCase()}`}
                className="text-xs text-text-muted hover:text-text-secondary transition-colors duration-150"
              >
                {social}
              </Link>
            ))}
          </div>
        </div>
      </div>
    </footer>
  );
}
