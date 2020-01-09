package copy

import (
	"errors"
	"fmt"
	"io"
	"os"
)

//Copy is for copyiing files and stuff. Pretty self-explanatiry
func Copy(src, dst string, offset int, limit int) error {

	// проверка состояния файла
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return fmt.Errorf("%s is not a regular file", src)
	}

	// открытие файла и дефер закрытия
	source, err := os.Open(src)
	if err != nil {
		return err
	}
	defer source.Close()

	// создание файла назначения и дефер закрытия
	destination, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destination.Close()

	// создание буфера на весь файл
	fileSize := int(sourceFileStat.Size())
	buf := make([]byte, sourceFileStat.Size())

	// проверка на смещение
	if offset >= fileSize {
		return errors.New("offset exceedes file size")
	}

	// если лимит слишком большой или ноль - выставим его на конец файла
	if offset+limit > fileSize || limit == 0 {
		limit = fileSize - offset
	}

	// все скопировали в буфер и переписали
	_, err = source.Read(buf)
	if err != nil && err != io.EOF {
		return err
	}

	nBytes, err := destination.Write(buf[offset : offset+limit])

	if err != nil {
		return err
	}

	fmt.Printf("%#v bytes copied", nBytes)

	return nil
}
