package bigint

import "testing"

func TestParserPrice(t *testing.T) {
	source := map[float64]string{
		6.3: "6300000000000000000",
	}
	for k, v := range source {
		if ParserPrice(k) != v {
			t.Errorf("want: 6000000000000000000,got %s", ParserPrice(k))
		}
	}
}
