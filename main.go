package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Config struct {
	SonarQubeURL string `json:"sonarqube_url"`
	APIToken     string `json:"api_token"`
}

type AuthenticationResponse struct {
	Valid bool `json:"valid"`
}

func main() {
	config, err := loadConfig("config.json")
	if err != nil {
		fmt.Println("Error loading config:", err)
		return
	}

	client := http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	url := fmt.Sprintf("%s/api/authentication/validate", config.SonarQubeURL)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", config.APIToken))

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Authentication request failed with status:", resp.Status)
		return
	}

	var authResponse AuthenticationResponse
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	err = json.Unmarshal(body, &authResponse)
	if err != nil {
		fmt.Println("Error decoding response JSON:", err)
		return
	}

	if authResponse.Valid {
		fmt.Println("Authentication successful!")
	} else {
		fmt.Println("Authentication failed: Invalid credentials")
	}
}

func loadConfig(filename string) (Config, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return Config{}, err
	}

	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}
