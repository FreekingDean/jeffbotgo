package slackevent

import (
	"context"
	"fmt"
	"net/http"

	"github.com/FreekingDean/jeffbotgo/utils/pubsub"
	"github.com/FreekingDean/slackevents"
)

func init() {
	slackevents.RegisterCallbackHandler(Event)
}

func SlackEvent(w http.ResponseWriter, r *http.Request) {
	slackevents.DefaultServer.ServeHTTP(w, r)
}

func Event(event *slackevents.Callback) error {
	return pubsub.Publish(context.Background(), fmt.Sprintf("%s.%s", "slack", event.InnerEvent.Type), event.InnerEvent.ParsedEvent)
}
