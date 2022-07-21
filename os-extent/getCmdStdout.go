package osxx

import (
	"bufio"
	"log"
	"os/exec"
)

func GetCmdStdout(cmd string) {
	// cmdString := "kubectl -n shencq logs -f --tail 20 k8s-ping-9"
	cmdExec := exec.Command("sh", "-c", cmd)
	out, _ := cmdExec.StdoutPipe()
	scanner := bufio.NewScanner(out)
	go func() {
		for scanner.Scan() {
			log.Println(scanner.Text())
		}

	}()
	cmdExec.Start()
	cmdExec.Wait()

}
