package gotelnet

import (
	"fmt"
	"testing"
)

func TestTelnet(t *testing.T) {
	client := New("127.0.0.1", 9200)
	err := client.Connect()
	if err != nil {
		println(err.Error())
	}
	println("cccccc")
	defer client.Close()
	println("bbbbbb")
	println("aaaaa", client.MakeServerType())
	client.Close()
}

func TestByte(t *testing.T) {
	fmt.Printf("%v", Closed)
	fmt.Printf("%v", UnauthorizedAccess)
	fmt.Printf("%v", OnlyPassword)
	fmt.Printf("%v", UsernameAndPassword)
}
