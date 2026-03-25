export function Features() {
  return (
    <section id="features" className="py-[var(--spacing-section)]">
      <div className="section-container">
        <div className="max-w-2xl mb-16">
          <p className="section-label">Capabilities</p>
          <h2 className="section-title">
            Infrastructure primitives that actually work together.
          </h2>
        </div>

        <div className="grid md:grid-cols-2 gap-5">
          {features.map((feature) => (
            <div
              key={feature.title}
              className="rounded-xl bg-bg-raised/40 border border-border/40 p-7 hover:border-border-accent/40 transition-colors duration-200"
            >
              <div className="flex items-start gap-4">
                <div className="shrink-0 w-9 h-9 rounded-lg bg-accent-subtle flex items-center justify-center mt-0.5">
                  <feature.icon className="w-[18px] h-[18px] text-accent" />
                </div>
                <div>
                  <h3 className="text-[15px] font-semibold text-text-primary mb-1.5">
                    {feature.title}
                  </h3>
                  <p className="text-sm text-text-secondary leading-relaxed">
                    {feature.description}
                  </p>
                </div>
              </div>
            </div>
          ))}
        </div>
      </div>
    </section>
  );
}

function GlobeIcon(props: React.SVGProps<SVGSVGElement>) {
  return (
    <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth={1.5} {...props}>
      <path d="M12 2a10 10 0 100 20 10 10 0 000-20zM2 12h20M12 2a15.3 15.3 0 014 10 15.3 15.3 0 01-4 10 15.3 15.3 0 01-4-10A15.3 15.3 0 0112 2z" strokeLinecap="round" strokeLinejoin="round" />
    </svg>
  );
}

function LayersIcon(props: React.SVGProps<SVGSVGElement>) {
  return (
    <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth={1.5} {...props}>
      <path d="M12 2L2 7l10 5 10-5-10-5zM2 17l10 5 10-5M2 12l10 5 10-5" strokeLinecap="round" strokeLinejoin="round" />
    </svg>
  );
}

function GitBranchIcon(props: React.SVGProps<SVGSVGElement>) {
  return (
    <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth={1.5} {...props}>
      <path d="M6 3v12M18 9a3 3 0 100-6 3 3 0 000 6zM6 21a3 3 0 100-6 3 3 0 000 6zM18 9a9 9 0 01-9 9" strokeLinecap="round" strokeLinejoin="round" />
    </svg>
  );
}

function EyeIcon(props: React.SVGProps<SVGSVGElement>) {
  return (
    <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth={1.5} {...props}>
      <path d="M1 12s4-8 11-8 11 8 11 8-4 8-11 8-11-8-11-8z" strokeLinecap="round" strokeLinejoin="round" />
      <circle cx="12" cy="12" r="3" />
    </svg>
  );
}

function LockIcon(props: React.SVGProps<SVGSVGElement>) {
  return (
    <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth={1.5} {...props}>
      <rect x="3" y="11" width="18" height="11" rx="2" ry="2" />
      <path d="M7 11V7a5 5 0 0110 0v4" strokeLinecap="round" strokeLinejoin="round" />
    </svg>
  );
}

function ZapIcon(props: React.SVGProps<SVGSVGElement>) {
  return (
    <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth={1.5} {...props}>
      <path d="M13 2L3 14h9l-1 8 10-12h-9l1-8z" strokeLinecap="round" strokeLinejoin="round" />
    </svg>
  );
}

const features = [
  {
    icon: GlobeIcon,
    title: "Multi-region by default",
    description:
      "Deploy to any combination of 12 global regions. Route traffic intelligently with built-in edge networking.",
  },
  {
    icon: LayersIcon,
    title: "Infrastructure as code",
    description:
      "Define everything in Terraform or Arctis HCL. No drift between your config and what's actually running.",
  },
  {
    icon: GitBranchIcon,
    title: "Preview environments",
    description:
      "Every pull request gets an isolated environment with its own URL. Review changes before they hit production.",
  },
  {
    icon: EyeIcon,
    title: "Built-in observability",
    description:
      "Logs, metrics, traces, and alerts — integrated from the start. No third-party agents to install or maintain.",
  },
  {
    icon: LockIcon,
    title: "SOC 2 & GDPR ready",
    description:
      "Enterprise-grade security with encrypted secrets, audit logs, RBAC, and compliance certifications built in.",
  },
  {
    icon: ZapIcon,
    title: "Instant rollbacks",
    description:
      "One click to revert to any previous deploy. Your data stays intact while the application state rolls back safely.",
  },
];
