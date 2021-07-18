package module

import (
	"strconv"
)


func SayHello(name string) string {
	return "Hallo " + name
}

func SayHi(name string, umur int) string {

	return "Hai nama saya " + name + " dan umur saya " + strconv.Itoa(umur) 
}