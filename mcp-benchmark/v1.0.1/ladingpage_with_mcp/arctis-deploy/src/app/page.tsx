import {
  ArrowRight,
  Shield,
  Zap,
  GitBranch,
  Layers,
  Lock,
  BarChart3,
  Cloud,
  Terminal,
  Check,
} from "lucide-react";
import { Badge } from "@/components/ui/badge";
import { Button } from "@/components/ui/button";
import { Separator } from "@/components/ui/separator";

const differentiators = [
  {
    icon: Cloud,
    title: "Multi-cloud sem abstrações vazias",
    description:
      "Deploy em AWS, GCP, Azure ou infraestrutura própria. O Arctis orquestra — os recursos são seus. Sem camadas mágicas que escondem o que está acontecendo.",
  },
  {
    icon: Lock,
    title: "Zero lock-in de runtime",
    description:
      "Seus containers rodam em ambientes padrão. Se você sair do Arctis amanhã, sua aplicação continua funcionando em qualquer lugar. Nenhuma dependência de SDK proprietário.",
  },
  {
    icon: BarChart3,
    title: "Custo proporcional ao consumo",
    description:
      "Pague pelo que usa, não por promessas de escala. Transparência total no billing, sem surpresas no fim do mês. Relatórios detalhados por projeto, ambiente e recurso.",
  },
  {
    icon: Shield,
    title: "Segurança por design, não por marketing",
    description:
      "Secrets management nativo, RBAC granular, audit log completo e integração com seu provedor de identidade. Segurança embutida, não vendida como add-on.",
  },
];

const flowSteps = [
  {
    step: "01",
    title: "Conecte seu repositório",
    description:
      "Integração nativa com GitHub, GitLab e Bitbucket. O Arctis detecta o stack, configura o build e sugere a infraestrutura ideal.",
    command: "git push origin main",
  },
  {
    step: "02",
    title: "Configure sua infraestrutura",
    description:
      "Escolha a cloud, a região e o tier de recursos. Tudo configurável via UI ou código. Exporte como Terraform se quiser.",
    command: "arctis infra plan",
  },
  {
    step: "03",
    title: "Deploy automático",
    description:
      "Cada push gera um preview. Cada merge em main faz deploy em produção. Rollback com um clique se algo der errado.",
    command: "arctis deploy --prod",
  },
];

const features = [
  {
    icon: GitBranch,
    title: "Preview por branch",
    description:
      "Cada branch gera um ambiente isolado com URL única. Teste features antes de merge. Compartilhe com QA e stakeholders.",
  },
  {
    icon: Zap,
    title: "Deploy incremental",
    description:
      "Só o que mudou é enviado. Builds mais rápidos, menor tempo de deploy, menos cobrança de transferência.",
  },
  {
    icon: Terminal,
    title: "CLI completa",
    description:
      "Gerencie tudo via terminal. Deploys, logs, secrets, variáveis de ambiente, escalonamento — tudo com autocomplete e documentação integrada.",
  },
  {
    icon: Shield,
    title: "Observabilidade nativa",
    description:
      "Métricas, logs e traces sem integração externa. Alertas configuráveis. Dashboard unificado por projeto e ambiente.",
  },
  {
    icon: Cloud,
    title: "Edge e serverless",
    description:
      "Funções serverless, CDN global e edge computing como parte do mesmo stack. Sem configuração separada, sem contas adicionais.",
  },
  {
    icon: Lock,
    title: "Compliance e governança",
    description:
      "SOC 2, LGPD, HIPAA-ready. Regiões de dados configuráveis. Controle granular de acesso por equipe, projeto e ambiente.",
  },
];

const comparisons = [
  { label: "Vendor lock-in", arctis: "Portabilidade total", other: "Migração custosa" },
  { label: "Modelo de cobrança", arctis: "Pay-per-use real", other: "Planos fixos + overage" },
  { label: "Infraestrutura", arctis: "Multi-cloud nativo", other: "Cloud proprietária" },
  { label: "Configuração", arctis: "IaC exportável", other: "YAML proprietário" },
  { label: "Observabilidade", arctis: "Nativa e aberta", other: "Add-on pago" },
];

const footerLinks = {
  produto: ["Recursos", "Preços", "Changelog", "Roadmap"],
  desenvolvedores: ["Documentação", "API Reference", "CLI", "Status"],
  empresa: ["Sobre", "Blog", "Carreiras", "Contato"],
};

