-- Crear tabla Usuario
CREATE TABLE Usuario (
    UsuarioID INT AUTO_INCREMENT PRIMARY KEY,
    Nombre VARCHAR(100) NOT NULL,
    Email VARCHAR(100) NOT NULL UNIQUE,
    Password VARCHAR(256) NOT NULL,
    Rol VARCHAR(20) NOT NULL
);

-- Crear tabla Actividad
CREATE TABLE Actividad (
    ActividadID INT AUTO_INCREMENT PRIMARY KEY,
    HorarioInicio TIME NOT NULL,
    HorarioFin TIME NOT NULL,
    Titulo VARCHAR(100) NOT NULL,
    Descripcion TEXT,
    Instructor VARCHAR(100) NOT NULL,
    Cupo INT NOT NULL,
    Categoria VARCHAR(50) NOT NULL
);

-- Crear tabla Inscripcion (relación entre Usuario y Actividad)
CREATE TABLE Inscripcion (
    UsuarioID INT NOT NULL,
    ActividadID INT NOT NULL,
    FechaInscripcion DATETIME NOT NULL,
    PRIMARY KEY (UsuarioID, ActividadID),
    FOREIGN KEY (UsuarioID) REFERENCES Usuario(UsuarioID) ON DELETE CASCADE,
    FOREIGN KEY (ActividadID) REFERENCES Actividad(ActividadID) ON DELETE CASCADE
);
