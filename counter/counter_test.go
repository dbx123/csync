package counter

import (
	"testing"
)

func TestCounter(t *testing.T) {

	ctr := New()

	numAdds := 3

	for i := 0; i < numAdds; i++ {
		ctr.Get()
	}

	count := ctr.Get()

	if count != numAdds {
		t.Fatalf("incorrect count expected %v, got %v", numAdds, count)
	}
}
