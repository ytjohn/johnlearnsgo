package integers

import (
	"testing"
	"fmt"
)

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}

func ExampleAdd() {
	fmt.Println(Add(1, 5)) // 6
	fmt.Println(Add(5, 4)) // 9
	// Output:
	// 6
	// 9

}
