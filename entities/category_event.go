package entities

type CategoryEvent struct {
	ID         uint       `gorm:"primaryKey;autoIncrement" json:"id"`
	Nama       string     `gorm:"type:varchar(50)" json:"nama"`
}