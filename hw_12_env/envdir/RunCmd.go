package envdir

import (
	"log"
	"os"
	"os/exec"
)

// RunCmd запускает программу с аргументами (cmd) c переопределннёым окружением
func RunCmd(cmd []string, env map[string]string) int {

	envSlice := []string{}
	for key, value := range env {
		currentEnv := key + "=" + value
		envSlice = append(envSlice, currentEnv)
	}

	runner := exec.Command(cmd[0])
	// runner.Path = cmd[0]
	runner.Env = envSlice
	runner.Stdout = os.Stdout
	runner.Stderr = os.Stderr

	if err := runner.Run(); err != nil {
		log.Fatalf("error occured: %s\n", err)
	}

	return 0
}
