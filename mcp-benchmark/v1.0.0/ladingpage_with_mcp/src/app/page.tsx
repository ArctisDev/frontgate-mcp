import { Button } from "@/components/ui/button";
import { SectionHeader } from "@/components/landing/section-header";
import { FeatureCard } from "@/components/landing/feature-card";
import { CTAButton } from "@/components/landing/cta-button";
import {
  ArrowRight,
  GitBranch,
  Shield,
  Zap,
  Globe,
  Layers,
  Terminal,
  Lock,
  Unlock,
  ChevronRight,
  Container,
  Cloud,
  Settings,
} from "lucide-react";

const NAV_LINKS = [
  { href: "#diferenciais", label: "Diferenciais" },
  { href: "#como-funciona", label: "Como funciona" },
  { href: "#recursos", label: "Recursos" },
  { href: "#posicionamento", label: "Comparativo" },
];

const DIFERENCIAIS = [
  {
    icon: Unlock,
    title: "Zero lock-in",
    description:
      "Seus deploy rodam sobre infraestrutura que você controla. Containers padrão, configs portáveis, sem runtime proprietário.",
  },
  {
    icon: Settings,
    title: "Custo transparente",
    description:
      "Você vê exatamente onde cada centavo vai. Sem tier intermediário, sem cobrança por recurso que não usa.",
  },
  {
    icon: GitBranch,
    title: "Git-native workflow",
    description:
      "Push faz deploy. Branch vira ambiente. Rollback é um comando. O workflow que você já conhece, sem camadas extras.",
  },
];

const FLUXO = [
  {
    step: "01",
    title: "Conecte seu repositório",
    description:
      "GitHub, GitLab ou Bitbucket. Uma vez configurado, cada push cria um build automaticamente. Suporte a monorepos e workspaces.",
    icon: GitBranch,
  },
  {
    step: "02",
    title: "Configure seu ambiente",
    description:
      "Dockerfile, variáveis de ambiente e domínio. Tudo em um arquivo declarativo que versiona junto com seu código.",
    icon: Container,
  },
  {
    step: "03",
    title: "Deploy e monitore",
    description:
      "Deploy zero-downtime com health checks nativos. Logs em tempo real, métricas de recursos e alertas configuráveis.",
    icon: Cloud,
  },
];

const RECURSOS = [
  {
    icon: Zap,
    title: "Deploy em < 30s",
    description:
      "Builds incrementais com cache distribuído. Se mudou uma linha, rebuilda uma linha.",
  },
  {
    icon: Shield,
    title: "Isolamento por workload",
    description:
      "Cada serviço roda em container dedicado com recursos garantidos e rede isolada.",
  },
  {
    icon: Globe,
    title: "Multi-região nativo",
    description:
      "Escolha onde sua app roda. Deploy simultâneo em múltiplas regiões sem configuração extra.",
  },
  {
    icon: Terminal,
    title: "CLI completa",
    description:
      "Gerencie tudo via terminal. Logs, deploys, rollbacks, variáveis — tudo com autocompletar.",
  },
  {
    icon: Lock,
    title: "Secrets encriptados",
    description:
      "Variáveis sensíveis encriptadas em repouso. Integração com Vault, AWS KMS e GCP KMS.",
  },
  {
    icon: GitBranch,
    title: "Preview environments",
    description:
      "Cada PR ganha um ambiente efêmero com URL única. Teste antes de mergear.",
  },
];

const COMPARATIVO = [
  { label: "Infraestrutura", paas: "Abstraída e opaca", arctis: "Visível e configurável" },
  { label: "Custo", paas: "Tier fixo, cobrança por add-on", arctis: "Pay-per-use real, sem surpresa" },
  { label: "Portabilidade", paas: "Migrar = reescrever", arctis: "Docker padrão, exportável" },
  { label: "Performance", paas: "Runtime compartilhado", arctis: "Recursos dedicados por serviço" },
  { label: "Debug", paas: "Logs limitados, sem SSH", arctis: "Acesso completo, logs em tempo real" },
];

