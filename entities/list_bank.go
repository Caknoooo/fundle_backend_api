package entities

type ListBank struct {
	ID         uint       `gorm:"primaryKey;autoIncrement" json:"id"`
	Nama       string     `gorm:"type:varchar(50)" json:"nama"`
	
	Pembayaran Pembayaran `gorm:"foreignKey:ListBankID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"pembayaran,omitempty"`
}
