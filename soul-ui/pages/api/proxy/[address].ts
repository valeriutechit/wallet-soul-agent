import type { NextApiRequest, NextApiResponse } from 'next'

export default async function handler(req: NextApiRequest, res: NextApiResponse) {
  const { address } = req.query

  try {
    const apiBase = process.env.NEXT_PUBLIC_API_URL || "http://localhost:8080"
    const resp = await fetch(`${apiBase}/api/wallet/${address}`)
    const data = await resp.json()
    res.status(200).json(data)
  } catch (err) {
    console.error('❌ Proxy API failed:', err)
    res.status(500).json({ error: 'Internal Server Error' })
  }
}