export default function Home() {
  return (
    <div className="dark min-h-screen bg-background text-foreground">
      {/* ─── NAV ─────────────────────────────────────────────────── */}
      <nav className="fixed inset-x-0 top-0 z-50 border-b border-arctis-border bg-background/80 backdrop-blur-xl">
        <div className="mx-auto flex h-16 max-w-7xl items-center justify-between px-6 lg:px-8">
          <div className="flex items-center gap-3">
            <div className="flex h-8 w-8 items-center justify-center rounded-lg bg-primary">
              <Layers className="h-4 w-4 text-primary-foreground" />
            </div>
            <span className="font-heading text-lg font-semibold tracking-tight">
              Arctis
            </span>
          </div>
          <div className="hidden items-center gap-8 md:flex">
            {[
              { href: "#diferenciais", label: "Diferenciais" },
              { href: "#fluxo", label: "Como funciona" },
              { href: "#recursos", label: "Recursos" },
            ].map((link) => (
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
            <Button variant="ghost" size="sm" className="hidden sm:inline-flex">
              Documentação
            </Button>
            <Button size="sm">
              Começar agora
              <ArrowRight className="ml-1.5 h-3.5 w-3.5" />
            </Button>
          </div>
        </div>
      </nav>

      <main>
        {/* ─── HERO ──────────────────────────────────────────────── */}
        <section className="relative flex min-h-screen items-center justify-center overflow-hidden">
          <div className="pointer-events-none absolute inset-0 bg-grid-subtle opacity-40" />
          <div className="pointer-events-none absolute left-1/2 top-1/2 h-[800px] w-[800px] -translate-x-1/2 -translate-y-1/2 rounded-full bg-glow-radial blur-3xl" />

          <div className="relative z-10 mx-auto max-w-5xl px-6 py-32 text-center lg:px-8">
            <Badge
              variant="outline"
              className="mb-8 border-primary/20 bg-primary/5 text-primary"
            >
              <span className="mr-1.5 inline-block h-1.5 w-1.5 rounded-full bg-primary animate-pulse" />
              Beta aberta — acesso antecipado disponível
            </Badge>

            <h1 className="font-heading text-5xl font-semibold leading-[1.05] tracking-tight sm:text-6xl md:text-7xl lg:text-[5.5rem]">
              Deploy com{" "}
              <span className="text-gradient-blue">controle total</span>
              <br />
              sobre sua infraestrutura
            </h1>

            <p className="mx-auto mt-8 max-w-2xl text-lg leading-relaxed text-muted-foreground md:text-xl">
              Arctis Deploy é a plataforma para times que não aceitam lock-in.
              Faça deploy em qualquer cloud, mantenha posse dos seus recursos e
              otimize custos sem abrir mão de performance.
            </p>

            <div className="mt-12 flex flex-col items-center justify-center gap-4 sm:flex-row">
              <Button size="lg" className="h-12 px-8 text-base">
                Criar conta gratuita
                <ArrowRight className="ml-2 h-4 w-4" />
              </Button>
              <Button
                variant="outline"
                size="lg"
                className="h-12 border-arctis-border px-8 text-base"
              >
                Ver documentação
              </Button>
            </div>

            <div className="mt-16 flex flex-wrap items-center justify-center gap-x-8 gap-y-3 text-sm text-arctis-text-dim">
              {["Sem vendor lock-in", "Cobrança por uso real", "Multi-cloud nativo"].map((text) => (
                <span key={text} className="flex items-center gap-2">
                  <Check className="h-3.5 w-3.5 text-primary" />
                  {text}
                </span>
              ))}
            </div>
          </div>

          <div className="pointer-events-none absolute bottom-0 left-0 right-0 h-32 bg-gradient-to-t from-background to-transparent" />
        </section>

        {/* ─── DIFERENCIAIS ──────────────────────────────────────── */}
        <section id="diferenciais" className="relative py-32">
          <div className="mx-auto max-w-7xl px-6 lg:px-8">
            <div className="grid gap-16 lg:grid-cols-12 lg:gap-24">
              <div className="lg:col-span-5 lg:sticky lg:top-32 lg:self-start">
                <p className="text-sm font-medium uppercase tracking-widest text-primary">
                  Por que Arctis
                </p>
                <h2 className="mt-4 font-heading text-4xl font-semibold leading-tight tracking-tight sm:text-5xl">
                  Não entregue
                  <br />
                  sua infraestrutura
                  <br />
                  para ninguém
                </h2>
                <p className="mt-6 text-lg leading-relaxed text-muted-foreground">
                  A maioria das plataformas de deploy cobra sua liberdade como
                  preço de conveniência. Nós achamos que é possível ter os dois.
                </p>
              </div>

              <div className="lg:col-span-7">
                {differentiators.map((item, index) => (
                  <div
                    key={index}
                    className="group relative flex gap-5 border-b border-arctis-border py-8 last:border-b-0"
                  >
                    <div className="flex h-10 w-10 shrink-0 items-center justify-center rounded-lg bg-primary/10 text-primary transition-colors group-hover:bg-primary/20">
                      <item.icon className="h-5 w-5" />
                    </div>
                    <div>
                      <h3 className="text-lg font-semibold">{item.title}</h3>
                      <p className="mt-2 leading-relaxed text-muted-foreground">
                        {item.description}
                      </p>
                    </div>
                  </div>
                ))}
              </div>
            </div>
          </div>
        </section>

        {/* ─── FLUXO ─────────────────────────────────────────────── */}
        <section
          id="fluxo"
          className="relative border-y border-arctis-border bg-arctis-surface py-32"
        >
          <div className="mx-auto max-w-7xl px-6 lg:px-8">
            <div className="mx-auto max-w-3xl text-center">
              <p className="text-sm font-medium uppercase tracking-widest text-primary">
                Fluxo do produto
              </p>
              <h2 className="mt-4 font-heading text-4xl font-semibold leading-tight tracking-tight sm:text-5xl">
                Do commit ao production
                <br />
                em três etapas
              </h2>
              <p className="mt-6 text-lg text-muted-foreground">
                Sem pipelines complicados. Sem YAML interminável. O Arctis
                entende seu projeto e faz o resto.
              </p>
            </div>

            <div className="mt-20 grid gap-1 md:grid-cols-3">
              {flowSteps.map((item, index) => (
                <div
                  key={index}
                  className="relative border border-arctis-border bg-background p-8 first:rounded-l-xl last:rounded-r-xl md:border-r-0 last:md:border-r"
                >
                  <span className="font-mono text-xs font-medium text-primary">
                    {item.step}
                  </span>
                  <h3 className="mt-4 text-xl font-semibold">{item.title}</h3>
                  <p className="mt-3 text-sm leading-relaxed text-muted-foreground">
                    {item.description}
                  </p>
                  <div className="mt-6 rounded-lg bg-arctis-surface px-4 py-3">
                    <code className="font-mono text-sm text-primary">
                      {item.command}
                    </code>
                  </div>
                </div>
              ))}
            </div>
          </div>
        </section>

        {/* ─── RECURSOS ──────────────────────────────────────────── */}
        <section id="recursos" className="relative py-32">
          <div className="mx-auto max-w-7xl px-6 lg:px-8">
            <div className="mx-auto max-w-3xl text-center">
              <p className="text-sm font-medium uppercase tracking-widest text-primary">
                Recursos
              </p>
              <h2 className="mt-4 font-heading text-4xl font-semibold leading-tight tracking-tight sm:text-5xl">
                Tudo que você precisa,
                <br />
                nada que você não usa
              </h2>
            </div>

            <div className="mt-20 grid gap-px md:grid-cols-2 lg:grid-cols-3">
              {features.map((feature, index) => (
                <div
                  key={index}
                  className="group relative bg-background p-8 transition-colors hover:bg-arctis-surface"
                >
                  <div className="flex h-10 w-10 items-center justify-center rounded-lg bg-primary/10 text-primary transition-colors group-hover:bg-primary/15">
                    <feature.icon className="h-5 w-5" />
                  </div>
                  <h3 className="mt-5 text-lg font-semibold">
                    {feature.title}
                  </h3>
                  <p className="mt-2 text-sm leading-relaxed text-muted-foreground">
                    {feature.description}
                  </p>
                </div>
              ))}
            </div>
          </div>
        </section>

        {/* ─── COMPARATIVO ───────────────────────────────────────── */}
        <section className="relative overflow-hidden border-y border-arctis-border bg-arctis-surface py-32">
          <div className="pointer-events-none absolute left-1/2 top-0 h-px w-1/2 -translate-x-1/2 bg-gradient-to-r from-transparent via-primary/30 to-transparent" />

          <div className="mx-auto max-w-7xl px-6 lg:px-8">
            <div className="grid gap-16 lg:grid-cols-2 lg:gap-24">
              <div>
                <p className="text-sm font-medium uppercase tracking-widest text-primary">
                  Comparativo
                </p>
                <h2 className="mt-4 font-heading text-4xl font-semibold leading-tight tracking-tight sm:text-5xl">
                  Soluções engessadas
                  <br />
                  custam mais
                  <br />
                  do que parece
                </h2>
                <p className="mt-6 text-lg leading-relaxed text-muted-foreground">
                  Plataformas que simplificam demais escondem a complexidade — e
                  a cobram quando você precisa de controle. O Arctis é
                  transparente por padrão.
                </p>
              </div>

              <div>
                {comparisons.map((item, index) => (
                  <div
                    key={index}
                    className="grid grid-cols-3 items-center gap-4 border-b border-arctis-border py-5 last:border-b-0"
                  >
                    <span className="text-sm font-medium text-muted-foreground">
                      {item.label}
                    </span>
                    <div className="flex items-center gap-2">
                      <Check className="h-3.5 w-3.5 text-primary" />
                      <span className="text-sm font-medium">{item.arctis}</span>
                    </div>
                    <div className="flex items-center gap-2 text-arctis-text-dim">
                      <span className="text-xs">vs</span>
                      <span className="text-sm">{item.other}</span>
                    </div>
                  </div>
                ))}
              </div>
            </div>
          </div>
        </section>

        {/* ─── CTA FINAL ─────────────────────────────────────────── */}
        <section className="relative py-32">
          <div className="pointer-events-none absolute left-1/2 top-1/2 h-[600px] w-[600px] -translate-x-1/2 -translate-y-1/2 rounded-full bg-glow-radial blur-3xl" />

          <div className="relative z-10 mx-auto max-w-3xl px-6 text-center lg:px-8">
            <h2 className="font-heading text-4xl font-semibold leading-tight tracking-tight sm:text-5xl md:text-6xl">
              Pare de pagar por
              <br />
              <span className="text-gradient-blue">conveniência que aprisiona</span>
            </h2>
            <p className="mx-auto mt-6 max-w-xl text-lg text-muted-foreground">
              Crie sua conta, conecte um repositório e faça seu primeiro deploy
              em menos de 5 minutos. Sem cartão de crédito.
            </p>
            <div className="mt-10 flex flex-col items-center justify-center gap-4 sm:flex-row">
              <Button size="lg" className="h-12 px-10 text-base">
                Criar conta gratuita
                <ArrowRight className="ml-2 h-4 w-4" />
              </Button>
              <Button
                variant="outline"
                size="lg"
                className="h-12 border-arctis-border px-10 text-base"
              >
                Falar com vendas
              </Button>
            </div>
          </div>
        </section>
      </main>

      {/* ─── FOOTER ──────────────────────────────────────────────── */}
      <footer className="border-t border-arctis-border bg-arctis-surface">
        <div className="mx-auto max-w-7xl px-6 py-16 lg:px-8">
          <div className="grid gap-12 md:grid-cols-4">
            <div>
              <div className="flex items-center gap-3">
                <div className="flex h-8 w-8 items-center justify-center rounded-lg bg-primary">
                  <Layers className="h-4 w-4 text-primary-foreground" />
                </div>
                <span className="font-heading text-lg font-semibold tracking-tight">
                  Arctis
                </span>
              </div>
              <p className="mt-4 text-sm leading-relaxed text-muted-foreground">
                Deploy com controle.
                <br />
                Infraestrutura sem lock-in.
              </p>
            </div>

            {Object.entries(footerLinks).map(([section, links]) => (
              <div key={section}>
                <h4 className="text-sm font-semibold capitalize">{section}</h4>
                <ul className="mt-4 space-y-3">
                  {links.map((link) => (
                    <li key={link}>
                      <a
                        href="#"
                        className="text-sm text-muted-foreground transition-colors hover:text-foreground"
                      >
                        {link}
                      </a>
                    </li>
                  ))}
                </ul>
              </div>
            ))}
          </div>

          <Separator className="my-10 bg-arctis-border" />

          <div className="flex flex-col items-center justify-between gap-4 md:flex-row">
            <p className="text-xs text-arctis-text-dim">
              &copy; 2026 Arctis Deploy. Todos os direitos reservados.
            </p>
            <div className="flex gap-6">
              {["Privacidade", "Termos", "Segurança"].map((link) => (
                <a
                  key={link}
                  href="#"
                  className="text-xs text-arctis-text-dim transition-colors hover:text-foreground"
                >
                  {link}
                </a>
              ))}
            </div>
          </div>
        </div>
      </footer>
    </div>
  );
}
