export function Benefits() {
  return (
    <section id="benefits" className="py-[var(--spacing-section)] md:py-[var(--spacing-section)]">
      <div className="section-container">
        <div className="max-w-2xl mb-16">
          <p className="section-label">Why teams switch</p>
          <h2 className="section-title">
            Built for teams that refuse to trade control for convenience.
          </h2>
        </div>

        <div className="grid md:grid-cols-3 gap-6">
          {benefits.map((benefit) => (
            <div
              key={benefit.title}
              className="group rounded-xl bg-bg-raised/50 border border-border/50 p-7 hover:border-border-accent/50 transition-colors duration-200"
            >
              <div className="w-10 h-10 rounded-lg bg-accent-subtle flex items-center justify-center mb-5">
                <benefit.icon className="w-5 h-5 text-accent" />
              </div>
              <h3 className="text-base font-semibold text-text-primary mb-2">
                {benefit.title}
              </h3>
              <p className="text-sm text-text-secondary leading-relaxed">
                {benefit.description}
              </p>
            </div>
          ))}
        </div>
      </div>
    </section>
  );
}

function LockOpenIcon(props: React.SVGProps<SVGSVGElement>) {
  return (
    <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth={1.5} {...props}>
      <path d="M8 11V7a4 4 0 118 0M5 21h14a2 2 0 002-2v-5a2 2 0 00-2-2H5a2 2 0 00-2 2v5a2 2 0 002 2z" strokeLinecap="round" strokeLinejoin="round" />
    </svg>
  );
}

function ChartIcon(props: React.SVGProps<SVGSVGElement>) {
  return (
    <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth={1.5} {...props}>
      <path d="M3 3v18h18M7 16l4-4 4 4 5-8" strokeLinecap="round" strokeLinejoin="round" />
    </svg>
  );
}

function ShieldIcon(props: React.SVGProps<SVGSVGElement>) {
  return (
    <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth={1.5} {...props}>
      <path d="M12 22s8-4 8-10V5l-8-3-8 3v7c0 6 8 10 8 10z" strokeLinecap="round" strokeLinejoin="round" />
    </svg>
  );
}

const benefits = [
  {
    icon: LockOpenIcon,
    title: "Zero lock-in",
    description:
      "Your infra runs on open standards. Kubernetes, Terraform, and your existing CI — not proprietary runtimes you can't migrate away from.",
  },
  {
    icon: ChartIcon,
    title: "Predictable costs",
    description:
      "No surprise egress bills. No per-seat pricing that punishes growth. Costs scale linearly with actual resource consumption.",
  },
  {
    icon: ShieldIcon,
    title: "Production-grade from day one",
    description:
      "Automated TLS, built-in observability, and rolling deploys with zero downtime. Not a toy — infrastructure your team can trust.",
  },
];
