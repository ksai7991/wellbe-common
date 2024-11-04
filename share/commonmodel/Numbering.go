package commonmodel

type Numbering struct {
	NumberingKey string
	InitialValue int64
	CurrentValue int64
	MaxValue int64
	FixLength int // 0 means no fixed
	CreateDatetime string
	CreateFunc string
	UpdateDatetime string
	UpdateFunc string
}