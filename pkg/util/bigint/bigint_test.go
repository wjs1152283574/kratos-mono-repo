package bigint

import "testing"

func TestParserPrice(t *testing.T) {
	source := map[float64]string{
		6.3: "63000000000000000001",
	}
	for k, v := range source {
		if ParserPrice(k) != v {
			t.Errorf("want: %s,got %s", v, ParserPrice(k))
		}
	}
}

func TestPrecisionPrice6(t *testing.T) {
	if PrecisionPrice6(1) != "1000000" {
		t.Errorf("want: %s,got %s", "1000000", PrecisionPrice6(1))
	}
}
