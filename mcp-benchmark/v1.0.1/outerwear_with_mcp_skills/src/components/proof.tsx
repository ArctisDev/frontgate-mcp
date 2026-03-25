const statements = [
  "SEAL THE ATMOSPHERE",
  "MAKE PRESSURE VISIBLE",
  "ARMOR WITHOUT THE NOISE",
];

export function ProofSection() {
  return (
    <section className="section-shell section-block">
      <div className="grid gap-5 lg:grid-cols-[0.75fr_1.25fr]">
        <div className="panel rounded-[1.8rem] p-7 md:p-8">
          <p className="section-number">06 / Proof</p>
          <p className="eyebrow">Proof / editorial quote</p>
          <p className="mt-8 text-3xl font-semibold uppercase leading-[0.96] tracking-[-0.05em] md:text-4xl">
            &ldquo;The rare outerwear brand that behaves like technical equipment and lands like a
            luxury campaign.&rdquo;
          </p>
          <p className="mt-6 font-mono text-xs uppercase tracking-[0.3em] text-fog">
            Aurora Review / Winter issue
          </p>
        </div>

        <div className="overflow-hidden rounded-[1.8rem] border border-white/10">
          {statements.map((statement, index) => (
            <div
              key={statement}
              className={`px-6 py-7 md:px-8 ${
                index === 1 ? "bg-accent text-black" : "bg-[rgba(255,255,255,0.03)] text-white"
              } ${index !== statements.length - 1 ? "border-b border-white/10" : ""}`}
            >
              <p className="text-2xl font-semibold uppercase tracking-[-0.04em] md:text-4xl">
                {statement}
              </p>
            </div>
          ))}
        </div>
      </div>
    </section>
  );
}
