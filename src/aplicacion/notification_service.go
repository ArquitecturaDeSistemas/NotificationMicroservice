package service

import (
	"context"
	"fmt"
	"strconv"

	model "github.com/ArquitecturaDeSistemas/notificationmicroservice/dominio"
	repository "github.com/ArquitecturaDeSistemas/notificationmicroservice/ports"
	pb "github.com/ArquitecturaDeSistemas/notificationmicroservice/proto"
)

// este servicio implementa la interfaz UserServiceServer
// que se genera a partir del archivo proto
type NotificacionService struct {
	pb.UnimplementedNotificacionServiceServer
	repo repository.NotificacionRepository
}

func NewNotificacionService(repo repository.NotificacionRepository) *NotificacionService {
	return &NotificacionService{repo: repo}
}

func (s *NotificacionService) CrearNotificacion(ctx context.Context, req *pb.CrearNotificacionRequest) (*pb.NotificacionResponse, error) {
	usuarioID, err := strconv.ParseUint(req.GetUsuarioId(), 10, 32)
	tareaID, err := strconv.ParseUint(req.GetTareaId(), 10, 32)
	if err != nil {
		return nil, err
	}
	input := model.CrearNotificacionInput{
		UsuarioID: uint(usuarioID),
		TareaID:   uint(tareaID),
		Titulo:    req.GetTitulo(),
	}
	notificacion, err := s.repo.CrearNotificacion(input)
	if err != nil {
		return nil, err
	}
	notificacionID := strconv.Itoa(int(notificacion.ID))
	response := &pb.NotificacionResponse{
		Id:        notificacionID,
		UsuarioId: strconv.FormatUint(uint64(notificacion.UsuarioID), 10),
		TareaId:   strconv.Itoa(int(notificacion.TareaID)),
		Titulo:    notificacion.Titulo,
	}
	fmt.Printf("Notificacion creada: %v", response)
	return response, nil
}

func (s *NotificacionService) ObtenerNotificacion(ctx context.Context, req *pb.NotificacionRequest) (*pb.NotificacionResponse, error) {
	notificacion, err := s.repo.Notificacion(req.GetId())
	if err != nil {
		return nil, err
	}
	response := &pb.NotificacionResponse{
		Id:        strconv.Itoa(int(notificacion.ID)),
		UsuarioId: strconv.FormatUint(uint64(notificacion.UsuarioID), 10),
		TareaId:   strconv.Itoa(int(notificacion.TareaID)),
		Titulo:    notificacion.Titulo,
	}
	fmt.Printf("Notificacion obtenida: %v", response)
	return response, nil
}

func (s *NotificacionService) EliminarNotificacion(ctx context.Context, req *pb.EliminarNotificacionRequest) (*pb.RespuestaEliminacion, error) {
	_, err := s.repo.EliminarNotificacion(req.GetId())
	if err != nil {
		return nil, err
	}
	response := &pb.RespuestaEliminacion{
		Mensaje: "Notificacion eliminada",
	}
	fmt.Printf("Notificacion eliminada: %v", response)
	return response, nil
}

func (s *NotificacionService) ListarNotificaciones(ctx context.Context, req *pb.ListarNotificacionesRequest) (*pb.ListarNotificacionesResponse, error) {
	notificaciones, err := s.repo.Notificaciones()
	if err != nil {
		return nil, err
	}
	response := &pb.ListarNotificacionesResponse{}
	for _, notificacion := range notificaciones {
		response.Notificaciones = append(response.Notificaciones, &pb.NotificacionResponse{
			Id:        strconv.Itoa(int(notificacion.ID)),
			UsuarioId: strconv.FormatUint(uint64(notificacion.UsuarioID), 10),
			TareaId:   strconv.Itoa(int(notificacion.TareaID)),
			Titulo:    notificacion.Titulo,
		})
	}
	fmt.Printf("Notificaciones obtenidas: %v", response)
	return response, nil
}

func (s *NotificacionService) NotificacionesPorUsuario(ctx context.Context, req *pb.NotificacionesPorUsuarioRequest) (*pb.NotificacionesPorUsuarioResponse, error) {
	if req.GetUsuarioId() == "" {
		return nil, fmt.Errorf("El usuarioID es requerido")
	}
	notificaciones, err := s.repo.NotificacionesPorUsuario(req.GetUsuarioId())
	if err != nil {
		return nil, err
	}
	response := &pb.NotificacionesPorUsuarioResponse{}
	for _, notificacion := range notificaciones {
		response.Notificaciones = append(response.Notificaciones, &pb.NotificacionResponse{
			Id:        strconv.Itoa(int(notificacion.ID)),
			UsuarioId: strconv.FormatUint(uint64(notificacion.UsuarioID), 10),
			TareaId:   strconv.Itoa(int(notificacion.TareaID)),
			Titulo:    notificacion.Titulo,
		})
	}
	fmt.Printf("Notificaciones obtenidas: %v", response)
	return response, nil
}
