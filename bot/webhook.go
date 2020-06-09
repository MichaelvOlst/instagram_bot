package bot

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

// WebhookResponse ...
type WebhookResponse struct {
	Response *InstagramResult `json:"response,omitempty"`
	Error    error            `json:"error,omitempty"`
}

func (b *Bot) makeRequest(w *WebhookResponse) {

	s, err := json.Marshal(w)
	if err != nil {
		log.Fatal(err)
		return
	}

	jsonBytes := []byte(s)
	req, err := http.NewRequest("POST", b.Webhook, bytes.NewBuffer(jsonBytes))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
}
