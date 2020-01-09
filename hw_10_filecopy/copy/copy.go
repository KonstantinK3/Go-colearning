package copy

import (
	"fmt"
	"io"
	"os"
)

//Copy is for copyiing files and stuff. Pretty self-explanatiry
func Copy(src, dst string, offset int, limit int) error {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)

	fmt.Printf("%#v bytes copied", nBytes)

	return err
}
