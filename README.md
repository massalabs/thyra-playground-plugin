# thyra-playground-plugin
Playground plugin for Thyra

This plugin embeds the massa smartcontract playground and serve it has a standalone server.
The backend api is used to execute the smart contract simulator and return the results to the frontend.

# Initialize the development setup:

## Linux

```./install.sh```

# How to build:

## Build plugin for linux
```./build.sh```

This script performs the following actions:
- create a ```build``` directory.
- downloads the simulator binary for the current OS in ```build/simulator```.
- launch the Go build process which produce a ```thyra-playground-plugin``` binary in ```build```
- copy the ```manifest.json``` file in ```build``` 

# How to run the app:
In development mode:

```go run cmd/main.go```

Run the built binary:

```cd build && ./thyra-playground-plugin```

Optionnal flags:
 - --port {int}
