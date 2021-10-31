// Run with go test -bench=.
package echo_test

import (
	"fmt"
	"strings"
	"testing"
)

var args = []string{"teste1", "teste2", "teste3", "teste4", "teste5", "teste6",
	"teste7", "teste8", "teste9"}

func echo1(args []string) {

	var s, sep string

	for i := 1; i < len(args); i++ {
		s += sep + args[i]
		sep = " "
	}

	fmt.Println(s)
}

func echo2(args []string) {

	s, sep := "", ""

	for _, arg := range args[1:] {
		s += sep + arg
		sep = " "
	}

	fmt.Println(s)
}

func echo3(args []string) {

	fmt.Println(strings.Join(args[1:], " "))
}

func BenchmarkEcho1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echo1(args)
	}
}

func BenchmarkEcho2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echo2(args)
	}
}
func BenchmarkEcho3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echo3(args)
	}

}
