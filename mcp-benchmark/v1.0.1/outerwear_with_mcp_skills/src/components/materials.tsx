const systems = [
  {
    index: "01",
    title: "Membrane",
    stat: "20K pressure block",
    copy: "Lâmina externa com repelência severa e respirabilidade controlada para transições extremas.",
  },
  {
    index: "02",
    title: "Insulation",
    stat: "Aerogel lattice",
    copy: "Câmaras técnicas isolam o corpo sem inflar a silhueta. Retenção térmica com desenho editorial enxuto.",
  },
  {
    index: "03",
    title: "Cut",
    stat: "Angular mobility",
    copy: "Articulação nos ombros, cintura alta e mangas estruturadas para movimento agressivo sem perder forma.",
  },
  {
    index: "04",
    title: "Build",
    stat: "Laser sealed",
    copy: "Painéis cortados com precisão industrial e manutenção modular em campo em menos de noventa segundos.",
  },
];

export function MaterialSection() {
  return (
    <section id="materiais" className="section-shell section-block">
      <div className="panel rounded-[2rem] p-6 md:p-8 xl:p-10">
        <div className="grid gap-8 xl:grid-cols-[0.42fr_0.58fr]">
          <div className="flex flex-col justify-between gap-8 border-b border-white/10 pb-8 xl:border-b-0 xl:border-r xl:pb-0 xl:pr-10">
            <div className="space-y-4">
              <p className="section-number">03 / Materials</p>
              <p className="eyebrow">System / anti-climate engineering</p>
              <h2 className="section-title max-w-md">
                Four layers. <span className="text-accent">One pressure language.</span>
              </h2>
              <p className="body-lead text-base">
                Em vez de cards genéricos, a seção opera como ficha técnica editorial: índice grande,
                linhas horizontais e ritmo de laboratório visual.
              </p>
            </div>
            <p className="display-kicker hidden text-white/10 xl:block">04</p>
          </div>

          <div className="grid gap-4">
            {systems.map((system) => (
              <article
                key={system.index}
                className="group grid gap-4 border-t border-white/10 py-5 md:grid-cols-[90px_1fr_220px] md:items-start"
              >
                <p className="font-mono text-sm uppercase tracking-[0.34em] text-accent">{system.index}</p>
                <div>
                  <h3 className="text-2xl font-semibold uppercase tracking-[-0.04em]">
                    {system.title}
                  </h3>
                  <p className="mt-3 max-w-xl text-base leading-relaxed text-white/70">
                    {system.copy}
                  </p>
                </div>
                <div className="pt-1 md:text-right">
                  <p className="caption-label">Performance</p>
                  <p className="mt-2 text-lg font-semibold uppercase tracking-[-0.03em] text-white/88">
                    {system.stat}
                  </p>
                </div>
              </article>
            ))}
          </div>
        </div>
      </div>
    </section>
  );
}
