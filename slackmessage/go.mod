module github.com/FreekingDean/jeffbotgo/slackmessage

require (
	cloud.google.com/go v0.35.1
	github.com/FreekingDean/jeffbotgo/messages v0.0.0
	github.com/FreekingDean/jeffbotgo/utils/pubsub v0.0.0
	github.com/FreekingDean/slackevents v0.0.1
)

replace (
	github.com/FreekingDean/jeffbotgo/messages => ../messages
	github.com/FreekingDean/jeffbotgo/utils/pubsub => ../utils/pubsub
)
