#!/bin/bash

# Verificar si el usuario tiene permisos de superusuario
if [ "$EUID" -ne 0 ]; then
  echo "Este script debe ser ejecutado como superusuario (root)."
  exit 1
fi

# Actualizar el sistema
apt update
apt upgrade -y

# Instalar dependencias
apt install -y apt-transport-https ca-certificates curl software-properties-common

# Agregar el repositorio oficial de Docker
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | gpg --dearmor -o /usr/share/keyrings/docker-archive-keyring.gpg
echo "deb [signed-by=/usr/share/keyrings/docker-archive-keyring.gpg] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable" | tee /etc/apt/sources.list.d/docker.list > /dev/null

# Actualizar el sistema nuevamente
apt update

# Instalar Docker
apt install -y docker-ce

# Iniciar y habilitar el servicio Docker
systemctl start docker
systemctl enable docker

# Agregar el usuario al grupo docker
usermod -aG docker $USER

# Mostrar mensaje de instalación exitosa
echo "Docker ha sido instalado y configurado correctamente."

# Salir del script
exit 0
