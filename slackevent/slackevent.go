package slackevent

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/FreekingDean/jeffbot/utils/pubsub"
	"github.com/FreekingDean/slackevents"
)

func init() {
	slackevents.RegisterCallbackHandler(Event)
}

func SlackEvent(w http.ResponseWriter, r *http.Request) {
	slackevents.DefaultServer.ServeHTTP(w, r)
}

func Event(event *slackevents.Callback) error {
	data, err := json.Marshal(event.InnerEvent.ParsedEvent)
	if err != nil {
		return err
	}

	return pubsub.Publish(fmt.Sprintf(context.Background(), "%s.%s", "slack", event.InnerEvent.Type), data)
}
