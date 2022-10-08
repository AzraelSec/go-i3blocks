package main

import "strings"

const (
	CHARGING = iota
	DISCHARGING
	UNKNOWN
)

type chargeStatus int

type chargeBlock struct {
	Current int
	Full    int
	Status  chargeStatus
}

func (b *chargeBlock) Percentage() int {
	return b.Current * 100 / b.Full
}

func (b *chargeBlock) Icon() string {
	switch b.Status {
	case CHARGING:
		return "🔌 "
	case DISCHARGING:
		if b.Percentage() < 20 {
			return "🪫 "
		} else {
			return "🔋 "
		}
	default:
		return ""
	}
}

func (b *chargeBlock) TextColor() string {
	p := b.Percentage()
	switch {
	case p < 30:
		return "#d65d0e"
	case p < 20:
		return "#cc241d"
	default:
		return "#fabd2f"
	}
}

func NewChargeStatus(s string) chargeStatus {
	switch strings.Trim(s, "\n") {
	case "Charging":
		return CHARGING
	case "Discharging":
		return DISCHARGING
	default:
		return UNKNOWN
	}
}
