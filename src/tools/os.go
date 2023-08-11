package tools

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func GrantVolumePermissions(destinatioonFolder string) (bool, string) {

	// Reemplaza "1000" con el UID deseado
	uid := "1000"

	cmd := exec.Command("sudo", "chown", uid, filepath.Join(destinatioonFolder, "volumen"))

	// Establece la salida estándar y de error para capturar la salida
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Ejecuta el comando
	err := cmd.Run()

	message := "chown completed successfully"
	success := true

	if err != nil {
		message = fmt.Sprintf("Error: %s", err)
	}

	return success, message
}

func ComposeUp(destinationFolder string) (bool, string) {

	message := "docker-compose up -d completed successfully"
	success := true

	// Define el comando y los argumentos
	command := "docker-compose"
	args := []string{"-f", "docker-compose.yml", "up", "-d"}

	// Crea un objeto Cmd para ejecutar el comando
	cmd := exec.Command(command, args...)
	cmd.Dir = destinationFolder // Establece el directorio de trabajo para el comando

	// Conecta la salida estándar y de error del comando al stdout y stderr del programa Go
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Ejecuta el comando
	err := cmd.Run()

	if err != nil {
		message = fmt.Sprintf("Error: %s", err)
		success = false
	}

	return success, message
}
