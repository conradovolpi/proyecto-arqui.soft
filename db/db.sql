CREATE TABLE Usuario (
    UsuarioID INT AUTO_INCREMENT PRIMARY KEY,
    Nombre VARCHAR(100) NOT NULL,
    Email VARCHAR(100) NOT NULL UNIQUE,
    Password VARCHAR(256) NOT NULL,
    Rol VARCHAR(20) NOT NULL
);

CREATE TABLE Actividad (
    ActividadID INT AUTO_INCREMENT PRIMARY KEY,
    HorarioInicio DATETIME NOT NULL,
    HorarioFin DATETIME NOT NULL,
    Titulo VARCHAR(100) NOT NULL,
    Descripcion TEXT,
    Instructor VARCHAR(100) NOT NULL,
    Duracion INT NOT NULL,
    Cupo INT NOT NULL,
    Categoria VARCHAR(50) NOT NULL
);

CREATE TABLE Inscripcion (
    UsuarioID INT NOT NULL,
    ActividadID INT NOT NULL,
    FechaInscripcion DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (UsuarioID, ActividadID),
    FOREIGN KEY (UsuarioID) REFERENCES Usuario(UsuarioID) ON DELETE CASCADE,
    FOREIGN KEY (ActividadID) REFERENCES Actividad(ActividadID) ON DELETE CASCADE
);
