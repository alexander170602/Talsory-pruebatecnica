# Imagen base de Node.js
FROM node:22-alpine

# Directorio de trabajo dentro del contenedor
WORKDIR /app

# Copiamos los archivos de dependencias
COPY package*.json ./

# Instalamos únicamente dependencias de producción
RUN npm ci --omit=dev

# Copiamos el código fuente
COPY src ./src

# Puerto de la API
ENV PORT=3002

EXPOSE 3002

# Iniciamos el servidor
CMD ["npm", "start"]