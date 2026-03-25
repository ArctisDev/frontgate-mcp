import { CardBase } from "@/components/ui/card-base";

export function Features() {
  return (
    <section id="features" className="border-t border-white/[0.04] py-24">
      <div className="mx-auto max-w-6xl px-6">
        <div className="mb-16 text-center">
          <p className="mb-3 text-[12px] font-medium uppercase tracking-widest text-blue-400">
            Capabilities
          </p>
          <h2 className="text-4xl font-semibold tracking-tight text-slate-200 sm:text-5xl">
            Infrastructure primitives,
            <br />
            not abstractions
          </h2>
        </div>

        <div className="grid gap-4 md:grid-cols-2">
          {FEATURES.map((feature) => (
            <CardBase key={feature.title}>
              <div className="flex items-start gap-4">
                <div className="mt-0.5 flex h-9 w-9 shrink-0 items-center justify-center rounded-lg border border-white/[0.08] bg-white/[0.03]">
                  <feature.icon className="h-4 w-4 text-slate-400" />
                </div>
                <div>
                  <h3 className="mb-1.5 text-[15px] font-semibold text-slate-200">
                    {feature.title}
                  </h3>
                  <p className="text-[14px] leading-relaxed text-slate-400">
                    {feature.description}
                  </p>
                </div>
              </div>
            </CardBase>
          ))}
        </div>
      </div>
    </section>
  );
}

function GlobeIcon({ className }: { className?: string }) {
  return (
    <svg className={className} viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="1.5" strokeLinecap="round" strokeLinejoin="round">
      <circle cx="12" cy="12" r="10" />
      <path d="M2 12h20M12 2a15.3 15.3 0 0 1 4 10 15.3 15.3 0 0 1-4 10 15.3 15.3 0 0 1-4-10 15.3 15.3 0 0 1 4-10z" />
    </svg>
  );
}

function DatabaseIcon({ className }: { className?: string }) {
  return (
    <svg className={className} viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="1.5" strokeLinecap="round" strokeLinejoin="round">
      <ellipse cx="12" cy="5" rx="9" ry="3" />
      <path d="M21 12c0 1.66-4 3-9 3s-9-1.34-9-3M3 5v14c0 1.66 4 3 9 3s9-1.34 9-3V5" />
    </svg>
  );
}

function GitBranchIcon({ className }: { className?: string }) {
  return (
    <svg className={className} viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="1.5" strokeLinecap="round" strokeLinejoin="round">
      <line x1="6" y1="3" x2="6" y2="15" />
      <circle cx="18" cy="6" r="3" />
      <circle cx="6" cy="18" r="3" />
      <path d="M18 9a9 9 0 0 1-9 9" />
    </svg>
  );
}

function ActivityIcon({ className }: { className?: string }) {
  return (
    <svg className={className} viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="1.5" strokeLinecap="round" strokeLinejoin="round">
      <polyline points="22 12 18 12 15 21 9 3 6 12 2 12" />
    </svg>
  );
}

function KeyIcon({ className }: { className?: string }) {
  return (
    <svg className={className} viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="1.5" strokeLinecap="round" strokeLinejoin="round">
      <path d="m21 2-2 2m-7.61 7.61a5.5 5.5 0 1 1-7.778 7.778 5.5 5.5 0 0 1 7.777-7.777zm0 0L15.5 7.5m0 0 3 3L22 7l-3-3m-3.5 3.5L19 4" />
    </svg>
  );
}

function BoxIcon({ className }: { className?: string }) {
  return (
    <svg className={className} viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="1.5" strokeLinecap="round" strokeLinejoin="round">
      <path d="M21 16V8a2 2 0 0 0-1-1.73l-7-4a2 2 0 0 0-2 0l-7 4A2 2 0 0 0 3 8v8a2 2 0 0 0 1 1.73l7 4a2 2 0 0 0 2 0l7-4A2 2 0 0 0 21 16z" />
      <polyline points="3.27 6.96 12 12.01 20.73 6.96" />
      <line x1="12" y1="22.08" x2="12" y2="12" />
    </svg>
  );
}

function ZapIcon({ className }: { className?: string }) {
  return (
    <svg className={className} viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="1.5" strokeLinecap="round" strokeLinejoin="round">
      <polygon points="13 2 3 14 12 14 11 22 21 10 12 10 13 2" />
    </svg>
  );
}

function RefreshCwIcon({ className }: { className?: string }) {
  return (
    <svg className={className} viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="1.5" strokeLinecap="round" strokeLinejoin="round">
      <polyline points="23 4 23 10 17 10" />
      <polyline points="1 20 1 14 7 14" />
      <path d="M3.51 9a9 9 0 0 1 14.85-3.36L23 10M1 14l4.64 4.36A9 9 0 0 0 20.49 15" />
    </svg>
  );
}

const FEATURES = [
  {
    icon: GlobeIcon,
    title: "Edge network included",
    description:
      "Global CDN with automatic routing, DDoS protection, and SSL termination. No separate CDN to configure.",
  },
  {
    icon: DatabaseIcon,
    title: "Managed databases",
    description:
      "Postgres, Redis, and object storage provisioned alongside your app. Automated backups, point-in-time recovery.",
  },
  {
    icon: GitBranchIcon,
    title: "Preview environments",
    description:
      "Every pull request gets its own isolated environment with the full infra stack. Review in production-like conditions.",
  },
  {
    icon: ActivityIcon,
    title: "Built-in observability",
    description:
      "Metrics, logs, traces — unified. No third-party APM needed. Search across all your services from one place.",
  },
  {
    icon: KeyIcon,
    title: "Secret management",
    description:
      "Encrypted at rest, injected at runtime. Rotate secrets without redeploying. Environment-scoped by default.",
  },
  {
    icon: BoxIcon,
    title: "Container & serverless",
    description:
      "Run long-lived containers or auto-scaling functions. Switch between models without changing your application code.",
  },
  {
    icon: ZapIcon,
    title: "Instant rollbacks",
    description:
      "One click to roll back to any previous deployment. Your infra state rolls back with it. Zero downtime.",
  },
  {
    icon: RefreshCwIcon,
    title: "Drift detection",
    description:
      "We detect and alert on infrastructure drift automatically. Your declared state always matches reality.",
  },
];
