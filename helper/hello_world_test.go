package helper

import "testing"

func TestHelloWorld(t *testing.T){
	result := HelloWorld("Fuat")
	if result != "Hello Fuat"{
		panic("Result does not match with Hello Fuat")
	}
}
func TestHelloWorldAnshori(t *testing.T){
	result := HelloWorld("Anshori")
	if result != "Hello Anshori"{
		panic("Result does not match with Hello Fuat")
	}
}