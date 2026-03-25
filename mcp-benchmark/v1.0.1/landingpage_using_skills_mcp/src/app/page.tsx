import { Nav } from "@/components/nav";
import { Hero } from "@/components/hero";
import { SocialProof } from "@/components/social-proof";
import { Benefits } from "@/components/benefits";
import { HowItWorks } from "@/components/how-it-works";
import { Features } from "@/components/features";
import { Differentiation } from "@/components/differentiation";
import { FinalCTA } from "@/components/final-cta";
import { Footer } from "@/components/footer";

export default function Home() {
  return (
    <>
      <Nav />
      <main className="flex-1">
        <Hero />
        <SocialProof />
        <Benefits />
        <HowItWorks />
        <Features />
        <Differentiation />
        <div className="gradient-divider" />
        <FinalCTA />
      </main>
      <Footer />
    </>
  );
}
