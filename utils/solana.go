// utils/solana.go
package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Полная структура ответа RPC
type SolanaRpcResponse struct {
	Jsonrpc string `json:"jsonrpc"`
	Result  struct {
		Context struct {
			ApiVersion string `json:"apiVersion"`
			Slot       uint64 `json:"slot"`
		} `json:"context"`
		Value uint64 `json:"value"`
	} `json:"result"`
	ID     string    `json:"id"`
	Error  *RPCError `json:"error,omitempty"`
}

type RPCError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Token struct {
	Mint     string  `json:"mint"`
	Name     string  `json:"name"`
	Symbol   string  `json:"symbol"`
	UiAmount float64 `json:"amount"`
}

func FetchTokens(address string) ([]Token, error) {
	// Используем публичный RPC-endpoint Solana
	url := "https://api.mainnet-beta.solana.com"
	
	fmt.Println("🔍 Requesting from:", url)
	
	// Создаем JSON-RPC запрос для getBalance
	requestBody := map[string]interface{}{
		"jsonrpc": "2.0",
		"id":      "1",
		"method":  "getBalance",
		"params":  []interface{}{address},
	}
	
	// Конвертируем запрос в JSON
	requestJSON, err := json.Marshal(requestBody)
	if err != nil {
		fmt.Println("❌ Error marshaling request:", err)
		return nil, err
	}
	
	// Отправляем POST запрос
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(requestJSON))
	if err != nil {
		fmt.Println("❌ Error during request:", err)
		return nil, err
	}
	defer resp.Body.Close()
	
	fmt.Printf("🔄 RPC response status code: %d\n", resp.StatusCode)
	
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("❌ Error reading response body:", err)
		return nil, err
	}
	
	fmt.Println("📦 Raw response:")
	fmt.Println(string(bodyBytes))
	
	// Разбираем ответ используя правильную структуру
	var rpcResp SolanaRpcResponse
	if err := json.Unmarshal(bodyBytes, &rpcResp); err != nil {
		fmt.Println("❌ Error unmarshaling JSON:", err)
		return nil, err
	}
	
	// Проверяем ошибки RPC
	if rpcResp.Error != nil {
		return nil, fmt.Errorf("RPC error %d: %s", rpcResp.Error.Code, rpcResp.Error.Message)
	}
	
	// Получаем значение в lamports
	lamports := rpcResp.Result.Value
	
	// Конвертируем lamports в SOL (1 SOL = 10^9 lamports)
	solBalance := float64(lamports) / 1000000000.0
	
	fmt.Printf("✅ Successfully retrieved SOL balance: %f SOL\n", solBalance)
	
	// Возвращаем один токен, представляющий SOL
	tokens := []Token{
		{
			Mint:     "So11111111111111111111111111111111111111112", // Native SOL mint address
			Name:     "Solana",
			Symbol:   "SOL",
			UiAmount: solBalance,
		},
	}
	
	return tokens, nil
}