package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"

	// "io/ioutil"
	// "log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)



func InitializePayment(email string, amount int64) (*PaystackInitializeResponse, error) {
	err := godotenv.Load()
    if err != nil {
        fmt.Println("⚠️  No .env file found, using system env")
		return nil, err
    }
	payload := map[string]interface{}{
		"email": email,
		"amount": amount,
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", "https://api.paystack.co/transaction/initialize", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+os.Getenv("PAYSTACK_SECRET"));
	req.Header.Set("Content-Type", "application/json");

	client := &http.Client{Timeout: 15 * time.Second}

	response, err :=client.Do(req)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var paystackResponse PaystackInitializeResponse
	if err := json.Unmarshal(body, &paystackResponse); err != nil {
		return nil, err
	}
	return &paystackResponse, nil
}