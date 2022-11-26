package statistical

type Stat struct {
	ID         uint `gorm:"primaryKey"`
	Url        string
	StatusCode int
	Duration   int64 // Its unit is millisecond
	ReceivedAt int64 `gorm:"index"` // Its unit is millisecond
}
