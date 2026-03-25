import Image from "next/image";

const metrics = [
  { label: "Weather Shield", value: "20K / 20K" },
  { label: "Thermal Floor", value: "-45 C" },
  { label: "Layer Weight", value: "840 G" },
  { label: "Series", value: "AW/45" },
];

export function Hero() {
  return (
    <section className="section-shell section-block pt-8 lg:pt-14">
      <div className="grid gap-10 lg:grid-cols-2">
        <div className="flex flex-col gap-8">
          <div className="space-y-5">
            <div className="flex flex-wrap items-center gap-4">
              <p className="section-number">01 / Manifesto</p>
              <p className="eyebrow text-accent">Extreme climate · editorial presence</p>
            </div>
            <div className="space-y-4">
              <h1 className="display-kicker">
                BUILT FOR
                <br />
                WHITEOUT
                <span className="block text-accent">CITY PRESSURE</span>
              </h1>
              <div className="accent-rule max-w-md" />
            </div>
            <p className="body-lead">
              POLAR//ARC constrói outerwear técnico com gesto editorial: volumes arquitetônicos,
              proteção climática agressiva e leitura visual imediata para gelo, concreto e vento.
            </p>
          </div>

          <div className="flex flex-wrap gap-3">
            <a href="#colecao" className="button-primary">
              Ver campanha
            </a>
            <a href="#manifesto" className="button-secondary">
              Ler manifesto
            </a>
          </div>

          <div className="grid gap-3 sm:grid-cols-2 lg:grid-cols-4">
            {metrics.map((metric) => (
              <div key={metric.label} className="panel rounded-[1.4rem] px-5 py-4">
                <p className="caption-label">{metric.label}</p>
                <p className="mt-4 text-2xl font-semibold uppercase tracking-[-0.02em]">
                  {metric.value}
                </p>
              </div>
            ))}
          </div>
        </div>

        <div className="panel hero-visual rounded-[2rem] p-0">
          <div className="h-[420px] overflow-hidden rounded-t-[2rem] lg:h-[520px]">
            <Image
              src="https://images.unsplash.com/photo-1515886657613-9f3515b0c78f?auto=format&fit=crop&w=1200&q=80"
              alt="POLAR//ARC manifesto look"
              width={1200}
              height={900}
              className="h-full w-full object-cover"
              priority
            />
          </div>
          <div className="space-y-5 px-6 py-6">
            <div className="flex flex-wrap items-center justify-between gap-4">
              <div>
                <p className="caption-label">Field note</p>
                <p className="mt-2 text-base leading-relaxed text-white/80">
                  Sharp silhouettes, burnt-orange signals e camadas seladas resistem à neve e à
                  pressão urbana sem perder presença.
                </p>
              </div>
              <p className="font-mono text-xs uppercase tracking-[0.3em] text-accent">68°N</p>
            </div>
            <div className="grid gap-4 border-t border-white/10 pt-4 sm:grid-cols-2">
              <div>
                <p className="caption-label">Sistema</p>
                <p className="mt-2 text-lg font-semibold uppercase">Storm membrane</p>
              </div>
              <div>
                <p className="caption-label">Construção</p>
                <p className="mt-2 text-lg font-semibold uppercase">Modular insulation</p>
              </div>
            </div>
          </div>
        </div>
      </div>
    </section>
  );
}
