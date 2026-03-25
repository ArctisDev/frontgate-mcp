export function Footer() {
  return (
    <footer className="section-shell border-t border-white/8 py-8">
      <div className="flex flex-col gap-3 md:flex-row md:items-center md:justify-between">
        <p className="font-mono text-xs uppercase tracking-[0.3em] text-fog">
          POLAR//ARC / outerwear for severe climates / {new Date().getFullYear()}
        </p>
        <p className="font-mono text-xs uppercase tracking-[0.3em] text-fog">
          manifest@polararc.studio / showroom by appointment
        </p>
      </div>
    </footer>
  );
}
