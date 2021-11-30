package primusdate

import (
	"errors"
	"testing"
)

const layoutISO = "2006-01-02"

type validLayoutTest struct {
	arg      string
	expected bool
}

var validLayoutTests = []validLayoutTest{
	validLayoutTest{"01.01.2020", true},
	validLayoutTest{"31.12.1999", false},
	validLayoutTest{"31.12.2039", true},
	validLayoutTest{"01.01.2040", false},
	validLayoutTest{"32.01.2020", false},
	validLayoutTest{"00.01.2020", false},
	validLayoutTest{"01.00.2020 ", false},
	validLayoutTest{"01.13.2020 ", false},
	validLayoutTest{"01.01.2020 ", false},
	validLayoutTest{" 01.01.2020", false},
	validLayoutTest{"01.01.2020 ", false},
	validLayoutTest{"1.1.2020", false},
	validLayoutTest{"01-01-2020", false},
	validLayoutTest{"2020-01-01", false},
	validLayoutTest{"2020/01/01", false},
	validLayoutTest{"01/01/2020", false},
	validLayoutTest{"", false},
	validLayoutTest{"28.02.2022", true},
	validLayoutTest{"29.02.2022", true}, // a leap day
	validLayoutTest{"29.02.2020", true}, // not a leap day
	validLayoutTest{"31.02.2020", true},
}

func TestValid(t *testing.T) {

	for _, test := range validLayoutTests {
		if output := ValidLayout(test.arg); output != test.expected {
			t.Errorf("ouput %v not equal to expected %v, argument %s", output, test.expected, test.arg)
		}
	}
}

type primusDate2Date struct {
	arg         string
	expectedErr error
}

var primusDate2Dates = []primusDate2Date{
	primusDate2Date{"01.01.2000", nil},
	primusDate2Date{"", errors.New("not a valid date")},
	primusDate2Date{"29.02.2022", nil},                            // a leap day
	primusDate2Date{"29.02.2020", errors.New("not a valid date")}, // not a leap day
	primusDate2Date{"31.02.2020", errors.New("not a valid date")},
	primusDate2Date{"01.01.2020", errors.New("not a valid date")},
	primusDate2Date{"31.12.1999", errors.New("not a valid date")},
	primusDate2Date{"31.12.2039", nil},
	primusDate2Date{"01.01.2040", nil},
}

func TestPrimusDate2Date(t *testing.T) {

	for _, test := range primusDate2Dates {

		_, err := PrimusDate2Date(test.arg)
		if test.expectedErr != nil && err != nil {
			if err.Error() != test.expectedErr.Error() {
				t.Errorf("arg: %s error %s not equal expected %s", test.arg, err, test.expectedErr)
			}
		}

	}
}
