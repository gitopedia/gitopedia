package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Fprintf(os.Stderr, "Usage: %s <APP_ID> <PRIVATE_KEY_PATH> <INSTALLATION_ID>\n", os.Args[0])
		os.Exit(1)
	}

	appID := os.Args[1]
	keyPath := os.Args[2]
	installID := os.Args[3]

	keyBytes, err := os.ReadFile(keyPath)
	if err != nil {
		log.Fatalf("Failed to read key: %v", err)
	}

	token, err := generateJWT(appID, keyBytes)
	if err != nil {
		log.Fatalf("Failed to generate JWT: %v", err)
	}

	installToken, err := getInstallationToken(token, installID)
	if err != nil {
		log.Fatalf("Failed to get installation token: %v", err)
	}

	fmt.Println(installToken)
}

func generateJWT(appID string, keyBytes []byte) (string, error) {
	key, err := jwt.ParseRSAPrivateKeyFromPEM(keyBytes)
	if err != nil {
		return "", err
	}

	now := time.Now()
	claims := jwt.MapClaims{
		"iat": now.Add(-60 * time.Second).Unix(),
		"exp": now.Add(10 * time.Minute).Unix(),
		"iss": appID,
	}

	t := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	return t.SignedString(key)
}

func getInstallationToken(jwtToken, installID string) (string, error) {
	url := fmt.Sprintf("https://api.github.com/app/installations/%s/access_tokens", installID)
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", "Bearer "+jwtToken)
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("API error %d: %s", resp.StatusCode, string(body))
	}

	var res struct {
		Token string `json:"token"`
	}
	if err := json.Unmarshal(body, &res); err != nil {
		return "", err
	}
	return res.Token, nil
}











