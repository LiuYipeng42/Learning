// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package memo

import (
	"testing"
)

func Test(t *testing.T) {
	m := New(httpGetBody)
	Sequential(t, m)
}

// NOTE: not concurrency-safe!  Test fails.
// 运行  go test -race -run TestConcurrent 就可能检测出竞态
func TestConcurrent(t *testing.T) {
	m := New(httpGetBody)
	Concurrent(t, m)
}
