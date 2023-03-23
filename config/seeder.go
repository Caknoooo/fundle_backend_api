package config

import (
	"github.com/Caknoooo/golang-clean_template/entities"
	"gorm.io/gorm"
)

func Seeder(db *gorm.DB) error{
	if err := ListBankSeeder(db); err != nil {
		return err
	}
	
	if err := StatusPembayaranSeeder(db); err != nil {
		return err
	}

	return nil
}

func ListBankSeeder(db *gorm.DB) error{
	var listBanks = []entities.ListBank{
		{
			ID: 1, 
			Nama: "BNI",
		},
		{
			ID: 2, 
			Nama: "BCA", 
		},
		{
			ID: 3, 
			Nama: "Mandiri",
		},
		{
			ID: 4,
			Nama: "BSI",
		},
		{
			ID: 5, 
			Nama: "OVO",
		}, 
		{
			ID: 6,
			Nama: "Gopay",
		},
	}

	for _, data := range listBanks {
		if err := db.Create(&data).Error; err != nil {
			return err
		}
	}

	return nil
}

func StatusPembayaranSeeder(db *gorm.DB) error {
	var statusPembayaran = []entities.StatusPembayaran{
		{
			ID: 1, 
			Status: "Failed", 
		},
		{
			ID: 2, 
			Status: "Success",
		}, 
		{
			ID: 3,
			Status: "Awaiting",
		},
	}

	for _, data := range statusPembayaran{
		if err := db.Create(&data).Error; err != nil {
			return err
		}
	}

	return nil
}