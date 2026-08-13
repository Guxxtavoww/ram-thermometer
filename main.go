package main

import (
	"fmt"
	"ram-thermometer/helpers"
	"ram-thermometer/utils"
	"time"
)

func main() {
	board_port := helpers.UnrwrapError(utils.GetSerialPort())
	serial_port := helpers.UnrwrapError(utils.EstablishBoardComunitation(board_port))

	for {
		usage := helpers.UnrwrapError(utils.GetMemoryUsage())

		message := fmt.Sprintf("%.2f\n", usage)

		_, err := serial_port.Write([]byte(message))

		if err != nil {
			panic(err)
		}

		time.Sleep(1 * time.Second)
	}
}
