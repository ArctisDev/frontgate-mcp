export function Differentiation() {
  return (
    <section className="py-[var(--spacing-section)] bg-bg-base">
      <div className="section-container">
        <div className="max-w-2xl mb-16">
          <p className="section-label">The difference</p>
          <h2 className="section-title">
            Not another managed platform. A better foundation.
          </h2>
        </div>

        <div className="overflow-x-auto -mx-6 px-6">
          <table className="w-full min-w-[640px] text-sm">
            <thead>
              <tr className="border-b border-border/50">
                <th className="text-left py-4 pr-6 text-text-tertiary font-medium">
                  Capability
                </th>
                <th className="text-left py-4 px-6 text-text-tertiary font-medium">
                  Typical PaaS
                </th>
                <th className="text-left py-4 pl-6 font-semibold text-accent">
                  Arctis Deploy
                </th>
              </tr>
            </thead>
            <tbody>
              {comparisons.map((row, index) => (
                <tr
                  key={row.capability}
                  className={
                    index < comparisons.length - 1
                      ? "border-b border-border-subtle"
                      : ""
                  }
                >
                  <td className="py-4 pr-6 text-text-secondary">
                    {row.capability}
                  </td>
                  <td className="py-4 px-6 text-text-tertiary">
                    {row.typical}
                  </td>
                  <td className="py-4 pl-6 text-text-primary font-medium">
                    {row.arctis}
                  </td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
      </div>
    </section>
  );
}

const comparisons = [
  {
    capability: "Infrastructure ownership",
    typical: "Abstracted away, locked in",
    arctis: "You own it, runs on open standards",
  },
  {
    capability: "Egress pricing",
    typical: "Unpredictable, often expensive",
    arctis: "Transparent, flat-rate tiers",
  },
  {
    capability: "Multi-cloud",
    typical: "Single provider, difficult to migrate",
    arctis: "AWS, GCP, Azure — deploy to any",
  },
  {
    capability: "Scaling model",
    typical: "Per-seat or per-app pricing",
    arctis: "Resource-based, linear cost scaling",
  },
  {
    capability: "CI/CD integration",
    typical: "Built-in only, limited customization",
    arctis: "BYO pipeline or use Arctis CI",
  },
  {
    capability: "Compliance",
    typical: "Add-on, extra cost",
    arctis: "SOC 2, GDPR, HIPAA included",
  },
];
