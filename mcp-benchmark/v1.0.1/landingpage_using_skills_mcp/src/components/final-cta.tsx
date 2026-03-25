import Link from "next/link";

export function FinalCTA() {
  return (
    <section className="py-[var(--spacing-section)] relative overflow-hidden">
      {/* Ambient glow */}
      <div
        className="absolute top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2 w-[600px] h-[400px] pointer-events-none"
        style={{
          background:
            "radial-gradient(ellipse at center, rgba(59,130,246,0.05) 0%, transparent 70%)",
        }}
      />

      <div className="section-container relative text-center">
        <h2 className="text-3xl md:text-4xl lg:text-5xl font-bold tracking-tight text-text-primary max-w-2xl mx-auto leading-tight">
          Stop renting your infrastructure.
          <br />
          <span className="text-text-secondary">Start owning it.</span>
        </h2>

        <p className="mt-5 text-lg text-text-secondary max-w-xl mx-auto leading-relaxed">
          Migrate in an afternoon, not a quarter. Arctis handles the hard parts
          — your team stays focused on shipping.
        </p>

        <div className="mt-10 flex flex-col sm:flex-row items-center justify-center gap-4">
          <Link
            href="#start"
            className="inline-flex h-12 items-center justify-center rounded-lg bg-accent px-8 text-sm font-semibold text-white hover:bg-accent-hover transition-colors duration-150"
          >
            Start free — no credit card
          </Link>
          <Link
            href="#contact"
            className="inline-flex h-12 items-center justify-center text-sm font-medium text-text-secondary hover:text-text-primary transition-colors duration-150"
          >
            Talk to sales
            <svg
              className="ml-1.5 w-4 h-4"
              viewBox="0 0 24 24"
              fill="none"
              stroke="currentColor"
              strokeWidth={2}
              strokeLinecap="round"
              strokeLinejoin="round"
            >
              <path d="M5 12h14M12 5l7 7-7 7" />
            </svg>
          </Link>
        </div>
      </div>
    </section>
  );
}
