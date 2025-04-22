import { GetServerSideProps } from 'next'

type Token = {
  mint: string
  name: string
  symbol: string
  amount: number
}

type Report = {
  Address: string
  Profile: string
  Reflection: string
  Tokens: Token[]
}

type Props = {
  report: Report | null
}

export default function WalletPage({ report }: Props) {
  if (!report) {
    return <div className="text-red-500 p-10">❌ Wallet not found or server error</div>
  }

  return (
    <div className="min-h-screen flex items-center justify-center bg-black text-white p-6">
      <div className="w-full max-w-md bg-zinc-900 border border-zinc-700 p-6 rounded-xl">
        <h1 className="text-2xl font-bold mb-4">🧙 Wallet Soul Agent</h1>
        <p className="mb-2 text-zinc-400 text-sm">📍 <span className="text-white">Address:</span> {report.Address}</p>
        <p className="mb-2">🧠 <span className="font-semibold">Archetype:</span> <span className="text-indigo-400">{report.Profile}</span></p>
        <p className="mb-2">🪞 <span className="font-semibold">Reflection:</span></p>
        <p className="italic text-zinc-300 border-l-4 border-indigo-500 pl-4">{report.Reflection}</p>
        <p className="mt-4 font-semibold">💎 Tokens:</p>
        <ul className="list-disc list-inside text-zinc-300 text-sm">
          {report.Tokens.map((t, idx) => (
            <li key={idx}>{t.symbol}: {t.amount.toFixed(4)}</li>
          ))}
        </ul>
      </div>
    </div>
  )
}

export const getServerSideProps: GetServerSideProps = async (context) => {
  const address = context.params?.address as string

  try {
    const res = await fetch(`${process.env.NEXT_PUBLIC_API_URL}/api/wallet/${address}`)
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
