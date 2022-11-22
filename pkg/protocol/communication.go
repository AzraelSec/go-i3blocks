package protocol

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func PrintBlock(output *I3BlocksOutput) {
	WriteBlock(output, os.Stdout)
}

func WriteBlock(output *I3BlocksOutput, w io.Writer) {
	json, err := json.Marshal(output)
	if err != nil {
		panic(err)
	}
	fmt.Fprintln(w, string(json))
}

func Print(s string) {
	Write(s, os.Stdout)
}

func Write(s string, w io.Writer) {
	fmt.Fprintf(w, "{ \"FullText\": \"%s\" }\n", s)
}

func PrintError() {
	PrintBlock(&I3BlocksOutput{
		FullText: "🛑 error",
	})
}
