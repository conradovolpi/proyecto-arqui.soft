package dto

// ActividadRequest representa los datos necesarios para crear o actualizar una actividad
type ActividadRequest struct {
	Titulo        string `json:"titulo" binding:"required"`
	Descripcion   string `json:"descripcion" binding:"required"`
	HorarioInicio string `json:"horario_inicio" binding:"required,datetime=2006-01-02T15:04:05"`
	HorarioFin    string `json:"horario_fin" binding:"required,datetime=2006-01-02T15:04:05"`
	Instructor    string `json:"instructor" binding:"required"`
	Cupo          uint   `json:"cupo" binding:"required,gte=1"`
	Categoria     string `json:"categoria" binding:"required"`
}

// ActividadResponse representa los datos expuestos al cliente
type ActividadResponse struct {
	ID            uint   `json:"id"`
	Titulo        string `json:"titulo"`
	Descripcion   string `json:"descripcion"`
	HorarioInicio string `json:"horario_inicio"`
	HorarioFin    string `json:"horario_fin"`
	Instructor    string `json:"instructor"`
	Cupo          uint   `json:"cupo"`
	Categoria     string `json:"categoria"`
}
