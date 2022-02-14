package memory

import "testing"


const creations = 20_000_000

type BigStruct struct {
	A, B, C int
	D, E, F string
	G, H, I bool
}
func CreateCopy() BigStruct {
	return BigStruct{
		A: 123, B: 456, C: 789,
		D: "ABC", E: "DEF", F: "HIJ",
		G: true, H: true, I: true,
	}
}

func CreatePointer() *BigStruct {
	return &BigStruct{
		A: 123, B: 456, C: 789,
		D: "ABC", E: "DEF", F: "HIJ",
		G: true, H: true, I: true,
	}
}


func TestCopyIt(t *testing.T) {
	for i := 0; i < creations; i++ {
		_ = CreateCopy()
	}
}

func TestPointerIt(t *testing.T) {
	for i := 0; i < creations; i++ {
		_ = CreatePointer()
	}
}
