package primusdate

import (
	"testing"
)

type validTest struct {
	arg      string
	expected bool
}

var validTests = []validTest{
	validTest{"01.01.2020", true},
	validTest{"31.12.1999", false},
	validTest{"31.12.2039", true},
	validTest{"01.01.2040", false},
	validTest{"32.01.2020", false},
	validTest{"00.01.2020", false},
	validTest{"01.00.2020 ", false},
	validTest{"01.13.2020 ", false},
	validTest{"01.01.2020 ", false},
	validTest{" 01.01.2020", false},
	validTest{"01.01.2020 ", false},
	validTest{"1.1.2020", false},
	validTest{"01-01-2020", false},
	validTest{"2020-01-01", false},
	validTest{"2020/01/01", false},
	validTest{"01/01/2020", false},
	//validTest{"", false},
}

func TestValid(t *testing.T) {

	for _, test := range validTests {
		if output := Valid(test.arg); output != test.expected {
			t.Errorf("Ouput %v not equal to expected %v, argument %s", output, test.expected, test.arg)
		}
	}
}
