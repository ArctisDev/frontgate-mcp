export function Differentiation() {
  return (
    <section className="border-t border-white/[0.04] py-24">
      <div className="mx-auto max-w-6xl px-6">
        <div className="mb-16 text-center">
          <p className="mb-3 text-[12px] font-medium uppercase tracking-widest text-blue-400">
            Different by design
          </p>
          <h2 className="text-4xl font-semibold tracking-tight text-slate-200 sm:text-5xl">
            What sets Arctis apart
          </h2>
        </div>

        {/* Comparison table */}
        <div className="overflow-hidden rounded-2xl border border-white/[0.06]">
          {/* Header */}
          <div className="grid grid-cols-3 border-b border-white/[0.06] bg-white/[0.02]">
            <div className="p-5 text-[13px] font-medium text-slate-500" />
            <div className="p-5 text-center text-[13px] font-semibold text-slate-300">
              Arctis
            </div>
            <div className="p-5 text-center text-[13px] font-medium text-slate-500">
              Traditional PaaS
            </div>
          </div>

          {/* Rows */}
          {COMPARISON.map((row) => (
            <div
              key={row.feature}
              className="grid grid-cols-3 border-b border-white/[0.04] last:border-b-0"
            >
              <div className="flex items-center p-5 text-[14px] text-slate-400">
                {row.feature}
              </div>
              <div className="flex items-center justify-center border-x border-white/[0.04] p-5">
                <span className="flex items-center gap-2 text-[14px] text-emerald-400">
                  <svg className="h-4 w-4" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round">
                    <polyline points="20 6 9 17 4 12" />
                  </svg>
                  {row.arctis}
                </span>
              </div>
              <div className="flex items-center justify-center p-5">
                <span className="text-[14px] text-slate-500">
                  {row.other}
                </span>
              </div>
            </div>
          ))}
        </div>

        {/* Key differentiators */}
        <div className="mt-16 grid gap-8 md:grid-cols-3">
          {DIFFERENTIATORS.map((item) => (
            <div key={item.title}>
              <h3 className="mb-2 text-[15px] font-semibold text-slate-200">
                {item.title}
              </h3>
              <p className="text-[14px] leading-relaxed text-slate-400">
                {item.description}
              </p>
            </div>
          ))}
        </div>
      </div>
    </section>
  );
}

const COMPARISON = [
  {
    feature: "Infrastructure export",
    arctis: "Full export, anytime",
    other: "Proprietary format",
  },
  {
    feature: "Pricing model",
    arctis: "Usage-based, transparent",
    other: "Per-seat + hidden fees",
  },
  {
    feature: "Database portability",
    arctis: "Standard Postgres/Redis",
    other: "Managed, locked-in",
  },
  {
    feature: "Custom domains & SSL",
    arctis: "Included, auto-renew",
    other: "Often paid add-on",
  },
  {
    feature: "Infrastructure as Code",
    arctis: "Native + Terraform/Pulumi",
    other: "Limited or none",
  },
  {
    feature: "Multi-cloud support",
    arctis: "AWS, GCP, Azure",
    other: "Single provider",
  },
];

const DIFFERENTIATORS = [
  {
    title: "Open by default",
    description:
      "Your infra config is standard formats. Your data is in standard databases. If you leave, everything comes with you.",
  },
  {
    title: "No abstraction ceiling",
    description:
      "Start simple, drop to the metal when you need to. Arctis grows with your requirements, not against them.",
  },
  {
    title: "Billed fairly",
    description:
      "You pay for what you use. No per-seat tax. No premium tiers for basic features. Audit your bill anytime.",
  },
];
