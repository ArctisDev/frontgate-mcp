export function HowItWorks() {
  return (
    <section
      id="how-it-works"
      className="border-t border-white/[0.04] py-24"
    >
      <div className="mx-auto max-w-6xl px-6">
        <div className="mb-16 text-center">
          <p className="mb-3 text-[12px] font-medium uppercase tracking-widest text-blue-400">
            Workflow
          </p>
          <h2 className="text-4xl font-semibold tracking-tight text-slate-200 sm:text-5xl">
            From code to production
          </h2>
          <p className="mx-auto mt-4 max-w-xl text-lg text-slate-400">
            Three steps. No yak shaving. No infrastructure meetings.
          </p>
        </div>

        <div className="grid gap-12 md:grid-cols-3">
          {STEPS.map((step, i) => (
            <div key={step.title} className="relative">
              {/* Connector line */}
              {i < STEPS.length - 1 && (
                <div className="absolute right-0 top-8 hidden h-px w-full translate-x-1/2 bg-gradient-to-r from-white/[0.08] to-transparent md:block" />
              )}

              <div className="relative">
                <div className="mb-5 inline-flex h-10 w-10 items-center justify-center rounded-xl border border-white/[0.08] bg-white/[0.03] text-[14px] font-semibold text-slate-400">
                  {i + 1}
                </div>
                <h3 className="mb-2 text-lg font-semibold text-slate-200">
                  {step.title}
                </h3>
                <p className="text-[15px] leading-relaxed text-slate-400">
                  {step.description}
                </p>
              </div>
            </div>
          ))}
        </div>

        {/* Code example */}
        <div className="mt-16 overflow-hidden rounded-2xl border border-white/[0.06] bg-navy-900/50">
          <div className="flex items-center gap-2 border-b border-white/[0.06] px-5 py-3">
            <div className="h-2.5 w-2.5 rounded-full bg-white/[0.08]" />
            <div className="h-2.5 w-2.5 rounded-full bg-white/[0.08]" />
            <div className="h-2.5 w-2.5 rounded-full bg-white/[0.08]" />
            <span className="ml-3 text-[12px] text-slate-500">arctis.config.ts</span>
          </div>
          <pre className="overflow-x-auto p-6 text-[13px] leading-relaxed">
            <code className="font-mono">
              <span className="text-slate-500">{"import { defineConfig } from '@arctis/deploy';\n\n"}</span>
              <span className="text-blue-300">{"export default "}</span>
              <span className="text-blue-400">{"defineConfig"}</span>
              <span className="text-slate-300">{"({\n"}</span>
              <span className="text-slate-300">{"  "}</span>
              <span className="text-slate-400">{"region"}</span>
              <span className="text-slate-300">{"   "}</span>
              <span className="text-slate-500">{"// auto-detected from nearest edge\n"}</span>
              <span className="text-slate-300">{"  "}</span>
              <span className="text-slate-400">{"scaling"}</span>
              <span className="text-slate-300">{"  "}</span>
              <span className="text-emerald-400">{"'auto'"}</span>
              <span className="text-slate-300">{",\n"}</span>
              <span className="text-slate-300">{"  "}</span>
              <span className="text-slate-400">{"preview"}</span>
              <span className="text-slate-300">{"  "}</span>
              <span className="text-blue-400">{"true"}</span>
              <span className="text-slate-300">{",\n"}</span>
              <span className="text-slate-300">{"})"}</span>
            </code>
          </pre>
        </div>
      </div>
    </section>
  );
}

const STEPS = [
  {
    title: "Connect your repo",
    description:
      "Link your GitHub, GitLab, or Bitbucket repository. Arctis detects your framework and configures the build pipeline automatically.",
  },
  {
    title: "Define your infra",
    description:
      "Use our config DSL or bring your own Terraform/Pulumi. Define databases, queues, storage — whatever your app needs.",
  },
  {
    title: "Push to deploy",
    description:
      "Every push triggers a preview deployment. Merge to main goes to production. Rollbacks are one click.",
  },
];
