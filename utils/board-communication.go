package utils

import (
	"bufio"
	"fmt"

	"go.bug.st/serial"
)

func EstablishBoardCommunication(board_port string) (serial.Port, error) {
	mode := &serial.Mode{BaudRate: 115200}

	port, err := serial.Open(board_port, mode)

	if err != nil {
		return nil, fmt.Errorf("Unable to open serial port: %d", err)
	}

	defer port.Close()

	// Reader goroutine
	go func() {
		reader := bufio.NewReader(port)

		for {
			line, err := reader.ReadString('\n')

			if err != nil {
				continue
			}

			fmt.Print("ESP32 says: ", line)
		}
	}()

	return port, nil
}
