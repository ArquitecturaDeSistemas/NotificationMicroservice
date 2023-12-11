package dominio

type NotificacionGORM struct {
	ID        uint   `gorm:"primaryKey:autoIncrement" json:"id"`
	Titulo    string `gorm:"type:varchar(255);not null"`
	UsuarioID uint   `gorm:"not null"`
	TareaID   uint   `gorm:"not null"`
	// Usuario string `gorm:"foreignKey:UsuarioID;not null"`
	// Tarea string `gorm:"foreignKey:TareaID;not null"`
}

// TableName especifica el nombre de la tabla para UsuarioGORM
func (NotificacionGORM) TableName() string {
	return "notificacions"
}

// func (notificacionGORM *NotificacionGORM) ToGQL() (*Notificacion, error) {
// 	return &Notificacion{
// 		ID:        strconv.Itoa(int(notificacionGORM.ID)),
// 		UsuarioID: strconv.Itoa(int(notificacionGORM.UsuarioID)),
// 		TareaID:   strconv.Itoa(int(notificacionGORM.TareaID)),
// 		Titulo:    notificacionGORM.Titulo,
// 	}, nil
// }
