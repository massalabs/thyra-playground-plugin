package simulator

import (
	"errors"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

func RunSimulation(pluginPath string) error {
	log.Println("Runnig massa sc simulator")

	configFileName := "simulator_config.json"
	simulatorBin := "massa-sc-tester"

	simulatorPath := filepath.Join(pluginPath, "simulator")

	exe := ""
	if runtime.GOOS == "windows" {
		exe = ".exe"
	}

	//check simulator binary
	if _, err := os.Stat(filepath.Join(simulatorPath, simulatorBin+exe)); err != nil {
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

	cmd := exec.Command(filepath.Join(simulatorPath, simulatorBin)+exe, filepath.Join(simulatorPath, configFileName))
	stdout, err := cmd.CombinedOutput()
	if err != nil {
		return errors.New("Smart contract simulation failed: " + string(stdout))
	}

	// move results files in plugin folder
	path, _ := os.Getwd()
	ledgerFile := "ledger.json"
	traceFile := "trace.json"
	os.Rename(filepath.Join(path, ledgerFile), filepath.Join(pluginPath, ledgerFile))
	os.Rename(filepath.Join(path, traceFile), filepath.Join(pluginPath, traceFile))

	return nil
}
