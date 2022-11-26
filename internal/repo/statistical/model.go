package statistical

type Statistical struct {
	ID         uint `gorm:"primaryKey"`
	Url        string
	StatusCode int   // Http status code
	Duration   int64 // Its unit is millisecond
	ReceivedAt int64 `gorm:"index"` // Its unit is millisecond
}
