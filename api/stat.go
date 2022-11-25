package api

type StatsRequest struct {
	Time int // Number of minutes in the past from current time
}

type StatsResponse struct {
	ForwardingStats ForwardingStats `json:"forwarding_stats"`
	MaxDurations    MaxDurations    `json:"max_durations"`
}

type ForwardingStats struct {
	SuccessCount int `json:"success_count"`
	FailCount    int `json:"fail_count"`
}

type MaxDurations struct {
	UrlDuration map[string]int
}
