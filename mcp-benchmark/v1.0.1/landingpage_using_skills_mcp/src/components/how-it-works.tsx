export function HowItWorks() {
  return (
    <section className="py-[var(--spacing-section)] bg-bg-base">
      <div className="section-container">
        <div className="max-w-2xl mb-16">
          <p className="section-label">How it works</p>
          <h2 className="section-title">
            From zero to production in minutes, not days.
          </h2>
        </div>

        <div className="grid md:grid-cols-3 gap-8 md:gap-12">
          {steps.map((step, index) => (
            <div key={step.title} className="relative">
              <div className="flex items-center gap-3 mb-5">
                <span className="flex items-center justify-center w-8 h-8 rounded-full bg-bg-overlay border border-border text-xs font-mono font-semibold text-accent">
                  {String(index + 1).padStart(2, "0")}
                </span>
                {index < steps.length - 1 && (
                  <div className="hidden md:block flex-1 h-px bg-gradient-to-r from-border to-transparent" />
                )}
              </div>
              <h3 className="text-base font-semibold text-text-primary mb-2">
                {step.title}
              </h3>
              <p className="text-sm text-text-secondary leading-relaxed">
                {step.description}
              </p>
            </div>
          ))}
        </div>

        {/* Terminal-style code block — decorative, contextual */}
        <div className="mt-16 rounded-xl bg-bg-raised border border-border/50 overflow-hidden">
          <div className="flex items-center gap-2 px-5 py-3 border-b border-border/50 bg-bg-overlay/30">
            <span className="w-3 h-3 rounded-full bg-red-500/50" />
            <span className="w-3 h-3 rounded-full bg-yellow-500/50" />
            <span className="w-3 h-3 rounded-full bg-green-500/50" />
            <span className="ml-3 text-xs text-text-tertiary font-mono">
              terminal
            </span>
          </div>
          <div className="p-6 font-mono text-sm leading-relaxed">
            <p className="text-text-tertiary">$ arctis init my-service</p>
            <p className="text-text-secondary mt-2">
              <span className="text-accent">→</span> Detected: Next.js 16, Node
              22
            </p>
            <p className="text-text-secondary">
              <span className="text-accent">→</span> Configuring infrastructure
              as code...
            </p>
            <p className="text-text-secondary">
              <span className="text-accent">→</span> Building container image...
            </p>
            <p className="mt-3 text-success font-medium">
              ✓ Deployed to us-east-1 — 38s
            </p>
            <p className="text-text-tertiary mt-1">
              https://my-service.arctis.dev
            </p>
          </div>
        </div>
      </div>
    </section>
  );
}

const steps = [
  {
    title: "Connect your repo",
    description:
      "Link your GitHub or GitLab repository. Arctis auto-detects your framework, runtime, and dependencies.",
  },
  {
    title: "Configure your infra",
    description:
      "Define regions, scaling policies, and env vars through code — not a dashboard with hidden defaults. Everything versioned.",
  },
  {
    title: "Push to deploy",
    description:
      "Every push triggers a build, runs your test suite, and deploys with automated rollbacks. Zero-downtime by default.",
  },
];