export default function Home() {
  return (
    <div className="flex flex-col flex-1">
      {/* Navigation */}
      <nav className="sticky top-0 z-50 border-b border-arctis-border bg-background/80 backdrop-blur-xl">
        <div className="mx-auto flex h-16 max-w-6xl items-center justify-between px-6">
          <div className="flex items-center gap-2">
            <div className="flex h-8 w-8 items-center justify-center rounded-lg bg-arctis-blue">
              <Layers className="h-4 w-4 text-primary-foreground" />
            </div>
            <span className="text-lg font-semibold tracking-tight">Arctis</span>
            <span className="text-sm font-medium text-arctis-text-muted">Deploy</span>
          </div>
          <div className="hidden items-center gap-8 md:flex">
            {NAV_LINKS.map((link) => (
              <a
                key={link.href}
                href={link.href}
                className="text-sm text-muted-foreground transition-colors hover:text-foreground"
              >
                {link.label}
              </a>
            ))}
          </div>
          <div className="flex items-center gap-3">
            <Button variant="ghost" size="sm" className="text-sm text-muted-foreground">
              Login
            </Button>
            <Button size="sm" className="bg-arctis-blue text-primary-foreground hover:bg-arctis-blue-light">
              Começar agora
              <ArrowRight className="ml-1 h-3.5 w-3.5" />
            </Button>
          </div>
        </div>
      </nav>

      <main className="flex flex-1 flex-col">
        {/* Hero */}
        <section className="relative overflow-hidden">
          <div className="absolute inset-0 bg-[radial-gradient(ellipse_80%_60%_at_50%_-20%,oklch(0.65_0.18_255_/_0.15),transparent)]" />
          <div className="relative mx-auto max-w-6xl px-6 pb-24 pt-28 md:pb-32 md:pt-36">
            <div className="mx-auto max-w-3xl text-center">
              <div className="mb-6 inline-flex items-center gap-2 rounded-full border border-arctis-border bg-arctis-surface-raised px-4 py-1.5 text-xs font-medium text-arctis-blue-light">
                <span className="h-1.5 w-1.5 rounded-full bg-arctis-blue" />
                Infraestrutura sob seu controle
              </div>
              <h1 className="mb-6 text-4xl font-bold leading-[1.1] tracking-tight md:text-5xl lg:text-6xl">
                Deploy sem intermediário.
                <br />
                <span className="text-arctis-blue">Infraestrutura que você entende.</span>
              </h1>
              <p className="mx-auto mb-10 max-w-xl text-lg leading-relaxed text-muted-foreground md:text-xl">
                Arctis é a plataforma de deploy que devolve o controle ao
                desenvolvedor. Sem vendor lock-in, sem custo opaco, sem
                abstrações que escondem o que importa.
              </p>
              <div className="flex flex-col items-center justify-center gap-3 sm:flex-row">
                <CTAButton icon={ArrowRight}>Começar gratuitamente</CTAButton>
                <CTAButton variant="outline" icon={Terminal}>Ver documentação</CTAButton>
              </div>
            </div>
          </div>
        </section>

        {/* Diferenciais Estratégicos */}
        <section id="diferenciais" className="border-t border-arctis-border bg-arctis-surface">
          <div className="mx-auto max-w-6xl px-6 py-24 md:py-32">
            <SectionHeader
              label="Por que Arctis"
              title="Decisões de engenharia, não de marketing."
              description="Cada escolha da plataforma foi feita para reduzir acoplamento, aumentar transparência e manter você no comando."
            />
            <div className="mt-16 grid gap-px md:grid-cols-3">
              {DIFERENCIAIS.map((item, i) => (
                <div
                  key={i}
                  className="group relative bg-background p-8 transition-colors hover:bg-arctis-surface-raised md:p-10"
                >
                  <div className="mb-5 inline-flex h-10 w-10 items-center justify-center rounded-lg border border-arctis-border bg-arctis-surface">
                    <item.icon className="h-5 w-5 text-arctis-blue" />
                  </div>
                  <h3 className="mb-2 text-lg font-semibold">{item.title}</h3>
                  <p className="text-sm leading-relaxed text-muted-foreground">{item.description}</p>
                </div>
              ))}
            </div>
          </div>
        </section>

        {/* Como funciona - Fluxo do produto */}
        <section id="como-funciona" className="border-t border-arctis-border">
          <div className="mx-auto max-w-6xl px-6 py-24 md:py-32">
            <SectionHeader
              label="Fluxo do produto"
              title="Do commit ao production em três passos."
              description="Sem painéis complexos. Sem configurações que exigem manual. Apenas o que importa para rodar sua aplicação."
              align="center"
              className="mx-auto mb-16"
            />
            <div className="mx-auto max-w-4xl">
              <div className="relative">
                <div className="absolute left-[27px] top-0 hidden h-full w-px bg-gradient-to-b from-arctis-blue/50 via-arctis-blue/20 to-transparent md:block" />
                <div className="space-y-12 md:space-y-16">
                  {FLUXO.map((item, i) => (
                    <div key={i} className="relative flex gap-6 md:gap-10">
                      <div className="flex h-14 w-14 shrink-0 items-center justify-center rounded-xl border border-arctis-border bg-arctis-surface">
                        <item.icon className="h-6 w-6 text-arctis-blue" />
                      </div>
                      <div className="pt-1">
                        <span className="mb-2 block font-mono text-xs font-medium text-arctis-blue-dim">{item.step}</span>
                        <h3 className="mb-2 text-xl font-semibold">{item.title}</h3>
                        <p className="max-w-md text-sm leading-relaxed text-muted-foreground">{item.description}</p>
                      </div>
                    </div>
                  ))}
                </div>
              </div>
            </div>
          </div>
        </section>

        {/* Recursos principais */}
        <section id="recursos" className="border-t border-arctis-border bg-arctis-surface">
          <div className="mx-auto max-w-6xl px-6 py-24 md:py-32">
            <SectionHeader
              label="Recursos"
              title="Construído para quem deploya todo dia."
            />
            <div className="mt-16 grid gap-4 md:grid-cols-2 lg:grid-cols-3">
              {RECURSOS.map((item, i) => (
                <FeatureCard
                  key={i}
                  icon={item.icon}
                  title={item.title}
                  description={item.description}
                />
              ))}
            </div>
          </div>
        </section>

        {/* Posicionamento contra soluções engessadas */}
        <section id="posicionamento" className="border-t border-arctis-border">
          <div className="mx-auto max-w-6xl px-6 py-24 md:py-32">
            <div className="grid gap-16 md:grid-cols-2 md:gap-20">
              <div>
                <p className="mb-3 text-xs font-semibold uppercase tracking-widest text-arctis-blue">Comparativo</p>
                <h2 className="mb-6 text-3xl font-bold tracking-tight md:text-4xl">
                  PaaS engessa.
                  <br />
                  Arctis libera.
                </h2>
                <p className="text-lg leading-relaxed text-muted-foreground">
                  Plataformas PaaS tradicionais simplificam no início, mas
                  travam quando você precisa de mais. O custo sobe, a
                  performance estagna e migrar vira projeto.
                </p>
                <p className="mt-4 text-lg leading-relaxed text-muted-foreground">
                  Arctis foi desenhado para crescer com sua aplicação — não
                  contra ela.
                </p>
              </div>
              <div className="space-y-1">
                {COMPARATIVO.map((row, i) => (
                  <div
                    key={i}
                    className="grid grid-cols-[120px_1fr_1fr] items-center gap-4 border-b border-arctis-border py-4 last:border-b-0"
                  >
                    <span className="text-xs font-semibold uppercase tracking-wider text-arctis-text-muted">{row.label}</span>
                    <span className="text-sm text-muted-foreground/70 line-through decoration-muted-foreground/30">{row.paas}</span>
                    <span className="flex items-center gap-1.5 text-sm font-medium text-foreground">
                      <ChevronRight className="h-3.5 w-3.5 text-arctis-blue" />
                      {row.arctis}
                    </span>
                  </div>
                ))}
              </div>
            </div>
          </div>
        </section>

        {/* CTA Final */}
        <section className="border-t border-arctis-border bg-arctis-surface">
          <div className="mx-auto max-w-6xl px-6 py-24 md:py-32">
            <div className="relative overflow-hidden rounded-2xl border border-arctis-border bg-background p-12 text-center md:p-20">
              <div className="absolute inset-0 bg-[radial-gradient(ellipse_60%_50%_at_50%_50%,oklch(0.65_0.18_255_/_0.08),transparent)]" />
              <div className="relative">
                <h2 className="mb-4 text-3xl font-bold tracking-tight md:text-4xl">
                  Pronto para deployar com controle?
                </h2>
                <p className="mx-auto mb-8 max-w-md text-lg text-muted-foreground">
                  Sem cartão de crédito. Sem compromisso. Deploy em menos de 5 minutos.
                </p>
                <div className="flex flex-col items-center justify-center gap-3 sm:flex-row">
                  <CTAButton icon={ArrowRight}>Criar conta gratuita</CTAButton>
                  <CTAButton variant="outline">Falar com vendas</CTAButton>
                </div>
              </div>
            </div>
          </div>
        </section>
      </main>

      {/* Footer */}
      <footer className="border-t border-arctis-border bg-background">
        <div className="mx-auto max-w-6xl px-6 py-12 md:py-16">
          <div className="grid gap-10 md:grid-cols-4">
            <div>
              <div className="mb-4 flex items-center gap-2">
                <div className="flex h-7 w-7 items-center justify-center rounded-lg bg-arctis-blue">
                  <Layers className="h-3.5 w-3.5 text-primary-foreground" />
                </div>
                <span className="text-sm font-semibold">Arctis Deploy</span>
              </div>
              <p className="text-sm leading-relaxed text-muted-foreground">
                Infraestrutura para desenvolvedores que valorizam controle, transparência e simplicidade.
              </p>
            </div>
            <div>
              <h4 className="mb-3 text-xs font-semibold uppercase tracking-wider text-muted-foreground">Produto</h4>
              <ul className="space-y-2">
                {["Deploy", "Ambientes", "Monitoramento", "CLI", "API"].map((item) => (
                  <li key={item}>
                    <a href="#" className="text-sm text-muted-foreground transition-colors hover:text-foreground">{item}</a>
                  </li>
                ))}
              </ul>
            </div>
            <div>
              <h4 className="mb-3 text-xs font-semibold uppercase tracking-wider text-muted-foreground">Recursos</h4>
              <ul className="space-y-2">
                {["Documentação", "Guias", "Status", "Changelog", "Blog"].map((item) => (
                  <li key={item}>
                    <a href="#" className="text-sm text-muted-foreground transition-colors hover:text-foreground">{item}</a>
                  </li>
                ))}
              </ul>
            </div>
            <div>
              <h4 className="mb-3 text-xs font-semibold uppercase tracking-wider text-muted-foreground">Empresa</h4>
              <ul className="space-y-2">
                {["Sobre", "Carreiras", "Contato", "Privacidade", "Termos"].map((item) => (
                  <li key={item}>
                    <a href="#" className="text-sm text-muted-foreground transition-colors hover:text-foreground">{item}</a>
                  </li>
                ))}
              </ul>
            </div>
          </div>
          <div className="mt-12 flex flex-col items-center justify-between gap-4 border-t border-arctis-border pt-8 md:flex-row">
            <p className="text-xs text-arctis-text-muted">
              &copy; {new Date().getFullYear()} Arctis Deploy. Todos os direitos reservados.
            </p>
            <div className="flex items-center gap-4">
              {["GitHub", "Twitter", "Discord"].map((item) => (
                <a key={item} href="#" className="text-xs text-arctis-text-muted transition-colors hover:text-foreground">{item}</a>
              ))}
            </div>
          </div>
        </div>
      </footer>
    </div>
  );
}
