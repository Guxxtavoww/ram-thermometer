package utils

import (
	"fmt"

	"github.com/shirou/gopsutil/v3/mem"
)

type MemoryUsage struct {
	TotalGB     float64
	UsedPercent float64
}

func GetMemoryUsage() (*MemoryUsage, error) {
	values, err := mem.VirtualMemory()

	if err != nil {
		return nil, fmt.Errorf("failed to get memory data: %w", err)
	}

	return &MemoryUsage{
		TotalGB:     float64(values.Total) / 1024 / 1024 / 1024,
		UsedPercent: values.UsedPercent,
	}, nil
}
