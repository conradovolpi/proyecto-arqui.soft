# 🐳 Guía de Docker para el Proyecto

Este documento explica cómo usar Docker para ejecutar el proyecto completo.

## 📋 Prerrequisitos

- Docker instalado (versión 20.10 o superior)
- Docker Compose instalado (versión 2.0 o superior)
- Al menos 4GB de RAM disponible
- Puertos 3000, 8080 y 3306 libres

## 🚀 Inicio Rápido

### Opción 1: Scripts Automatizados (Recomendado)

```bash
# Dar permisos de ejecución a los scripts
chmod +x docker-start.sh docker-stop.sh

# Iniciar el proyecto
./docker-start.sh

# Detener el proyecto
./docker-stop.sh
```

### Opción 2: Comandos Manuales

```bash
# Construir y levantar todos los servicios
docker-compose up --build -d

# Ver logs en tiempo real
docker-compose logs -f

# Detener todos los servicios
docker-compose down
```

## 🏗️ Arquitectura del Proyecto

El proyecto está compuesto por 3 servicios:

### 🗄️ MySQL (Base de Datos)
- **Puerto**: 3306
- **Usuario**: appuser
- **Contraseña**: apppass
- **Base de datos**: appdb
- **Volumen persistente**: mysql_data

### 🔧 Backend (API Go)
- **Puerto**: 8080
- **Framework**: Gin
- **Base de datos**: MySQL
- **Health check**: `/health`
- **Endpoints principales**:
  - `POST /usuarios/login` - Login
  - `GET /actividades/` - Listar actividades
  - `POST /actividades/` - Crear actividad (admin)
  - `POST /inscripciones/` - Inscribirse en actividad

### 🌐 Frontend (React + Vite)
- **Puerto**: 3000
- **Framework**: React 18
- **Build tool**: Vite
- **Servidor**: Nginx
- **Proxy API**: Configurado para redirigir llamadas al backend

## 🔧 Comandos Útiles

### Gestión de Servicios

```bash
# Ver estado de todos los servicios
docker-compose ps

# Ver logs de un servicio específico
docker-compose logs backend
docker-compose logs frontend
docker-compose logs mysql

# Reiniciar un servicio específico
docker-compose restart backend

# Reconstruir un servicio específico
docker-compose up --build backend
```

### Debugging

```bash
# Acceder al contenedor del backend
docker-compose exec backend sh

# Acceder al contenedor de MySQL
docker-compose exec mysql mysql -u appuser -papppass appdb

# Ver logs con timestamps
docker-compose logs -t

# Ver logs de los últimos 100 líneas
docker-compose logs --tail=100
```

### Limpieza

```bash
# Detener y eliminar contenedores
docker-compose down

# Detener, eliminar contenedores y volúmenes
docker-compose down -v

# Eliminar imágenes también
docker-compose down --rmi all

# Limpiar sistema Docker (cuidado: elimina todo)
docker system prune -a
```

## 🌐 Acceso a la Aplicación

Una vez iniciado el proyecto:

- **Frontend**: http://localhost:3000
- **Backend API**: http://localhost:8080
- **Base de datos**: localhost:3306

### Endpoints de Health Check

- **Backend**: http://localhost:8080/health
- **Frontend**: http://localhost:3000 (debería mostrar la aplicación)

## 🔒 Configuración de Seguridad

### Variables de Entorno

Las siguientes variables están configuradas en `docker-compose.yml`:

```yaml
# Base de datos
DB_USER: appuser
DB_PASSWORD: apppass
DB_HOST: mysql
DB_PORT: 3306
DB_NAME: appdb

# JWT
JWT_SECRET_KEY: supersecreto123
```

### 🔐 Cambiar Credenciales

Para cambiar las credenciales de la base de datos:

1. Edita `docker-compose.yml`
2. Cambia las variables `MYSQL_*` y `DB_*`
3. Reconstruye los servicios: `docker-compose up --build -d`

## 🐛 Solución de Problemas

### Error de Conexión a Base de Datos

```bash
# Verificar que MySQL esté corriendo
docker-compose logs mysql

# Verificar conectividad
docker-compose exec backend ping mysql
```

### Error de Build

```bash
# Limpiar caché de Docker
docker system prune -a

# Reconstruir sin caché
docker-compose build --no-cache
```

### Puerto en Uso

```bash
# Verificar qué proceso usa el puerto
netstat -tulpn | grep :3000
netstat -tulpn | grep :8080
netstat -tulpn | grep :3306

# Matar proceso si es necesario
sudo kill -9 <PID>
```

### Problemas de Permisos

```bash
# Dar permisos a los scripts
chmod +x docker-start.sh docker-stop.sh

# Si hay problemas con volúmenes
sudo chown -R $USER:$USER .
```

## 📊 Monitoreo

### Verificar Salud de los Servicios

```bash
# Health check del backend
curl http://localhost:8080/health

# Health check del frontend
curl http://localhost:3000

# Estado de contenedores
docker-compose ps
```

### Métricas de Recursos

```bash
# Uso de recursos por contenedor
docker stats

# Información detallada de un contenedor
docker inspect <container_name>
```

## 🔄 Desarrollo

### Modo Desarrollo

Para desarrollo activo, puedes ejecutar los servicios por separado:

```bash
# Solo base de datos
docker-compose up mysql -d

# Ejecutar backend localmente
cd backend && go run main.go

# Ejecutar frontend localmente
cd frontend && npm run dev
```

### Hot Reload

Los Dockerfiles están optimizados para producción. Para desarrollo con hot reload, ejecuta los servicios localmente y solo usa Docker para MySQL.

## 📝 Notas Importantes

1. **Persistencia de Datos**: Los datos de MySQL se guardan en el volumen `mysql_data`
2. **Red Interna**: Los servicios se comunican a través de la red `app-network`
3. **Seguridad**: Los contenedores ejecutan con usuarios no-root
4. **Optimización**: Las imágenes están optimizadas para producción
5. **Health Checks**: Todos los servicios tienen health checks configurados

## 🆘 Soporte

Si encuentras problemas:

1. Revisa los logs: `docker-compose logs`
2. Verifica el estado: `docker-compose ps`
3. Consulta este documento
4. Revisa la configuración de red y puertos
