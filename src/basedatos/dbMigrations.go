package basedatos

import (
	"log"
	"notificationmicroservice/schema"

	"gorm.io/gorm"
)

func EjecutarMigraciones(db *gorm.DB) {
	if !db.Migrator().HasTable(&schema.Notification{}) {
		db.AutoMigrate(schema.Notification{})
		log.Println("Migraciones completadas")
	}
	log.Println("Migraciones antigua, Tabla existe")
}
