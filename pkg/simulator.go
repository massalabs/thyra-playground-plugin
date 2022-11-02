package simulator

import (
	"errors"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

func RunSimulation() error {
	log.Println(" Runnig massa sc simulator")

	configFileName := "simulator_config.json"
	simulatorBin := "massa-sc-tester"

	path, err := os.Getwd()
	simulatorPath := filepath.Join(path, "simulator")

	//check simulator binary
	if _, err := os.Stat(filepath.Join(simulatorPath, simulatorBin)); err != nil {
		errMsg := "unable to find simulator " + simulatorBin
		log.Fatal(errMsg)
		return errors.New(errMsg)
	}
	//check simulator config file
	if _, err := os.Stat(filepath.Join(simulatorPath, configFileName)); err != nil {
		errMsg := "unable to find " + configFileName
		log.Fatal(errMsg)
		return errors.New(errMsg)
	}

	exe := ""
	if runtime.GOOS == "windows" {
		exe = ".exe"
	}
	cmd := exec.Command(filepath.Join(simulatorPath, simulatorBin)+exe, filepath.Join(simulatorPath, configFileName))
	stdout, err := cmd.Output()
	// Print the output
	log.Println(string(stdout))
	if err != nil {
		log.Println(" Fatal1")

		log.Fatal(err.Error())
		return err
	}

	// Move simulator result to static files
	os.Rename("ledger.json", filepath.Join(path, "static", "ledger.json"))
	os.Rename("trace.json", filepath.Join(path, "static", "trace.json"))
	return nil
}
