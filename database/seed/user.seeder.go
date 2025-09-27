package seed

import (
	"errors"
	"log"

	"github.com/bagusyanuar/app-hr-be/internal/domain/entity"
	"gorm.io/gorm"
)

func UserSeeder(db *gorm.DB) {
	if !db.Migrator().HasTable("users") {
		log.Println("⛔ Table users not found, seeding cancelled.")
		return
	}

	email := "superdev@web.id"
	username := "superdev"
	password := "$2a$13$23Ysb/ADn.t/OmRNVXyf0eOFhpyM5NMelOxClqZ2lz8Uf/HRQ/4C2"

	data := entity.User{
		Email:    email,
		Username: username,
		Password: password,
	}

	var user *entity.User
	err := db.Where("email = ?", email).First(&user).Error

	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		if err := db.Create(&data).Error; err != nil {
			log.Printf("❌ failed insert user: %v", err)
			return
		}
		log.Println("✅ successfully insert user (insert)")
	case err != nil:
		log.Printf("❌ failed to find user: %v", err)
		return
	default:
		if errUpdate := db.Model(&user).Updates(&data).Error; errUpdate != nil {
			log.Printf("❌ failed to update seed user: %v", err)
			return
		}
		log.Println("✅ successfully update user (update)")
	}
}
