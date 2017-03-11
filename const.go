package vtm

const (
	defaultEventsURL = "/event"

	/* --- api related constants --- */
	marathonAPIVersion      = "v2"
	marathonAPIEventStream  = marathonAPIVersion + "/events"
	marathonAPISubscription = marathonAPIVersion + "/eventSubscriptions"
	marathonAPIPools        = "api/tm/3.5/config/active/pools"
	marathonAPITasks        = marathonAPIVersion + "/tasks"
	marathonAPIDeployments  = marathonAPIVersion + "/deployments"
	marathonAPIGroups       = marathonAPIVersion + "/groups"
	marathonAPIQueue        = marathonAPIVersion + "/queue"
	marathonAPIInfo         = marathonAPIVersion + "/info"
	marathonAPILeader       = marathonAPIVersion + "/leader"
	marathonAPIPing         = "ping"
)

const (
	// EventsTransportCallback activates callback events transport
	EventsTransportCallback EventsTransport = 1 << iota

	// EventsTransportSSE activates stream events transport
	EventsTransportSSE
)
