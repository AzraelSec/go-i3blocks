package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/AzraelSec/go-i3blocks/internal/files"
	"github.com/AzraelSec/go-i3blocks/pkg/protocol"
)

const (
	BASE_PATH        = "/sys/class/power_supply/BAT0"
	CURRENT_REL_PATH = "energy_now"
	FULL_REL_PATH    = "energy_full"
	STATUS_REL_PATH  = "status"
)

func main() {
	file, err := os.Open(BASE_PATH)
	if err != nil {
		fmt.Printf("Directory %s does not exist", BASE_PATH)
		return
	}
	defer file.Close()

	info, err := batteryStat(BASE_PATH)
	if err != nil {
		fmt.Printf("Directory %s does not exist", BASE_PATH)
		return
	}

	if !(*info).IsDir() {
		fmt.Printf("%s entry is not a directory", BASE_PATH)
		return
	}

	current, err := files.FileWrapper(BASE_PATH+"/"+CURRENT_REL_PATH, func(f *os.File) (int, error) {
		return files.GetIntFileValue(f)
	})
	if err != nil {
		fmt.Printf("Impossible to read current battery level")
		return
	}

	full, err := files.FileWrapper(BASE_PATH+"/"+FULL_REL_PATH, func(f *os.File) (int, error) {
		return files.GetIntFileValue(f)
	})
	if err != nil {
		fmt.Printf("Impossible to read full battery level")
		return
	}

	status, _ := files.FileWrapper(BASE_PATH+"/"+STATUS_REL_PATH, func(f *os.File) (string, error) {
		content, err := io.ReadAll(f)
		if err != nil {
			return "", err
		}
		return string(content), nil
	})

	battery := &chargeBlock{
		Full:    full,
		Current: current,
		Status:  NewChargeStatus(status),
	}

	output := &protocol.I3BlocksOutput{
		FullText: fmt.Sprintf("%s%d%%", battery.Icon(), battery.Percentage()),
		Color:    "#fabd2f",
	}
	json, err := json.Marshal(output)
	if err != nil {
		fmt.Print("Error")
		return
	}
	fmt.Println(string(json))
}

func batteryStat(path string) (*os.FileInfo, error) {
	info, err := os.Stat(path)
	if err != nil {
		return nil, err
	}
	return &info, nil
}
