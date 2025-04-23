import { Html, Head, Main, NextScript } from 'next/document'

export default function Document() {
  return (
    <Html>
      <Head>
        <title>Wallet Soul Agent</title>
        <meta name="description" content="Discover the soul of any Solana wallet. AI-powered archetype & reflection. Telegram-ready." />

        {/* Open Graph */}
        <meta property="og:title" content="Wallet Soul Agent" />
        <meta property="og:description" content="AI-powered soul analysis of Solana wallets. Archetypes. Reflections. Telegram-ready." />
        <meta property="og:image" content="https://walletsoul.xyz/og-image.png" />
        <meta property="og:url" content="https://walletsoul.xyz" />
        <meta property="og:type" content="website" />

        {/* Twitter Card */}
        <meta name="twitter:card" content="summary_large_image" />
        <meta name="twitter:title" content="Wallet Soul Agent" />
        <meta name="twitter:description" content="Discover the soul of any wallet on Solana. GPT-powered archetype + poetic reflection." />
        <meta name="twitter:image" content="https://walletsoul.xyz/og-image.png" />
      </Head>
      
      <body>
        <Main />
        <NextScript />
      </body>
    </Html>
  )
}
