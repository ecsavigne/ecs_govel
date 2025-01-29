package migration

type Test struct {
	ID  uint   `gorm:"primary_key"`
	Url string `gorm:"type:varchar(255);not null"`
}
