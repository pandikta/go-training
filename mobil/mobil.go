package mobil

import "fmt"

type Mobil struct {
	Ban   int
	Pintu int
}

type Ban interface {
	Tipe() string
}

type BanKaret struct {
	Ban
}

func (bk BanKaret) Tipe() string {
	return "Ban Karet"
}

func (m Mobil) Open() {
	fmt.Println("Membuka pintu")
}

func (m Mobil) Knock() {
	fmt.Println("Mengetuk")
}
