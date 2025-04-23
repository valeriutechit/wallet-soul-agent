import type { NextApiRequest, NextApiResponse } from 'next'

export default async function handler(req: NextApiRequest, res: NextApiResponse) {
  const { address } = req.query

  try {
    const resp = await fetch(`http://localhost:8080/api/wallet/${address}`)
    const data = await resp.json()
    res.status(200).json(data)
  } catch (err) {
    console.error('❌ Proxy API failed:', err)
    res.status(500).json({ error: 'Internal Server Error' })
  }
}
