package main

import "fmt"

func jalanTugas() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover from: ", r)
		}
	}()

	fmt.Println("menjalankan tugas")
	panic("Panic panic panic")
}

func main() {
	jalanTugas()
	fmt.Println("sudah selesai")
}
