package protocol

import (
	"encoding/json"
	"os"
)

// Note: This should stay aligned with the serialization name
// of `State` field of `I3BlocksOutput` structure in order to
// preserve the cyclic read/update state logic.
const STATE_ENV_NAME = "_state"

// - center
// - right
// - left
type Alignment string

// - Pango
// - None
type MarkupEngine string

type State map[string]string

type I3BlocksOutput struct {
	FullText            string       `json:"full_text,omitempty"`
	ShortText           string       `json:"short_text,omitempty"`
	Color               string       `json:"color,omitempty"`
	Background          string       `json:"background,omitempty"`
	Border              string       `json:"border,omitempty"`
	BorderTop           int          `json:"border_top,omitempty"`
	BorderRight         int          `json:"border_right,omitempty"`
	BorderBottom        int          `json:"border_bottom,omitempty"`
	BorderLeft          int          `json:"border_left,omitempty"`
	MinWidth            int          `json:"min_width,omitempty"`
	Align               Alignment    `json:"align,omitempty"`
	Urgent              bool         `json:"urgent,omitempty"`
	Name                string       `json:"name,omitempty"`
	Instance            string       `json:"instance,omitempty"`
	Separator           bool         `json:"separator,omitempty"`
	SeparatorBlockWidth int          `json:"separator_block_width,omitempty"`
	Markup              MarkupEngine `json:"markup,omitempty"`
	State               State        `json:"_state,omitempty"`
}

func GetState() (State, bool) {
	raw, ok := os.LookupEnv(STATE_ENV_NAME)
	if !ok {
		return make(State), false
	}

	state := map[string]string{}
	if err := json.Unmarshal([]byte(raw), &state); err != nil {
		return make(State), false
	}

	return state, true
}

func (o *I3BlocksOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(*o)
}
