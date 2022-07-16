package syscmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var OsRelease = map[string]string{}

const OsReleasePath = "/etc/os-release"

func GetOsOsReleaseHTML() string {
	osinfo := ParseOsRelease(OsReleasePath)
	result := "<b>Hostsystem: </b>" + osinfo["NAME"] + " <b>OS version:</b> " + osinfo["VERSION_ID"] + "<br/>"
	return result
}

func ParseOsRelease(file string) (osrelease map[string]string) {
	fmt.Println("parsing: " + file)
	var result = make(map[string]string)
	readFile, err := os.Open(file)

	if err != nil {
		fmt.Println(err)
	}

	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		value := strings.Split(line, "=")
		result[value[0]] = strings.Trim(value[1], "\"")
	}

	return result
}
