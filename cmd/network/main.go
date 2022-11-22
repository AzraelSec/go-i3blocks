package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/AzraelSec/go-i3blocks/internal/files"
	"github.com/AzraelSec/go-i3blocks/pkg/protocol"
)

const (
	INTERFACE_KEY = "interface"
	DATETIME_KEY  = "last_update"
	RX_KEY        = "last_rx"
	TX_KEY        = "last_tx"

	BASE_PATH   = "/sys/class/net"
	RX_REL_PATH = "statistics/rx_bytes"
	TX_REL_PATH = "statistics/tx_bytes"
)

func main() {
	state, ok := protocol.GetState()
	if !ok {
		protocol.PrintError()
		return
	}

	var network *Network
	name, has := state[INTERFACE_KEY]
	if has {
		found, err := FindNetwork(name)
		if err == nil {
			network = found
		}
	}
	if network == nil {
		found, err := CurrentNetwork()
		if err != nil {
			protocol.PrintError()
			return
		}
		network = found
	}

	intrfPath := fmt.Sprintf("%s/%s", BASE_PATH, network.Interface.Name)
	if _, err := os.Stat(intrfPath); err != nil {
		protocol.PrintError()
		return
	}

	rx, _ := strconv.Atoi(state.GetValue(RX_KEY))
	rx1, _ := files.FileWrapper(fmt.Sprintf("%s/%s", intrfPath, RX_REL_PATH), func(f *os.File) (int, error) {
		return files.GetIntFileValue(f)
	})
	tx, _ := strconv.Atoi(state.GetValue(TX_KEY))
	tx1, _ := files.FileWrapper(fmt.Sprintf("%s/%s", intrfPath, TX_REL_PATH), func(f *os.File) (int, error) {
		return files.GetIntFileValue(f)
	})
	ot, err := parseUnixTime(state.GetValue(DATETIME_KEY))
	if err != nil {
		now := time.Now()
		ot = &now
	}

	drx := rx1 - rx
	dtx := tx1 - tx
	dt := time.Since(*ot).Seconds()

	state[DATETIME_KEY] = strconv.Itoa(int(time.Now().Unix()))
	state[RX_KEY] = strconv.Itoa(rx1)
	state[TX_KEY] = strconv.Itoa(tx1)

	var output *protocol.I3BlocksOutput
	if int(dt) <= 0 {
		output = &protocol.I3BlocksOutput{
			FullText: fmt.Sprintf("[%s] calc...", network.Interface.Name),
			Color:    "#fe8019",
			State:    state,
		}
	} else {
		rxRate, rtRate := float64(drx)/dt, float64(dtx)/dt

		output = &protocol.I3BlocksOutput{
			FullText:  fmt.Sprintf("[%s] 🔻 %s / 🔺 %s", network.Interface.Name, formatRate(rxRate), formatRate(rtRate)),
			ShortText: fmt.Sprintf("🔻 %s / 🔺 %s", formatRate(rxRate), formatRate(rtRate)),
			Color:     "#fe8019",
			State:     state,
		}
	}

	protocol.PrintBlock(output)
}

func parseUnixTime(v string) (*time.Time, error) {
	i, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("impossible to convert %s to valid unix timestamp", v)
	}
	tm := time.Unix(i, 0)
	return &tm, nil
}

func formatRate(r float64) string {
	if int(r) >= 1024 {
		return fmt.Sprintf("%0.1fM", r/1024)
	}
	return fmt.Sprintf("%0.1fkb", r)
}
