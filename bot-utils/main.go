package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/EdgeJay/psg-navi-bot/bot-backend/utils"
)

type InitResponse struct {
	Status  string `json:"status"`
	Version string `json:"version"`
	Error   string `json:"error,omitempty"`
}

func doInitBot(url, hashed string) {
	req, err := http.NewRequest("POST", url, strings.NewReader(""))
	if err != nil {
		log.Fatalln("unable to create request")
	}

	req.Header.Add("X-PSGNaviBot-Init-Token", hashed)

	client := http.Client{
		Timeout: time.Second * 30,
	}

	if res, err := client.Do(req); err != nil {
		log.Fatalln("request failed:", err)
	} else {
		defer res.Body.Close()

		var initResponse InitResponse
		if err := json.NewDecoder(res.Body).Decode(&initResponse); err != nil {
			log.Fatalln("invalid response:", err)
		} else {
			if initResponse.Error != "" {
				log.Println("request failed with error:", initResponse.Error)
			}
			log.Println(initResponse)
		}
	}
}

func main() {
	url := flag.String("url", "", "url to post request to")
	appVersion := flag.String("version", "0.0.0", "app version")
	tokenSecret := flag.String("secret", "invalid_init_token_secret", "init token secret")
	flag.Parse()

	if hashed, err := utils.CreateHmacHexString(*appVersion, []byte(*tokenSecret)); err != nil {
		log.Fatalln("unable to generate hmac hex string", err)
	} else {
		doInitBot(*url, hashed)
	}
}
