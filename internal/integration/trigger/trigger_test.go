package trigger

import "testing"

func TestCalcTrigger(t *testing.T) {

	js := "  (1 == 1)"
	res, err := calcTrigger(js)

	if err != nil {
		t.Fatal(err)
	}
	if !res {
		t.Fatal("expect true")
	}
}
