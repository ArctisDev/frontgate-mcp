export function FinalCTA() {
  return (
    <section id="cta" className="section-shell section-block pb-20">
      <div className="panel rounded-[2rem] p-8 md:p-10 xl:p-12">
        <div className="grid gap-8 lg:grid-cols-[1fr_0.55fr] lg:items-start">
          <div className="space-y-5">
            <p className="section-number">07 / Appointment</p>
            <p className="eyebrow">Private appointment</p>
            <h2 className="display-kicker text-[clamp(3rem,7vw,5.5rem)]">
              Enter the polar arc.
            </h2>
            <p className="body-lead max-w-2xl text-base">
              Sessões privadas em Reykjavik, São Paulo e Zurich apresentam campanha, corte e material
              em clima de galeria editorial.
            </p>
          </div>

          <div className="rounded-[1.6rem] border border-white/12 p-6">
            <div className="space-y-4">
              <div>
                <p className="caption-label">Formato</p>
                <p className="mt-2 text-lg font-semibold uppercase tracking-[-0.02em]">
                  Fitting · briefing · material archive
                </p>
              </div>
              <div className="flex flex-wrap gap-3 pt-2">
                <a href="mailto:manifest@polararc.studio" className="button-primary w-full sm:w-auto">
                  Agendar prova
                </a>
                <a href="mailto:manifest@polararc.studio" className="button-secondary w-full sm:w-auto">
                  Receber dossier
                </a>
              </div>
            </div>
          </div>
        </div>
      </div>
    </section>
  );
}
