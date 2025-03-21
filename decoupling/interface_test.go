package decoupling

import "testing"

func TestRetriever(t *testing.T) {
	/*
		var f file
		err := Retriever(f)
		if err != nil {
			return
		}
	*/

	f := file{"data.json"}
	err1 := Retriever(f)
	if err1 != nil {
		return
	}

	p := pipe{"cfg_service"}
	err2 := Retriever(p)
	if err2 != nil {
		return
	}
}
