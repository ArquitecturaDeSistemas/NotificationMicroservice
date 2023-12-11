package schema

import "time"

type Notification struct {
	Titulo       string    `gorm:"not null" json:"titulo"`
	UsuarioID    string    `gorm:"not null;primaryKey" json:"usuario_id"`
	TareaID      string    `gorm:"not null;primaryKey" json:"tarea_id"`
	FechaTermino time.Time `gorm:"not null" json:"fecha_termino"`
	Status       int       `gorm:"not null;default:0"` //0 sin revisar, 1 activo, 2 terminado
}

type Notifications []Notification
