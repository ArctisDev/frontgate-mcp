import { CTAButton } from "@/components/ui/cta-button";

export function Hero() {
  return (
    <section className="relative flex min-h-[92vh] items-center justify-center overflow-hidden pt-16">
      <div className="pointer-events-none absolute inset-0">
        <div className="absolute left-1/2 top-0 h-[600px] w-[800px] -translate-x-1/2 -translate-y-1/4 rounded-full bg-blue-500/[0.04] blur-[120px]" />
        <div className="absolute bottom-0 left-1/4 h-[400px] w-[600px] rounded-full bg-blue-400/[0.03] blur-[100px]" />
      </div>

      <div className="relative mx-auto max-w-4xl px-6 text-center">
        <div className="animate-fade-in-up mb-8 inline-flex items-center gap-2 rounded-full border border-white/[0.08] bg-white/[0.04] px-4 py-1.5">
          <span className="h-1.5 w-1.5 rounded-full bg-emerald-400" />
          <span className="text-[12px] font-medium text-slate-400">
            All systems operational — v2.4 released
          </span>
        </div>

        <h1 className="animate-fade-in-up-delay-1 text-balance text-5xl font-semibold leading-[1.1] tracking-tight text-slate-200 sm:text-6xl md:text-7xl">
          Deploy with full
          <br />
          <span className="text-gradient">control of infrastructure</span>
        </h1>

        <p className="animate-fade-in-up-delay-2 mx-auto mt-6 max-w-2xl text-balance text-lg leading-relaxed text-slate-400 sm:text-xl">
          Own your stack. Scale without lock-in. Arctis gives your team the
          infrastructure primitives you need — without the platform tax.
        </p>

        <div className="animate-fade-in-up-delay-3 mt-10 flex flex-col items-center justify-center gap-3 sm:flex-row">
          <CTAButton href="#">Start free trial</CTAButton>
          <CTAButton href="#" variant="secondary">
            View documentation
          </CTAButton>
        </div>

        <p className="mt-8 text-[13px] text-slate-500">
          No credit card required · 14-day free trial · Deploy in minutes
        </p>
      </div>
    </section>
  );
}
