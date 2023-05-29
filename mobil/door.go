package mobil

import "fmt"

type Door struct {
	Side string
}

func (d Door) Open(side string) {
	if side == "Left" {
		fmt.Println("")
	} else {
		fmt.Println("fmt")
	}
}

func (d Door) Knock(side string) {
	if side == "Left" {
		fmt.Println("")
	} else {
		fmt.Println("fmt")
	}
}
