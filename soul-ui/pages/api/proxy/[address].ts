import { API_BASE_URL } from '@/lib/config'
import type { NextApiRequest, NextApiResponse } from 'next'

export default async function handler(req: NextApiRequest, res: NextApiResponse) {
  const { address } = req.query

  try {
    const resp = await fetch(`${API_BASE_URL}/api/wallet/${address}`)
    const data = await resp.json()
    res.status(200).json(data)
  } catch (err) {
    console.error('❌ Proxy API failed:', err)
    res.status(500).json({ error: 'Internal Server Error' })
  }
}
