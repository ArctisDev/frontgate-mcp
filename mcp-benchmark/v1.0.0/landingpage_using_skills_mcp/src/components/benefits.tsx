import { CardBase } from "@/components/ui/card-base";

export function Benefits() {
  return (
    <section id="benefits" className="py-24">
      <div className="mx-auto max-w-6xl px-6">
        <div className="mb-16 max-w-2xl">
          <p className="mb-3 text-[12px] font-medium uppercase tracking-widest text-blue-400">
            Why Arctis
          </p>
          <h2 className="text-4xl font-semibold tracking-tight text-slate-200 sm:text-5xl">
            Built for teams that
            <br />
            outgrow platforms
          </h2>
          <p className="mt-4 text-lg text-slate-400">
            Stop paying the &quot;easy now, painful later&quot; tax. Arctis gives you
            real control from day one.
          </p>
        </div>

        <div className="grid gap-6 md:grid-cols-3">
          {BENEFITS.map((benefit) => (
            <CardBase key={benefit.title} className="p-8">
              <div className="mb-5 flex h-11 w-11 items-center justify-center rounded-xl bg-blue-500/10">
                <benefit.icon className="h-5 w-5 text-blue-400" />
              </div>
              <h3 className="mb-2 text-lg font-semibold text-slate-200">
                {benefit.title}
              </h3>
              <p className="text-[15px] leading-relaxed text-slate-400">
                {benefit.description}
              </p>
            </CardBase>
          ))}
        </div>
      </div>
    </section>
  );
}

function LockIcon({ className }: { className?: string }) {
  return (
    <svg className={className} viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="1.5" strokeLinecap="round" strokeLinejoin="round">
      <rect x="3" y="11" width="18" height="11" rx="2" ry="2" />
      <path d="M7 11V7a5 5 0 0 1 10 0v4" />
    </svg>
  );
}

function ScaleIcon({ className }: { className?: string }) {
  return (
    <svg className={className} viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="1.5" strokeLinecap="round" strokeLinejoin="round">
      <path d="M12 2v20M2 12h20" />
      <path d="m4.93 4.93 14.14 14.14M19.07 4.93 4.93 19.07" />
      <circle cx="12" cy="12" r="3" />
    </svg>
  );
}

function ShieldIcon({ className }: { className?: string }) {
  return (
    <svg className={className} viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="1.5" strokeLinecap="round" strokeLinejoin="round">
      <path d="M12 22s8-4 8-10V5l-8-3-8 3v7c0 6 8 10 8 10z" />
      <path d="m9 12 2 2 4-4" />
    </svg>
  );
}

const BENEFITS = [
  {
    icon: LockIcon,
    title: "No vendor lock-in",
    description:
      "Your infrastructure definitions are portable. Export your configs anytime, run anywhere. You own your stack — not the other way around.",
  },
  {
    icon: ScaleIcon,
    title: "Costs that scale with you",
    description:
      "Transparent pricing based on actual resource usage. No hidden egress fees, no surprise bills. Predictable costs from your first deploy to your millionth.",
  },
  {
    icon: ShieldIcon,
    title: "Enterprise-grade security",
    description:
      "SOC 2 Type II certified. End-to-end encryption, RBAC, audit logs, and secret management built in — not bolted on as an afterthought.",
  },
];
