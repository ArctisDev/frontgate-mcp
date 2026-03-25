export function Credibility() {
  return (
    <section className="border-t border-white/[0.04] py-20">
      <div className="mx-auto max-w-6xl px-6">
        {/* Logos */}
        <p className="mb-10 text-center text-[12px] font-medium uppercase tracking-widest text-slate-500">
          Trusted by engineering teams at
        </p>
        <div className="flex flex-wrap items-center justify-center gap-x-12 gap-y-6">
          {COMPANIES.map((company) => (
            <div
              key={company.name}
              className="flex items-center gap-2 text-slate-500"
            >
              <span className="text-[15px] font-semibold tracking-tight">
                {company.name}
              </span>
            </div>
          ))}
        </div>

        {/* Stats */}
        <div className="mt-16 grid grid-cols-2 gap-8 md:grid-cols-4">
          {STATS.map((stat) => (
            <div key={stat.label} className="text-center">
              <div className="text-3xl font-semibold tracking-tight text-slate-200 md:text-4xl">
                {stat.value}
              </div>
              <div className="mt-1.5 text-[13px] text-slate-500">
                {stat.label}
              </div>
            </div>
          ))}
        </div>
      </div>
    </section>
  );
}

const COMPANIES = [
  { name: "Meridian" },
  { name: "Vantage" },
  { name: "Helios" },
  { name: "Strata" },
  { name: "Prism" },
];

const STATS = [
  { value: "12k+", "label": "Deployments / day" },
  { value: "99.97%", "label": "Uptime SLA" },
  { value: "< 4min", "label": "Avg deploy time" },
  { value: "38", "label": "Global regions" },
];
