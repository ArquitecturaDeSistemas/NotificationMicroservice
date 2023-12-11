package adapters

import (
	"errors"
	"log"

	"github.com/ArquitecturaDeSistemas/notificationmicroservice/database"
	model "github.com/ArquitecturaDeSistemas/notificationmicroservice/dominio"
	"github.com/ArquitecturaDeSistemas/notificationmicroservice/ports"

	"gorm.io/gorm"
)

/**
* Es un adaptador de salida

 */

type notificacionRepository struct {
	db             *database.DB
	activeSessions map[string]string
}

func NewNotificacionRepository(db *database.DB) ports.NotificacionRepository {
	return &notificacionRepository{
		db:             db,
		activeSessions: make(map[string]string),
	}
}

func (r *notificacionRepository) CrearNotificacion(input model.CrearNotificacionInput) (*model.Notificacion, error) {
	if input.UsuarioID == 0 || input.TareaID == 0 || input.Titulo == "" {
		return nil, errors.New("Faltan datos")
	}
	notificacion := model.Notificacion{
		UsuarioID: input.UsuarioID,
		TareaID:   input.TareaID,
		Titulo:    input.Titulo,
	}
	if err := r.db.GetConn().Create(&notificacion).Error; err != nil {
		return nil, err
	}
	return &notificacion, nil
}

func (r *notificacionRepository) Notificacion(id string) (*model.Notificacion, error) {
	if id == "" {
		return nil, errors.New("Falta el id")
	}
	var notificacion model.Notificacion
	if err := r.db.GetConn().First(&notificacion, "id = ?", id).Error; err != nil {
		log.Printf("Error al obtener notificaciones: %v", err)
		return nil, err
	}
	return &notificacion, nil
}

func (r *notificacionRepository) ActualizarNotificacion(id string, input *model.ActualizarNotificacionInput) (*model.Notificacion, error) {
	if id == "" {
		return nil, errors.New("Falta el id")
	}
	notificacion, err := r.Notificacion(id)
	if err != nil {
		return nil, err
	}
	if input.Titulo != "" {
		notificacion.Titulo = input.Titulo
	}
	if err := r.db.GetConn().Save(&notificacion).Error; err != nil {
		return nil, err
	}
	return notificacion, nil
}
func (ur *notificacionRepository) EliminarNotificacion(id string) (*model.RespuestaEliminacion, error) {
	var notificacionGORM model.NotificacionGORM
	result := ur.db.GetConn().First(&notificacionGORM, id)

	if result.Error != nil {
		// Manejo de errores
		if result.Error == gorm.ErrRecordNotFound {
			response := &model.RespuestaEliminacion{
				Mensaje: "La notificación no existe",
			}
			return response, result.Error

		}
		log.Printf("Error al buscar la notificación con ID %s: %v", id, result.Error)
		response := &model.RespuestaEliminacion{
			Mensaje: "Error al buscar la notificación",
		}
		return response, result.Error
	}

	result = ur.db.GetConn().Delete(&notificacionGORM, id)

	if result.Error != nil {
		log.Printf("Error al eliminar la notificación con ID %s: %v", id, result.Error)
		response := &model.RespuestaEliminacion{
			Mensaje: "Error al eliminar la notificación",
		}
		return response, result.Error
	}

	response := &model.RespuestaEliminacion{
		Mensaje: "Notificación eliminada con éxito",
	}
	return response, result.Error

}

func (r *notificacionRepository) Notificaciones() ([]*model.Notificacion, error) {
	var notificaciones []*model.Notificacion
	if err := r.db.GetConn().Find(&notificaciones).Error; err != nil {
		log.Printf("Error al obtener notificaciones: %v", err)
		return nil, err
	}
	return notificaciones, nil
}

func (r *notificacionRepository) NotificacionesPorUsuario(usuarioID string) ([]*model.Notificacion, error) {
	var notificaciones []*model.Notificacion
	if err := r.db.GetConn().Where("usuario_id = ?", usuarioID).Find(&notificaciones).Error; err != nil {
		log.Printf("Error al obtener notificaciones: %v", err)
		return nil, err
	}
	return notificaciones, nil
}
