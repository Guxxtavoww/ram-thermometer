package utils

import (
	"fmt"

	"go.bug.st/serial/enumerator"
)

func GetSerialPort() (string, error) {
	ports, err := enumerator.GetDetailedPortsList()

	if err != nil {
		panic(err)
	}

	if len(ports) == 0 {
		return "", fmt.Errorf("No serial ports found")
	}

	for _, port := range ports {
		if !port.IsUSB {
			continue
		}

		switch {
		case port.VID == "10C4" && port.PID == "EA60": // CP2102
			return port.Name, nil
		case port.VID == "1A86" && port.PID == "7523": // CH340
			return port.Name, nil
		case port.VID == "10C4" && port.PID == "EA60": // duplicate guard, remove if unused
			return port.Name, nil
		}
	}

	return "", fmt.Errorf("no ESP32-like device found among %d ports", len(ports))
}
