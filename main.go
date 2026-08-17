package main

import (
	"fmt"
	"ram-thermometer/helpers"
	"ram-thermometer/utils"
	"time"
)

func run() {
	board_port := helpers.UnrwrapError(utils.GetSerialPort())
	serial_port := helpers.UnrwrapError(utils.EstablishBoardCommunication(board_port))

	defer serial_port.Close()

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

func main() {
	for {
		func() {
			defer func() {
				if err := recover(); err != nil {
					fmt.Printf("Program failed: %v. Restarting...\n", err)

				}
			}()

			run()
		}()

		time.Sleep(1 * time.Second)
	}
}
