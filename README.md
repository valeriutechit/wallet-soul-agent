# 🧙 Wallet Soul Agent
Analyze the soul of any Solana wallet through GPT-powered archetypes and poetic reflection.

## 🌐 Live Demo  
[wallet-soul-agent.vercel.app](https://wallet-soul-agent.vercel.app)

## 🧠 Stack
- **Frontend**: Next.js + Tailwind CSS
- **Backend**: Go + SQLite + OpenAI API
- **Bot**: Telegram Bot API (go-telegram-bot-api)

## 📦 Features
- GPT-based soul archetype + poetic reflection
- Telegram bot for wallet analysis
- API endpoint: `/api/wallet/:address`
- SSR frontend for sharing
- Caching in SQLite for faster re-queries

## 🚀 Getting Started

### Prerequisites
- Node.js + npm
- Go (1.20+)

### Environment Variables
Create `.env` at the root:
```bash
OPENAI_API_KEY=your_openai_key
TELEGRAM_BOT_TOKEN=your_telegram_token
PORT=8080
```

Frontend uses `.env.local`:
```bash
NEXT_PUBLIC_SITE_URL=http://localhost:3000
```

### Install Dependencies
```bash
go mod tidy

cd soul-ui
npm install
```

### Run Locally
```bash
# frontend (http://localhost:3000)
cd soul-ui
npm run dev

# backend API (http://localhost:8080)
cd ..
go run main.go

# OR telegram bot (http://localhost:8080, responds in chat)
cd telegram
go run main.go
```

## 🧪 Testing
To test a wallet, try:
```
1BWutmTvYPwDtmw9abTkS4Ssr8no61spGAvW1X6NDix
```

## 📥 Deploy

### Frontend: Vercel
Deploy the `soul-ui` directory to Vercel.

### Backend API: Render or Railway
Deploy the root directory with the build command:
```
go build -o app ./main.go
```
Start command:
```
./app
```

### Telegram Bot: Railway
Create a new service with the build command:
```
go build -o app ./telegram
```
Start command:
```
./app
```

> **Note:** If you encounter deployment issues on Railway, try creating a new service instead of modifying an existing one to avoid cache problems.

## ✍️ Author
Valerii Bodnarchuk — [LinkedIn](https://www.linkedin.com/in/valerii-bodnarchuk) | [Telegram](https://t.me/valerii_bodnarchuk)

## ✨ Credit
Built in 4 days by @valerii_bodnarchuk & GPT (architect-assistant).

## 🪐 License
MIT