import Link from "next/link";

export function Hero() {
  return (
    <section className="relative pt-32 pb-20 md:pt-44 md:pb-32 overflow-hidden">
      {/* Subtle ambient glow */}
      <div
        className="absolute top-0 left-1/2 -translate-x-1/2 w-[800px] h-[600px] pointer-events-none"
        style={{
          background:
            "radial-gradient(ellipse at center, rgba(59,130,246,0.06) 0%, transparent 70%)",
        }}
      />

      <div className="section-container relative">
        <div className="max-w-3xl">
          <p className="section-label mb-5">Infrastructure without the overhead</p>

          <h1 className="text-4xl md:text-[3.5rem] lg:text-6xl font-bold tracking-tight leading-[1.1] text-text-primary">
            Deploy on your terms.
            <br />
            <span className="text-text-secondary">Scale on your budget.</span>
          </h1>

          <p className="mt-6 text-lg md:text-xl text-text-secondary leading-relaxed max-w-2xl">
            Arctis Deploy gives engineering teams full control of their
            infrastructure — no proprietary abstractions, no hidden egress
            fees, no surprises when you scale. Ship fast with the freedom
            to move when you need to.
          </p>

          <div className="mt-10 flex flex-col sm:flex-row gap-4">
            <Link
              href="#start"
              className="inline-flex h-12 items-center justify-center rounded-lg bg-accent px-8 text-sm font-semibold text-white hover:bg-accent-hover transition-colors duration-150"
            >
              Start free — no credit card
            </Link>
            <Link
              href="#docs"
              className="inline-flex h-12 items-center justify-center rounded-lg border border-border bg-bg-raised px-8 text-sm font-medium text-text-secondary hover:text-text-primary hover:border-border-accent transition-colors duration-150"
            >
              Read the docs
            </Link>
          </div>
        </div>

        {/* Stats row */}
        <div className="mt-20 md:mt-28 grid grid-cols-2 md:grid-cols-4 gap-6 md:gap-12 max-w-3xl">
          {heroStats.map((stat) => (
            <div key={stat.label}>
              <p className="text-2xl md:text-3xl font-bold tracking-tight text-text-primary">
                {stat.value}
              </p>
              <p className="mt-1 text-sm text-text-tertiary">{stat.label}</p>
            </div>
          ))}
        </div>
      </div>
    </section>
  );
}

const heroStats = [
  { value: "38%", label: "Lower avg. infra cost" },
  { value: "< 4 min", label: "Time to first deploy" },
  { value: "99.99%", label: "Uptime SLA" },
  { value: "12", label: "Cloud regions" },
];
