package greetings

import "testing"

func TestHello(t *testing.T) {
	t.Log(Hello(""))
	t.Log(Hello("John"))
	t.Log(Hello("Jane"))
}
