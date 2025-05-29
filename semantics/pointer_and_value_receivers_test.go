package semantics

import "testing"

func TestCounter_Increment(t *testing.T) {
	c := Counter{
		total: 1,
	}

	c.IncrementCounterPointerSemantic()
	c.InspectCounter()
	c.IncrementCounterValueSematic()
	c.InspectCounter()
	c.IncrementCounterPointerSemantic()
	c.InspectCounter()

	UpdateCounterByPassingValue(c)
	c.InspectCounter()
	UpdateCounterByPassingPointer(&c)
	c.InspectCounter()
}
