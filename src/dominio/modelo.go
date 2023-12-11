package dominio

type Notificacion struct {
	ID        uint   `json:"id"`
	UsuarioID uint   `json:"usuarioId"`
	TareaID   uint   `json:"tareaId"`
	Titulo    string `json:"titulo"`
}

func (Notificacion) IsEntity() {}

type CrearNotificacionInput struct {
	UsuarioID uint   `json:"usuarioId"`
	TareaID   uint   `json:"tareaId"`
	Titulo    string `json:"titulo"`
}

type ActualizarNotificacionInput struct {
	Titulo string `json:"titulo"`
}

type RespuestaEliminacion struct {
	Mensaje string `json:"mensaje"`
}
