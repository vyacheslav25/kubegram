package tools

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func RewriteK3() string {
	srcFile, err := os.Open(os.Getenv("K3PATH"))
	checkError(err)
	defer srcFile.Close()

	destFile, err := os.Create(os.Getenv("KUBECONFIG"))
	checkError(err)
	defer destFile.Close()

	_, err = io.Copy(destFile, srcFile)
	checkError(err)

	err = destFile.Sync()
	checkError(err)

	
	kubeconfig := os.Getenv("KUBECONFIG")
	result := grepKubeConfig(kubeconfig)

	return result

}

func grepKubeConfig(string) string {
	kubeconfig, err := os.Open(os.Getenv("KUBECONFIG"))
	checkError(err)

	defer kubeconfig.Close()

	scanner := bufio.NewScanner(kubeconfig)
	temp := []string{}
	var result string

	for i := 1; scanner.Scan(); i++ {
		if strings.Contains(scanner.Text(), "server:") {
			temp = append(temp, scanner.Text())
			result = strings.Join(temp, " ")
		}
	}
	return "Your kubeconfig:\n" + result

}

func MyKubeConfig() string {
	kubeconfig := os.Getenv("KUBECONFIG")
	result := grepKubeConfig(kubeconfig)

	return result
}
