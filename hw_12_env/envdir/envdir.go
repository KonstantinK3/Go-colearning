package envdir

// ReadDir сканирует указанный каталог
// и возвращает все переменные окружения, определенные в нем
func ReadDir(dir string) (map[string]string, error) {

	return map[string]string{}, nil

}

// RunCmd запускает программу с аргументами (cmd) c переопределнным окружением
func RunCmd(cmd []string, env map[string]string) int {
	return 0
}
