package simulator

import (
	"errors"
	"log"
	"os"
	"os/exec"
	"runtime"
)

func RunSimulation() error {
	log.Println(" Runnig massa sc simulator")

	configFileName := "simulator_config.json"
	simulatorBin := "massa-sc-tester"

	//check simulator binary
	if _, err := os.Stat("./simulator/" + simulatorBin); err != nil {
		errMsg := "unable to find simulator " + simulatorBin
		log.Fatal(errMsg)
		return errors.New(errMsg)
	}
	//check simulator config file
	if _, err := os.Stat("./simulator/" + configFileName); err != nil {
		errMsg := "unable to find " + configFileName
		log.Fatal(errMsg)
		return errors.New(errMsg)
	}

	//check simulator binary
	if _, err := os.Stat("./simulator/" + simulatorBin); err != nil {
		errMsg := "unable to find simulator " + simulatorBin
		log.Fatal(errMsg)
		return errors.New(errMsg)
	}

	exe := ""
	if runtime.GOOS == "windows" {
		exe = ".exe"
	}
	cmd := exec.Command("./simulator/massa-sc-tester"+exe, "simulator/"+configFileName)
	stdout, err := cmd.Output()
	// Print the output
	log.Println(string(stdout))
	if err != nil {
		log.Println(" Fatal1")

		log.Fatal(err.Error())
		return err
	}

	// Move simulator result to static files
	os.Rename("ledger.json", "./static/ledger.json")
	os.Rename("trace.json", "./static/trace.json")
	return nil
}
