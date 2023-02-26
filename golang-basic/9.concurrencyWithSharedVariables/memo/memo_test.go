package memo

import (
	"testing"
	"memo1"
	// "memo2"
)

// go test -run=Test
func Test(t *testing.T) {
	m := memo1.New(HTTPGetBody)
	Sequential(t, m)
}

func TestConcurrent(t *testing.T) {
	m := memo1.New(HTTPGetBody)
	Concurrent(t, m)
}
