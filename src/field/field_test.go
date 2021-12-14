package field

import (
	"fmt"
	"testing"
)

func TestField_Index(t *testing.T) {
	q, r := reverseIndex(110, 10)
	i := index(q, r, 10)
	fmt.Println(i)

	i = index(10, 6, 10)
	q, r = reverseIndex(i, 10)
	fmt.Printf("%d, %d %d\n", i, q, r)
}
