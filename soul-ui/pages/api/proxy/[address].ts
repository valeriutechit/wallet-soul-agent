// pages/api/proxy/[address].ts
export default async function handler(req, res) {
  const { address } = req.query
  try {
    const resp = await fetch(`http://localhost:8080/api/wallet/${address}`)
    const data = await resp.json()
    res.status(200).json(data)
  } catch (e) {
    res.status(500).json({ error: 'Proxy error' })
  }
}
