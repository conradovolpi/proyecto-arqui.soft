-- Tabla de usuarios
CREATE TABLE Usuario (
    UsuarioID INT AUTO_INCREMENT PRIMARY KEY,
    Nombre VARCHAR(100) NOT NULL,
    Email VARCHAR(100) NOT NULL UNIQUE,
    Password VARCHAR(256) NOT NULL,
    Rol VARCHAR(20) NOT NULL
);

-- Tabla de actividades
CREATE TABLE Actividad (
    ActividadID INT AUTO_INCREMENT PRIMARY KEY,
    HorarioInicio VARCHAR(20) NOT NULL,
    HorarioFin VARCHAR(20) NOT NULL,
    Titulo VARCHAR(100) NOT NULL,
    Descripcion TEXT,
    Instructor VARCHAR(100) NOT NULL,
    Duracion INT NOT NULL, -- minutos
    Cupo INT NOT NULL
);

-- Tabla de inscripciones (relación muchos a muchos)
CREATE TABLE Inscripcion (
    UsuarioID INT NOT NULL,
    ActividadID INT NOT NULL,
    FechaInscripcion DATETIME NOT NULL,
    PRIMARY KEY (UsuarioID, ActividadID),
    FOREIGN KEY (UsuarioID) REFERENCES Usuario(UsuarioID) ON DELETE CASCADE,
    FOREIGN KEY (ActividadID) REFERENCES Actividad(ActividadID) ON DELETE CASCADE
);
