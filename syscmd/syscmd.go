package syscmd

import (
	"fmt"
	"os/exec"
)

const Webservice string = "webserver.service"
const Tempsensorservice string = "tempsensor.service"

func CmdRestart(service string) {
	fmt.Println("Restart: " + service)
	prg := "systemctl"
	action := "restart"

	fmt.Println("cmd: " + prg + " " + action + " " + service)

	command := exec.Command(prg, action, service)

	stdout, err := command.Output()

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(stdout)
	}
}
