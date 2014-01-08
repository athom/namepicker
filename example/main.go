package main

import (
	"fmt"
	"github.com/athom/namepicker"
)

func main() {
	names := namepicker.RandomName()
	fmt.Println("RandomName:")
	fmt.Println(names)
	fmt.Println("")

	mnames := namepicker.RandomName()
	fmt.Println("MaleRandomName:")
	fmt.Println(mnames)
	fmt.Println("")

	fnames := namepicker.RandomName()
	fmt.Println("FemaleRandomName:")
	fmt.Println(fnames)
	fmt.Println("")

	fi, la := namepicker.RandomNames()
	fmt.Println("RandomNames:")
	fmt.Println(fi + "." + la)
	fmt.Println("")

	mfi, mla := namepicker.RandomMaleNames()
	fmt.Println("RandomMaleNames:")
	fmt.Println(mfi + "." + mla)
	fmt.Println("")

	ffi, fla := namepicker.RandomFemaleNames()
	fmt.Println("RandomFemaleNames:")
	fmt.Println(ffi + "." + fla)
	fmt.Println("")
	return
}
