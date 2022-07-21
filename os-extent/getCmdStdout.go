package osxx

import (
	"bufio"
	"log"
	"os/exec"
)

func GetCmdStdout(cmd string) error {
	// cmdString := "kubectl -n shencq logs -f --tail 20 k8s-ping-9"
	cmdExec := exec.Command("sh", "-c", cmd)
	out, err := cmdExec.StdoutPipe()
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(out)
	go func() {
		for scanner.Scan() {
			log.Println(scanner.Text())
		}

	}()
	err1 := cmdExec.Start()
	if err != nil {
		return err1
	}
	err2 := cmdExec.Wait()
	if err2 != nil {
		return err2
	}

}
