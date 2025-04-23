import Link from 'next/link'
import { GetServerSideProps } from 'next'
import { type WalletReport } from '@/types/report'
import { API_BASE_URL, SITE_BASE_URL } from '@/lib/config'

type Props = {
  report: WalletReport | null
}

export default function WalletPage({ report }: Props) {
  if (!report) {
    return <div className="text-red-500 p-10">❌ Wallet not found or server error</div>
  }

  const currentUrl = `${SITE_BASE_URL}/wallet/${report.address}`
  const customText = `🧙 Check out the soul of this wallet: ${report.profile} — ${report.reflection}`

  return (
    <div className="min-h-screen flex flex-col items-center justify-center bg-black text-white p-6">
      <Link href="/" className="mb-4 text-center w-full text-sm text-indigo-400 hover:underline self-start">
        ← Back to Home
      </Link>

      <div className="w-full max-w-md bg-zinc-900 border border-zinc-700 p-6 rounded-xl">
        <h1 className="text-2xl font-bold mb-4">🧙 Wallet Soul Agent</h1>
        <p className="mb-2 text-zinc-400 text-sm">📍 <span className="text-white">Address:</span> <span className="block truncate">{report.address}</span></p>
        <p className="mb-2">🧠 <span className="font-semibold">Archetype:</span> <span className="text-indigo-400">{report.profile}</span></p>
        <p className="mb-2">🪞 <span className="font-semibold">Reflection:</span></p>
        <p className="italic text-zinc-300 border-l-4 border-indigo-500 pl-4">{report.reflection}</p>
        <p className="mt-4 font-semibold">💎 Tokens:</p>
        {report.tokens && report.tokens.length > 0 ? (
          <ul className="list-disc list-inside text-zinc-300 text-sm">
            {report.tokens.map((token, idx) => (
              <li key={idx}>{token.symbol}: {token.amount.toFixed(4)}</li>
            ))}
          </ul>
        ) : (
          <p className="text-zinc-500 italic mt-2">No tokens found in this wallet.</p>
        )}
      </div>

      {report && (
        <a
          href={`https://t.me/share/url?url=${encodeURIComponent(currentUrl)}&text=${encodeURIComponent(customText)}`}
          target="_blank"
          rel="noopener noreferrer"
          className="bg-blue-800 hover:bg-blue-900 text-white font-semibold py-2 px-4 rounded shadow transition mt-6"
        >
          📤 Share in Telegram
        </a>
      )}
    </div>
  )
}

export const getServerSideProps: GetServerSideProps = async (context) => {
  const address = context.params?.address as string

  try {
    const res = await fetch(`${API_BASE_URL}/api/wallet/${address}`)
    const json = await res.json()

    return {
      props: {
        report: json
      }
    }
  } catch (error) {
    console.error("❌ Fetch failed:", error)
    return {
      props: {
        report: null
      }
    }
  }
}
