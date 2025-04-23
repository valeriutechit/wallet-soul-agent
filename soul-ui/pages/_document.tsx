import { Html, Head, Main, NextScript } from 'next/document'

export default function Document() {
  return (
    <Html>
      <Head>
        <link rel="icon" type="image/png" href="/favicon/favicon-96x96.png" />
        <link rel="apple-touch-icon" href="/favicon/apple-touch-icon.png" />

        <meta name="description" content="Discover the soul of any Solana wallet. AI-powered archetype & reflection. Telegram-ready." />

        <meta property="og:title" content="Wallet Soul Agent" />
        <meta property="og:description" content="AI-powered soul analysis of Solana wallets. Archetypes. Reflections. Telegram-ready." />
        <meta property="og:image" content="https://wallet-soul-agent.vercel.app/og-image.png" />
        <meta property="og:url" content="https://wallet-soul-agent.vercel.app" />
        <meta property="og:type" content="website" />

        <meta name="twitter:card" content="summary_large_image" />
        <meta name="twitter:title" content="Wallet Soul Agent" />
        <meta name="twitter:description" content="Discover the soul of any wallet on Solana. GPT-powered archetype + poetic reflection." />
        <meta name="twitter:image" content="https://wallet-soul-agent.vercel.app/og-image.png" />
      </Head>
      <body>
        <Main />
        <NextScript />
      </body>
    </Html>
  )
}
