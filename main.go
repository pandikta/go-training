package main

import (
	"go-training/mobil"
)

func main() {
	// var str string = "asd"
	// var str1 = "a"
	// st2 := "asddd"

	// fmt.Println(str)

	// map
	// var mapBuah = make(map[string]string)
	// mapBuah["buah pasar"] = "mangga"
	// mapBuah["buah manis"] = "jeruk"

	// dataMahasiswa := map[string]string{
	// 	"nama":   "adi",
	// 	"alamat": "bintaro",
	// }

	// fmt.Println(mapBuah)
	// fmt.Println(dataMahasiswa)

	// for mhs := range dataMahasiswa {
	// 	fmt.Println(mhs)
	// }

	// employee := employee.Employee{
	// 	EmpNo: 1,
	// 	Name:  "ari",
	// }

	// fmt.Println(employee.Name)
	// fmt.Println(employee.getName())

	// PendaftaranDataOn(&employee)
	// fmt.Println(employee)

	// x := 10
	// var y *int = &x
	// x = 20

	// fmt.Println(*y)

	// haya := employee.NewProgrammer("asd", "asd")
	// fmt.Println("helo")
	// go func() {
	// 	time.Sleep(100 * time.Millisecond)
	// 	fmt.Println("done")
	// }()

	// wg := &sync.WaitGroup{}
	// makanan := []string{"bakso", "mie ayam", "sosis"}

	// for _, v := range makanan {
	// 	wg.Add(1)
	// 	go func(food string) {
	// 		fmt.Println(food)
	// 		wg.Done()
	// 	}(v)
	// }
	// wg.Wait()

	// var once sync.Once

	// var yangmausayamakan string
	// for _, m := range makanan {
	// 	once.Do(
	// 		func() {
	// 			yangmausayamakan = m
	// 		},
	// 	)
	// }
	// fmt.Println(yangmausayamakan)
	// time.Sleep(200 * time.Millisecond)

	//Race condition
	// mutex
	// var a int
	// wg := &sync.WaitGroup{}
	// mu := &sync.Mutex{}
	// for i := 0; i < 1000; i++ {
	// 	wg.Add(1)
	// 	go func(i int) {
	// 		for j := 0; j < 1000; j++ {
	// 			mu.Lock()
	// 			a++
	// 			mu.Unlock()
	// 		}
	// 		wg.Done()
	// 	}(i)
	// }
	// wg.Wait()
	// fmt.Println(a)

	// ch := make(chan string)
	// defer close(ch)
	// for _, v := range makanan {
	// 	go func(food string) {
	// 		ch <- food
	// 	}(v)
	// }

	// for a := range ch {
	// 	fmt.Println(a)
	// }

	mobil := mobil.Mobil{
		Ban:   4,
		Pintu: 2,
	}

	mobil.Open()
	mobil.Knock()

}

// type Mobil struct {
// 	Ban   int
// 	Pintu int
// }

// type Ban interface {
// 	Tipe() string
// }

// func PendaftaranDataOn(emp *Employee.Employee) {
// 	emp.Name = "budi"
// }

// func (e employee) PakaiPakaian(pakaian Pakaian) {
// 	pakaian.Pakai()
// }
