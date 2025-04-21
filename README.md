# 🧙 Wallet Soul Agent

> AI-powered Solana wallet profiler that reflects the soul behind every address.

## ✨ What it does

- 🔍 Connects to Solana RPC to fetch wallet balances and tokens
- 🧠 Uses OpenAI to generate reflections and wallet archetypes
- 🌐 Exposes a clean JSON API
- 💻 Includes a frontend UI built with Next.js & Tailwind v4

## 📂 Structure

```
wallet-soul-agent/
│
├── agent/            # core logic (archetype, reflection, engine)
├── utils/            # Solana RPC & token fetchers
├── main.go           # Go entry point (API server)
│
├── soul-ui/          # Frontend app (Next.js 14, Tailwind v4)
│   └── pages/
│   └── tailwind.config.ts
│   └── ...
│
└── README.md
```

## 🚀 Run locally

### Backend (Go)

```bash
go run main.go
# API runs at: http://localhost:8080/api/wallet/:address
```

Requires `OPENAI_API_KEY` as environment variable.

### Frontend (Next.js)

```bash
cd soul-ui
npm install
npm run dev
# UI runs at: http://localhost:3000
```

## 📦 Example Response

```json
{
  "address": "So1111...xxxx",
  "profile": "Degen Wizard",
  "reflection": "This wallet echoes the chaos of meme cycles, yet hints at a hidden method...",
  "tokens": [
    { "symbol": "SOL", "amount": 3.1415 }
  ]
}
```

## 💡 Implementation Details

### Solana RPC Integration

The application uses JSON-RPC calls to fetch wallet data from Solana's blockchain:

```go
// Example of fetching SOL balance
func FetchTokens(address string) ([]Token, error) {
    // API endpoint
    url := "https://api.mainnet-beta.solana.com"
    
    // Create the RPC request
    requestBody := RPCRequest{
        Jsonrpc: "2.0",
        ID:      "1",
        Method:  "getBalance",
        Params:  []interface{}{address},
    }
    
    // Process response...
    // Convert lamports to SOL (1 SOL = 10^9 lamports)
    solBalance := float64(lamports) / 1000000000.0
    
    return tokens, nil
}
```

### Reflection Generator

The agent analyzes wallet contents and transaction patterns to generate personality reflections using OpenAI:

```go
func GenerateReflection(tokens []Token) (string, error) {
    // Format tokens for prompt
    tokensDescription := formatTokensForPrompt(tokens)
    
    // Generate via OpenAI
    prompt := fmt.Sprintf(
        "Analyze this Solana wallet contents and generate a witty, mystical reflection:\n%s",
        tokensDescription,
    )
    
    // Call OpenAI API and return reflection...
}
```

## 🛣️ Roadmap

- [ ] Add token transaction history analysis
- [ ] Implement NFT collection profiling
- [ ] Create Telegram bot interface
- [ ] Add wallet memory to track changes over time
- [ ] Deploy public demo instance

## 🌍 Live demo coming soon...

Stay tuned for the public deploy + Telegram bot + memory in v2.

---

## 🧠 Built by [@valeriibodnarchuk](https://github.com/valeriutechit)