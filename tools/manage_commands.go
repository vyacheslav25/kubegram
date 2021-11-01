package tools

import (
	"fmt"
	"os/exec"
)

func GetCommand(kubeclient string, args ...string) string {
	cmd := exec.Command(kubeclient, args...)
	res, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}

	return string(res)

}
