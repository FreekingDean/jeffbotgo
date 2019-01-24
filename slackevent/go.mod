module github.com/FreekingDean/slackevent

require (
	github.com/FreekingDean/jeffbot/utils/pubsub v0.0.0
	github.com/FreekingDean/slackevents v0.0.0
)

replace (
	github.com/FreekingDean/jeffbot/utils/pubsub => ../utils/pubsub
	github.com/FreekingDean/slackevents => ../../slackevents
)
