import { CTAButton } from "@/components/ui/cta-button";

export function FinalCTA() {
  return (
    <section className="border-t border-white/[0.04] py-24">
      <div className="mx-auto max-w-3xl px-6 text-center">
        <div className="relative rounded-3xl border border-white/[0.06] bg-gradient-to-b from-white/[0.03] to-transparent p-12 sm:p-16">
          <div className="pointer-events-none absolute -top-px left-1/2 h-px w-3/4 -translate-x-1/2 bg-gradient-to-r from-transparent via-blue-500/50 to-transparent" />

          <h2 className="text-4xl font-semibold tracking-tight text-slate-200 sm:text-5xl">
            Ready to own your
            <br />
            infrastructure?
          </h2>
          <p className="mx-auto mt-4 max-w-lg text-lg text-slate-400">
            Join teams that stopped fighting their platform and started
            building with it.
          </p>
          <div className="mt-10 flex flex-col items-center justify-center gap-3 sm:flex-row">
            <CTAButton href="#">Start free trial</CTAButton>
            <CTAButton href="#" variant="secondary">
              Talk to sales
            </CTAButton>
          </div>
          <p className="mt-6 text-[13px] text-slate-500">
            Free for individual developers · Teams start at $29/mo
          </p>
        </div>
      </div>
    </section>
  );
}
