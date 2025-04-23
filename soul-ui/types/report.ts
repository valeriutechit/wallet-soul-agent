export type Token = {
  mint: string
  name: string
  symbol: string
  amount: number
}

export type WalletReport = {
  address: string
  profile: string
  reflection: string
  tokens: Token[]
}