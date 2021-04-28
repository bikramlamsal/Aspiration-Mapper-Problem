package mapper

import "testing"

func TestCapitalizeEveryThirdAlphanumericChar(t *testing.T) {
	var testCases = []struct {
		Input  string
		Output string
	}{
		{

			Input:  "Aspiration.com",
			Output: "asPirAtiOn.cOm",
		},
		{

			Input:  "@$piration.com",
			Output: "@$piRatIon.Com",
		},
	}

	for _, v := range testCases {

		t.Run(v.Input, func(t *testing.T) {
			output := CapitalizeEveryThirdAlphanumericChar(v.Input)
			if v.Output != output {
				t.Logf("expected %s, got %s", v.Output, output)
				t.Fail()
			}
		})

	}
}
