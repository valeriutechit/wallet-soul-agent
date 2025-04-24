import bs58 from 'bs58'

export function isValidSolanaAddress(address: string): boolean {
  try {
    const decoded = bs58.decode(address.trim())
    return decoded.length === 32
  } catch {
    return false
  }
}
