package input

import "testing"

func TestGetInput(t *testing.T) {
	d1 := Data{
		Value: "test",
		Mask:  false,
	}
	out, _ := d1.GetInput()

	if out != "test" {
		t.Errorf("Value didn't match")
	}
}
