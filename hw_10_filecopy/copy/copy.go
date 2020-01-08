package copy

import (
	"errors"
	"fmt"
)

//Copy is for copyiing files and stuff. Pretty self-explanatiry
func Copy(from string, to string, offset int, limit int) error {
	fmt.Println("started")

	fmt.Println(from, to, limit, offset)

	return errors.New("Put error here")
}
