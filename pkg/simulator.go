package simulator

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func RunSimulation() {
	fmt.Println(" Runnig massa sc simulator")

	configFileName := "simulator_config.json"
	if _, err := os.Stat("./pkg/" + configFileName); err == nil {
		// path/to/whatever exists
		cmd := exec.Command("./pkg/massa-sc-tester", "pkg/"+configFileName)
		stdout, err := cmd.Output()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		// Print the output
		fmt.Println(string(stdout))

	} else {
		log.Fatal("unable to find " + configFileName)
	}

	// delete config
	// os.Remove("./pkg/" + configFileName)

	// Move simulator result to static files
	os.Rename("ledger.json", "./static/ledger.json")
	os.Rename("trace.json", "./static/trace.json")

}
