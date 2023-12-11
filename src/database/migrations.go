package database

import (
	"log"

	model "github.com/ArquitecturaDeSistemas/notificationmicroservice/dominio"

	"gorm.io/gorm"
)

// EjecutarMigraciones realiza todas las migraciones necesarias en la base de datos.
func EjecutarMigraciones(db *gorm.DB) {

	db.AutoMigrate(&model.NotificacionGORM{})

	log.Println("Migraciones completadas")
}
