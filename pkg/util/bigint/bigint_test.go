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
