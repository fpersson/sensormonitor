package logreader

import (
	"example/user/webserver/model"
	"fmt"
	"os/exec"
	"strings"
)

func ReadLog() (data *model.AllMessages, err error) {
	fmt.Println("Read log...")
	retval := model.AllMessages{}
	prg := "journalctl"
	arg := "-u"
	arg2 := "tempsensor.service"

	cmd := exec.Command(prg, arg, arg2)

	/*
		prg := "cat"
		arg := "./testdata/fejklog.txt"

		cmd := exec.Command(prg, arg)
	*/
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println("wtf: " + err.Error())
		return &retval, err
	}

	v := strings.Split(string(stdout), "\n")

	retval.LogMessages = &v

	return &retval, nil
}

func ReadStatus() (status *model.SystemdStatus) {
	fmt.Println("Read satus")
	retval := model.SystemdStatus{}

	prg := "systemctl"
	arg := "status"
	arg2 := "tempsensor.service"

	cmd := exec.Command(prg, arg, arg2)

	/*
		prg := "cat"
		arg := "./testdata/systemd_fejk_ok.txt"

		cmd := exec.Command(prg, arg)
	*/
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err)
		retval.Active = false
		return &retval
	}

	v := strings.Split(string(stdout), "\n")
	fmt.Println("checking: " + v[2])
	if strings.Contains(v[2], "Active: active") {
		retval.Active = true
	} else {
		retval.Active = false
	}
	retval.Message = &v

	return &retval
}
