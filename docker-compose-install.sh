#!/bin/bash

# Verificar si el usuario tiene permisos de superusuario
if [ "$EUID" -ne 0 ]; then
  echo "Este script debe ser ejecutado como superusuario (root)."
  exit 1
fi

# Descargar la versión más reciente de Docker Compose
curl -L "https://github.com/docker/compose/releases/latest/download/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose

# Dar permisos de ejecución a Docker Compose
chmod +x /usr/local/bin/docker-compose

# Crear un enlace simbólico para usar "docker-compose" en lugar de la ruta completa
ln -s /usr/local/bin/docker-compose /usr/bin/docker-compose

# Mostrar mensaje de instalación exitosa
echo "Docker Compose ha sido instalado correctamente."

# Salir del script
exit 0
