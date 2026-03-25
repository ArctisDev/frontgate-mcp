import Image from "next/image";

const looks = [
  {
    title: "Meridian Shell",
    note: "Overshell angular com leitura de revista e proteção de tempestade.",
    image: "https://images.unsplash.com/photo-1483985988355-763728e1935b?auto=format&fit=crop&w=1200&q=80",
    className: "md:col-span-2 md:row-span-2",
  },
  {
    title: "Urban Whiteout",
    note: "Sinalização mínima, contraste alto e construção para baixa visibilidade.",
    image: "https://images.unsplash.com/photo-1515886657613-9f3515b0c78f?auto=format&fit=crop&w=900&q=80",
    className: "",
  },
  {
    title: "Field Uniform",
    note: "Rigidez editorial, bolsos táticos e acabamento sem excesso visual.",
    image: "https://images.unsplash.com/photo-1507679799987-c73779587ccf?auto=format&fit=crop&w=900&q=80",
    className: "",
  },
];

export function EditorialGrid() {
  return (
    <section className="section-shell section-block">
      <div className="mb-8 flex flex-col gap-4 lg:flex-row lg:items-end lg:justify-between">
        <div className="space-y-3">
          <p className="section-number">04 / Editorial grid</p>
          <p className="eyebrow">Editorial grid / strong rhythm</p>
          <h2 className="section-title max-w-2xl">
            Image first. <span className="text-accent">Copy second.</span> Memory stays.
          </h2>
        </div>
        <p className="max-w-md font-mono text-[0.74rem] uppercase tracking-[0.28em] text-fog">
          Reference translation: strong branding, asymmetric composition, visual anchors and modular breaks.
        </p>
      </div>

      <div className="grid gap-5 md:grid-cols-3">
        {looks.map((look) => (
          <article key={look.title} className={`panel overflow-hidden rounded-[1.8rem] ${look.className}`}>
            <div className="h-[260px] w-full">
              <Image
                src={look.image}
                alt={look.title}
                width={900}
                height={900}
                className="h-full w-full object-cover"
              />
            </div>
            <div className="space-y-3 px-6 py-6">
              <p className="caption-label">Look study</p>
              <h3 className="text-2xl font-semibold uppercase tracking-[-0.04em]">{look.title}</h3>
              <p className="text-sm leading-relaxed text-white/70">{look.note}</p>
            </div>
          </article>
        ))}

        <div className="panel rounded-[1.8rem] p-6 md:col-span-3">
          <div className="grid gap-6 lg:grid-cols-[0.45fr_0.55fr]">
            <div>
              <p className="caption-label">Visual proof</p>
              <p className="mt-3 text-3xl font-semibold uppercase leading-tight tracking-[-0.04em]">
                Controlled color. Massive type. Outerwear with campaign gravity.
              </p>
            </div>
            <p className="text-base leading-relaxed text-white/70">
              Imagem vem antes da copy, mas os blocos tipográficos pesados sustentam a identidade – tradução direta das referências `typography-bold` e `editorial-identity`.
            </p>
          </div>
        </div>
      </div>
    </section>
  );
}
