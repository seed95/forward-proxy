package model

type Statistical struct {
	Url                  string
	StatusCode           int
	DurationResponseTime int64 // Its unit is millisecond
	ReceivedAt           int64 // Its unit is millisecond
}
