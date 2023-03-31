package entities

type StatusPembayaran struct {
	ID         uint       `gorm:"primaryKey;autoIncrement" json:"id"`
	Status     string     `gorm:"type:varchar(50)" json:"status"`
	
	Pembayaran Pembayaran `gorm:"foreignKey:StatusPembayaranID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
}
