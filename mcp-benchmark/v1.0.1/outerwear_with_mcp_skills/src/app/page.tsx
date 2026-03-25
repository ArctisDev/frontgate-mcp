import { Header } from "@/components/header";
import { Hero } from "@/components/hero";
import { CollectionShowcase } from "@/components/collection";
import { MaterialSection } from "@/components/materials";
import { EditorialGrid } from "@/components/editorial-grid";
import { Manifesto } from "@/components/manifesto";
import { ProofSection } from "@/components/proof";
import { FinalCTA } from "@/components/final-cta";
import { Footer } from "@/components/footer";

export default function Home() {
  return (
    <div className="page-shell">
      <Header />
      <main>
        <Hero />
        <CollectionShowcase />
        <MaterialSection />
        <EditorialGrid />
        <Manifesto />
        <ProofSection />
        <FinalCTA />
      </main>
      <Footer />
    </div>
  );
}
