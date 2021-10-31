package tools

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func GetFiles(command string, args ...string) string {

	cmd := exec.Command(command, args...)
	res, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}

	return string(res)

}

func CatFiles(command, arg string) string {

	cmd := exec.Command(command, arg)
	res, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}

	return string(res)

}

func CreateFolder() {
	if _, err := os.Stat(os.Getenv("FILE_STORAGE")); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(os.Getenv("FILE_STORAGE"), 0755)
		if err != nil {
			log.Println(err)
		}
	}

}

func checkError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
