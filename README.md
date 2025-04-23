# 🧙 Wallet Soul Agent

> AI-powered Web3 soul analyzer — get an archetype + poetic reflection of any Solana wallet. Telegram-ready. Open-source. Instant.

---

## ⚡ What is it?
**Wallet Soul Agent** is a minimalistic AI + Web3 project that:

- Takes a **Solana wallet address**
- Fetches its token balance
- Sends it to **GPT** for analysis
- Returns:
  - 🧠 **Archetype** (Monk, Alchemist, etc.)
  - 🪞 **Reflection** (poetic GPT response)

---

## 📦 Tech Stack
- **Go (Golang)** — backend + Telegram bot
- **OpenAI API** — for archetype + reflection
- **Solana RPC** — to fetch wallet token balances
- **Next.js (React)** — SSR frontend
- **Tailwind CSS** — UI styling
- **SQLite** — lightweight caching

---

## 🧪 Try it out
- 👉 [Telegram bot](https://t.me/wallet_soul_agent_bot)
- 👉 [Web UI](https://soon.com) ← _coming after deploy_

---

## 🔧 Run locally

```bash
git clone https://github.com/your-name/wallet-soul-agent.git
cd wallet-soul-agent

# 1. Backend
cp .env.example .env
# Add your OpenAI + Telegram tokens

go run main.go

# 2. Frontend (in soul-ui/)
cd soul-ui
npm install
npm run dev
```

---

## 📸 Screenshot
![wallet-soul-agent-preview](public/screenshot.png)

---

## 💭 Why?
Because:
> _"The soul leaves a trace — even on-chain."_

---

## ✨ Credit
Built in 3 days by [Valerii Bodnarchuk](https://t.me/valeriubodnarchuk) & GPT.

---

## 🪐 License
MIT. Use it. Fork it. Remix it. Let's make wallets soulful.

