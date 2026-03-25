export function SocialProof() {
  return (
    <section className="py-16 md:py-20 border-y border-border/50">
      <div className="section-container">
        <p className="text-center text-xs font-medium tracking-widest uppercase text-text-tertiary mb-10">
          Trusted by engineering teams at
        </p>

        <div className="flex flex-wrap items-center justify-center gap-x-12 gap-y-8">
          {companies.map((company) => (
            <span
              key={company.name}
              className="text-lg md:text-xl font-semibold tracking-tight text-text-muted/60 select-none"
            >
              {company.name}
            </span>
          ))}
        </div>
      </div>
    </section>
  );
}

const companies = [
  { name: "NovaCorp" },
  { name: "Helios Systems" },
  { name: "Driftwave" },
  { name: "Paladin Labs" },
  { name: "Axis Cloud" },
  { name: "Stratum" },
];
