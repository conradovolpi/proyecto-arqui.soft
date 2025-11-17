# Etapa 1: build de la app con Node
FROM node:20-alpine AS builder

WORKDIR /app

# Copiamos solo el package.json
COPY package.json ./

# Instalamos dependencias sin errores de peer dependencies
RUN npm install --legacy-peer-deps

# Copiamos el resto del proyecto
COPY . .

# Build de producción
RUN npm run build

# Etapa 2: servir el contenido generado con Nginx
FROM nginx:alpine

# Instalar dependencias adicionales si es necesario
RUN apk add --no-cache curl

# Copiar configuración personalizada de nginx
COPY nginx.conf /etc/nginx/conf.d/default.conf

# Borramos el HTML default de nginx
RUN rm -rf /usr/share/nginx/html/*

# Copiamos el contenido compilado
COPY --from=builder /app/dist /usr/share/nginx/html

# Configurar permisos
RUN chown -R nginx:nginx /usr/share/nginx/html

EXPOSE 80

# Healthcheck
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD curl -f http://localhost/ || exit 1

CMD ["nginx", "-g", "daemon off;"]
