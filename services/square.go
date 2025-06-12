package services

import (
	"context"
	"encoding/json"

	// "errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/joho/godotenv"
)

var (
	squareClient   *resty.Client
	locationID     string
	loyaltyAcctID  string
	squareBaseURL  = "https://connect.squareup.com/v2"
	squareAuth     string
	requestTimeout = 5 * time.Second
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	accessToken := os.Getenv("SQUARE_ACCESS_TOKEN")
	if accessToken == "" {
		log.Fatalf("SQUARE_ACCESS_TOKEN environment variable not set...")
	}
	squareAuth = "Bearer " + accessToken

	locationID = os.Getenv("SQUARE_LOCATION_ID")
	if locationID == "" {
		log.Fatalf("SQUARE_LOCATION_ID environment variable not set.")
	}

	loyaltyAcctID = os.Getenv("SQUARE_LOYALTY_ACCOUNT_ID")
	if loyaltyAcctID == "" || loyaltyAcctID == "YOUR_SQUARE_LOYALTY_ACCOUNT_ID_HERE" {
		log.Println("WARNING: loyaltyAcctID is missing or still a placeholder. Please provide a valid Square Loyalty Account ID.")
	}

	squareClient = resty.New().
		SetHostURL(squareBaseURL).
		SetHeader("Authorization", squareAuth).
		SetHeader("Content-Type", "application/json")
}

func EarnPoints(points int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()

	idempotencyKey := time.Now().Format("20060102150405.000000")

	payload := map[string]interface{}{
		"accumulate_points": map[string]interface{}{
			"points": points,
		},
		"idempotency_key": idempotencyKey,
	}

	resp, err := squareClient.R().
		SetContext(ctx).
		SetBody(payload).
		Post(fmt.Sprintf("/loyalty/accounts/%s/accumulate", loyaltyAcctID))

	if err != nil {
		return fmt.Errorf("failed to call Square API: %w", err)
	}

	if resp.IsError() {
		return fmt.Errorf("Square API error: %s", resp.String())
	}

	log.Printf("Successfully earned %d points", points)
	return nil
}

func RedeemPoints(points int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()

	idempotencyKey := time.Now().Format("20060102150405.000000")

	payload := map[string]interface{}{
		"redeem_points": map[string]interface{}{
			"points": points,
		},
		"idempotency_key": idempotencyKey,
	}

	resp, err := squareClient.R().
		SetContext(ctx).
		SetBody(payload).
		Post(fmt.Sprintf("/loyalty/accounts/%s/redeem", loyaltyAcctID))

	if err != nil {
		return fmt.Errorf("failed to call Square API: %w", err)
	}

	if resp.IsError() {
		return fmt.Errorf("Square API error: %s", resp.String())
	}

	log.Printf("Successfully redeemed %d points", points)
	return nil
}

func GetBalance() (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()

	resp, err := squareClient.R().
		SetContext(ctx).
		Get(fmt.Sprintf("/loyalty/accounts/%s", loyaltyAcctID))

	if err != nil {
		log.Printf("Error calling Square API: %v", err)
		return 0, fmt.Errorf("failed to call Square API: %w", err)
	}

	if resp.IsError() {
		log.Printf("Square API responded with error: %s", resp.String())
		return 0, fmt.Errorf("Square API error: %s", resp.String())
	}

	log.Printf("Square API raw response: %s", resp.String())

	var result struct {
		LoyaltyAccount struct {
			AccumulatedPoints int64 `json:"accumulated_points"`
		} `json:"loyalty_account"`
	}
	err = json.Unmarshal(resp.Body(), &result)
	if err != nil {
		log.Printf("Failed to parse Square API response: %v", err)
		return 0, fmt.Errorf("failed to parse response: %w", err)
	}

	return result.LoyaltyAccount.AccumulatedPoints, nil
}

func GetHistory() ([]map[string]interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()

	resp, err := squareClient.R().
		SetContext(ctx).
		Get(fmt.Sprintf("/loyalty/accounts/%s/events", loyaltyAcctID))

	if err != nil {
		return nil, fmt.Errorf("failed to call Square API: %w", err)
	}

	if resp.IsError() {
		return nil, fmt.Errorf("Square API error: %s", resp.String())
	}

	var result struct {
		Events []map[string]interface{} `json:"events"`
	}
	err = json.Unmarshal(resp.Body(), &result)
	if err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return result.Events, nil
}
