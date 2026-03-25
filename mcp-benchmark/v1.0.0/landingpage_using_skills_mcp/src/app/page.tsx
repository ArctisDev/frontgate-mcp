import { Navigation } from "@/components/navigation";
import { Hero } from "@/components/hero";
import { Credibility } from "@/components/credibility";
import { Benefits } from "@/components/benefits";
import { HowItWorks } from "@/components/how-it-works";
import { Features } from "@/components/features";
import { Differentiation } from "@/components/differentiation";
import { FinalCTA } from "@/components/final-cta";
import { Footer } from "@/components/footer";

export default function Home() {
  return (
    <>
      <Navigation />
      <main className="flex-1">
        <Hero />
        <Credibility />
        <Benefits />
        <HowItWorks />
        <Features />
        <Differentiation />
        <FinalCTA />
      </main>
      <Footer />
    </>
  );
}
