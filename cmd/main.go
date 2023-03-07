package main

import "fmt"

func main() {
	fmt.Printf("\n\nHello! Wahyu here ^-^ \n\n")

	var i int
	var j bool
	var k float64 = 123.456

	i = 21

	fmt.Println("Menampilakan nilai i : 21")
	fmt.Printf("%v \n\n", i)

	fmt.Println("Menampilkan tipe data dari variabel i")
	fmt.Printf("%T \n\n", i)

	fmt.Println("Menampilkan tanda %")
	fmt.Printf("%% \n\n")

	j = true
	fmt.Println("Menampilkan nilai boolean j : true")
	fmt.Printf("%t \n\n", j)

	j = false
	fmt.Println("Menampilkan nilai boolean j : false")
	fmt.Printf("%t \n\n", j)

	fmt.Println("Menampilkan unicode russia : Я (ya)")
	fmt.Printf("%+q \n\n", "Я")

	fmt.Println("Menampilkan nilai base 10 : 21")
	fmt.Printf("%d \n\n", 21)

	fmt.Println("Menampilkan nilai base 8 : 25")
	fmt.Printf("%o \n\n", 25)

	fmt.Println("Menampilkan nilai base 16 : f")
	fmt.Printf("%x \n\n", "f")

	fmt.Println("Menampilkan nilai base 16 : F")
	fmt.Printf("%X \n\n", "F")

	fmt.Println("Menampilkan unicode karakter Я : U+042F")
	fmt.Printf("\u042F \n\n")

	fmt.Println("Menampilkan float : 123.456000")
	fmt.Printf("%f \n\n", k)

	fmt.Println("Menampilkan float scientific : 1.234560E+02")
	fmt.Printf("%E \n", k)
}