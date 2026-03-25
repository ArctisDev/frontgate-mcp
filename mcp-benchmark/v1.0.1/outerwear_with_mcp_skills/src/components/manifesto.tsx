const principles = [
  "Nós desenhamos clima em vez de acompanhar estação.",
  "Toda peça precisa performar sob pressão e parecer capa de campanha.",
  "A identidade vem antes da tendência; a composição vem antes do efeito.",
];

export function Manifesto() {
  return (
    <section id="manifesto" className="section-shell section-block">
      <div className="grid gap-8 xl:grid-cols-[0.6fr_0.4fr]">
        <div className="panel rounded-[2rem] p-8 md:p-10">
          <p className="section-number">05 / Manifesto</p>
          <p className="eyebrow">Manifesto / institutional tone</p>
          <div className="mt-8 space-y-8">
            <h2 className="display-kicker max-w-4xl text-[clamp(3.5rem,8vw,7rem)]">
              The brand is a weather position.
            </h2>
            <p className="max-w-2xl text-lg leading-relaxed text-white/72">
              POLAR//ARC mistura utilitarismo sofisticado, construção técnica e direção de arte de
              luxo. Não entrega catálogo neutro. Entrega presença, proteção e linguagem visual.
            </p>
          </div>
        </div>

        <div className="grid gap-4">
            {principles.map((principle, index) => (
              <article key={principle} className="panel manifest-line rounded-[1.6rem] px-6 py-6">
                <p className="font-mono text-xs uppercase tracking-[0.32em] text-accent">
                  0{index + 1}
                </p>
                <p className="mt-4 text-xl font-semibold uppercase leading-tight tracking-[-0.04em]">
                  {principle}
                </p>
              </article>
            ))}
        </div>
      </div>
    </section>
  );
}
