package strings

import (
	"fmt"
	"strings"
	"testing"
)

type sli []string

func(s *sli) add(a string){
	*s = append(*s, a)
}

func TestStrings(t *testing.T) {
	a := []byte{'a', 'b', 'c'}
	b := string(a)
	a[1] = 'c'
	t.Log(a)
	t.Log(b)
	fmt.Printf("strings.Clone(b[:1]): %v\n", strings.Clone(b[:1]))

	var s sli
	t.Log(s)
	s.add("aaa")
	t.Log(s)
}
