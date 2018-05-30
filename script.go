package main

import (
	"bytes"
	"os/exec"
	"os"
	"fmt"
	"log"
)

//AllLocations will return locations
func RunScript(script string) (error, string, string) {

	curd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	//os.Chdir("scripts")


	fmt.Println("CHECKING "+curd+"/scripts/"+script+".sh")
	//CHECK IF FILE EXIST otherwise return error... so one cannot pass malicious bash -c cmd
	if _, err := os.Stat(curd+"/scripts/"+script+".sh"); os.IsNotExist(err) {
		fmt.Print(err.Error())
		return err, "no such file or dir or smh", "just no"
	}

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command(curd+"/scripts/"+script+".sh", "", ) // "-c", script)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	errun := cmd.Run()
	return errun, stdout.String(), stderr.String()
}
