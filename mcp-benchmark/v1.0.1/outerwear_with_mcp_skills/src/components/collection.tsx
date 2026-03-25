import Image from "next/image";

const lines = [
  {
    code: "A1",
    name: "Zenith Shell",
    copy: "Escudo externo de leitura escultórica, ombros altos e costura selada com expressão brutalista limpa.",
  },
  {
    code: "B4",
    name: "Tectonic Down",
    copy: "Câmaras densas, volume controlado e compressão tática para trânsito urbano e altitude.",
  },
  {
    code: "C7",
    name: "Signal Parka",
    copy: "Marcações em laranja queimado funcionam como assinatura de presença e leitura noturna.",
  },
];

export function CollectionShowcase() {
  return (
    <section id="colecao" className="section-shell section-block">
      <div className="grid gap-6 xl:grid-cols-[0.34fr_0.66fr]">
        <div className="flex flex-col justify-between gap-10 xl:pr-6">
          <div className="space-y-4">
            <p className="section-number">02 / Collection</p>
            <p className="eyebrow">Cinematic line / AW 45</p>
            <h2 className="section-title max-w-sm">
              Collection as
              <span className="block text-accent">campaign object</span>
            </h2>
          </div>
          <p className="body-lead max-w-sm text-base">
            A linha principal opera como filme congelado: quadro dominante, blocos laterais enxutos
            e leitura imediata da silhueta antes da ficha técnica.
          </p>
        </div>

        <div className="grid gap-6 lg:grid-cols-[1.15fr_0.85fr]">
          <div className="panel overflow-hidden rounded-[2rem]">
            <div className="h-[440px] w-full">
              <Image
                src="https://images.unsplash.com/photo-1491553895911-0055eca6402d?auto=format&fit=crop&w=1400&q=80"
                alt="Colecao principal POLAR ARC"
                width={1200}
                height={900}
                className="h-full w-full object-cover"
              />
            </div>
            <div className="space-y-4 px-6 py-6">
              <p className="caption-label">Dominant frame</p>
              <p className="text-3xl font-semibold uppercase leading-[1.05] tracking-[-0.04em]">
                The alpine silhouette for city weather fronts
              </p>
              <p className="text-base text-white/70">
                Ombros altos, selagem manual e painel reflexivo posicionam a peça como escultura climática.
              </p>
            </div>
          </div>

          <div className="flex flex-col gap-4">
            {lines.map((line) => (
              <article key={line.code} className="panel rounded-[1.8rem] px-6 py-6">
                <div className="flex items-center justify-between gap-4">
                  <p className="caption-label">Line {line.code}</p>
                  <span className="font-mono text-xs uppercase tracking-[0.34em] text-accent">
                    Limited
                  </span>
                </div>
                <h3 className="mt-5 text-3xl font-semibold uppercase tracking-[-0.05em]">
                  {line.name}
                </h3>
                <p className="mt-3 text-base leading-relaxed text-white/70">{line.copy}</p>
              </article>
            ))}
          </div>
        </div>
      </div>
    </section>
  );
}
