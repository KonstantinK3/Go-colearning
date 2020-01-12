package envdir

import (
	"io/ioutil"
	"path/filepath"
)

// ReadDir сканирует указанный каталог
// и возвращает все переменные окружения, определенные в нем
func ReadDir(dir string) (map[string]string, error) {

	envs := map[string]string{}

	absPath, err := filepath.Abs(dir)
	if err != nil {
		return nil, err
	}

	files, err := ioutil.ReadDir(absPath)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		buf, err := ioutil.ReadFile(filepath.Join(absPath, file.Name()))
		if err != nil {
			return nil, err
		}
		if len(file.Name()) > 100 || len(string(buf)) > 100 {
			continue
		}
		envs[file.Name()] = string(buf)
		// fmt.Printf("=====%#v\n", envs)
	}

	return envs, nil

}
