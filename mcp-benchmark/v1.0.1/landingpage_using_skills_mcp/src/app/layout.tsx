import type { Metadata } from "next";
import { Inter, Geist_Mono } from "next/font/google";
import "./globals.css";

const inter = Inter({
  variable: "--font-sans",
  subsets: ["latin"],
  display: "swap",
  weight: ["400", "500", "600", "700"],
});

const geistMono = Geist_Mono({
  variable: "--font-mono",
  subsets: ["latin"],
  display: "swap",
});

export const metadata: Metadata = {
  title: "Arctis Deploy — Deploy infrastructure on your terms",
  description:
    "More control, less lock-in, sustainable costs. Arctis Deploy gives developers a deployment platform that scales with you, not against you.",
  keywords: [
    "deploy",
    "infrastructure",
    "cloud",
    "devops",
    "CI/CD",
    "serverless",
  ],
  openGraph: {
    title: "Arctis Deploy",
    description:
      "Deploy infrastructure on your terms. More control, less lock-in.",
    type: "website",
  },
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html
      lang="en"
      className={`${inter.variable} ${geistMono.variable} h-full antialiased`}
    >
      <body className="min-h-full flex flex-col">{children}</body>
    </html>
  );
}
