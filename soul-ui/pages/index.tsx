import { useState } from 'react'
import Head from 'next/head'

export default function Home() {
  const [address, setAddress] = useState('')
  const [report, setReport] = useState<any>(null)
  const [error, setError] = useState<string | null>(null)

  const fetchSoul = async () => {
    setError(null)
    try {
      const res = await fetch(`http://localhost:8080/api/wallet/${address}`)
      if (!res.ok) throw new Error('Server error or wallet not found')
      const data = await res.json()
      setReport(data)
    } catch (err: any) {
      setError(err.message || 'Unexpected error')
      setReport(null)
    }
  }

  return (
    <>
      <Head>
        <title>Wallet Soul Agent</title>
        <link href="https://fonts.googleapis.com/css2?family=JetBrains+Mono:wght@400;600&display=swap" rel="stylesheet" />
      </Head>

      <div className="min-h-screen bg-gradient-to-br from-zinc-900 via-black to-zinc-800 text-white font-mono flex flex-col items-center justify-center px-4 py-10">
        <h1 className="text-4xl md:text-5xl font-bold mb-8 text-center tracking-tight">
          🧙 Wallet Soul Agent
        </h1>

        <div className="flex flex-col w-full max-w-[520px] overflow-hidden rounded-xl border border-zinc-700">
          <input
            className="bg-zinc-900 px-4 py-2 text-sm text-white placeholder-zinc-400 focus:outline-none focus:ring-2 focus:ring-indigo-500"
            placeholder="Enter Solana wallet address"
            value={address}
            onChange={(e) => setAddress(e.target.value)}
          />
          <button
            onClick={fetchSoul}
            className="bg-indigo-600 hover:bg-indigo-700 transition-colors px-4 py-2 text-sm text-white font-semibold"
          >
            🔍 Analyze Soul
          </button>
        </div>
        {error && (
          <div className="my-4 text-sm text-red-400 bg-red-900 bg-opacity-20 p-2 rounded">
            ⚠️ {error}
          </div>
        )}

        {report && (
          <div className="mt-10 w-full max-w-[640px] bg-zinc-900 border border-zinc-700 p-6 rounded-xl shadow-xl">
            <p className="mb-3 text-zinc-400 text-sm">📍 <span className="text-white">Address:</span> {report.address}</p>
            <p className="mb-2 text-sm">🧠 <span className="font-semibold">Archetype:</span> <span className="text-indigo-400">{report.profile}</span></p>
            <p className="mb-2 text-sm">🪞 <span className="font-semibold">Reflection:</span></p>
            <p className="italic text-zinc-300 text-sm border-l-4 border-indigo-500 pl-4">{report.reflection}</p>
            <p className="mt-4 font-semibold text-sm">💎 Tokens:</p>
            <ul className="list-disc list-inside text-zinc-300 text-sm">
              {report.tokens.map((t: any, idx: number) => (
                <li key={idx}>
                  {t.symbol}: {t.amount.toFixed(4)}
                </li>
              ))}
            </ul>
          </div>
        )}
      </div>
    </>
  )
}
