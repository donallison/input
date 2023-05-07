package input

import (
	"fmt"
	"syscall"

	"golang.org/x/term"
)

type Data struct {
	Value    string
	Original string
	Prompt   string
	Mask     bool
	Loop     bool
}

// GetInput will query the user for additional information needed for the command
func (d *Data) GetInput() (string, error) {
	if d.Mask {
		// Mask input
		for d.Value == "" {
			var bs []byte
			fmt.Print(d.Prompt)
			bs, err := term.ReadPassword(int(syscall.Stdin))
			if err != nil {
				return "", err
			}
			d.Value = string(bs)
			fmt.Printf("\n")
		}
	} else {
		if d.Value == "" {
			fmt.Print(d.Prompt + "(" + d.Original + "): ")
			fmt.Scanln(&d.Value)
			if d.Value == "" && d.Original != "" {
				d.Value = d.Original
			}
		}
	}

	return d.Value, nil
}
