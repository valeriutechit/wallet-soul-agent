import { useState } from 'react'
import Head from 'next/head'
import { useRouter } from 'next/router'
import { API_BASE_URL } from '@/lib/config'
import { isValidSolanaAddress } from '@/utils/solana'

export default function Home() {
  const [address, setAddress] = useState('')
  const [loading, setLoading] = useState(false)
  const [error, setError] = useState<string | null>(null)
  const router = useRouter()

  const fetchSoul = async () => {
    setError(null)

    if (!isValidSolanaAddress(address)) {
      setError('Invalid Solana address')
      return
    }

    setLoading(true)
    try {
      const res = await fetch(`${API_BASE_URL}/api/wallet/${address}`)
      if (!res.ok) throw new Error('Server error or wallet not found')
      await res.json()
      router.push(`/wallet/${address}`)
    } catch (err) {
      if (err instanceof Error) {
        setError(err.message)
      } else {
        setError('Unexpected error')
      }
    } finally {
      setLoading(false)
    }
  }

  return (
    <>
      <Head>
        <title>Wallet Soul Agent</title>
        <link href="https://fonts.googleapis.com/css2?family=JetBrains+Mono:wght@400;600&display=swap" rel="stylesheet" />
      </Head>

      <div className="min-h-screen flex items-center justify-center bg-gradient-to-br from-zinc-900 via-black to-zinc-800 text-white font-mono">
        <div className="relative flex flex-col items-center justify-center px-4 py-10">
          {error && (
            <div className="absolute -top-10 transition-all duration-300 max-w-[300px] text-sm text-red-400 bg-red-900/70 border border-red-600 px-4 py-2 rounded-lg shadow-md animate-fade-in">
              ⚠️ {error}
            </div>
          )}

          <h1 className="text-4xl md:text-5xl font-bold mb-8 text-center tracking-tight">
            🧙 Wallet Soul Agent
          </h1>

          <div className="flex flex-col w-full max-w-[520px] overflow-hidden rounded-xl border border-zinc-700">
            <input
              className="bg-zinc-900 px-4 py-2 text-sm text-white placeholder-zinc-400 focus:outline-none focus:ring-2 focus:ring-indigo-500"
              placeholder="Enter Solana wallet address"
              value={address}
              onChange={(e: React.ChangeEvent<HTMLInputElement>) => setAddress(e.target.value)}
            />

            {loading ? (
              <div className="flex justify-center items-center min-h-[36px] text-indigo-400 transition-all duration-300 animate-fade-in text-sm">🔄 Analyzing...</div>
            ) : (
              <button
                onClick={loading ? undefined : fetchSoul}
                disabled={loading || address.trim().length === 0}
                className="bg-indigo-600 disabled:opacity-50 disabled:bg-gray-600 min-h-[36px] hover:bg-indigo-700 focus:ring-4 focus:ring-indigo-400/40 transition-all animate-fade-in duration-300"
              >
                🔍 Analyze Soul
              </button>
            )}
          </div>

          <a
            href="https://t.me/wallet_soul_agent_bot"
            target="_blank"
            rel="noopener noreferrer"
            className="bg-blue-800 hover:bg-blue-900 text-white font-semibold py-2 px-4 rounded shadow transition mt-6 md:mt-10"
          >
            🤖 Run in Telegram
          </a>
        </div>
      </div>
    </>
  )
}
