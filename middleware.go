package line

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// BotMiddleware ...
func BotMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Our middleware logic goes here...

		log.Print("===exampleMiddleware===")

		var receivedMessage ReceivedMessage
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Print(err)
		}
		err = json.Unmarshal(b, &receivedMessage)
		if err != nil {
			log.Print(err)
		}

		for _, result := range receivedMessage.Result {
			from := result.Content.From
			log.Print(from, result.Content.Text)
		}
		next.ServeHTTP(w, r)
	})
}
