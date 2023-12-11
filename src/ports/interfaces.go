package ports

import (
	model "github.com/ArquitecturaDeSistemas/notificationmicroservice/dominio"
)

// NotificacionRepository define las operaciones que puedes realizar con notificaciones
type NotificacionRepository interface {
	CrearNotificacion(input model.CrearNotificacionInput) (*model.Notificacion, error)
	Notificacion(id string) (*model.Notificacion, error)
	ActualizarNotificacion(id string, input *model.ActualizarNotificacionInput) (*model.Notificacion, error)
	EliminarNotificacion(id string) (*model.RespuestaEliminacion, error)
	Notificaciones() ([]*model.Notificacion, error)
	NotificacionesPorUsuario(usuarioID string) ([]*model.Notificacion, error)
}
