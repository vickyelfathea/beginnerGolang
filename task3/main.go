package main

import ("fmt"
"strconv")

func main(){ 
	data := []int{1, 7, 3, 4, 8, 9}
	durasi := 10
	search(data, durasi)
}

func search(input []int, durasi int) {
	panjang := len(input)-1
	isi := ""
	for i := panjang; i > 0; {
		for a := i; a > 0; {
			if input [a-1] + input[i] == durasi {
				isi = getFilm(input, a, i)
				fmt.Printf("%v\n", isi)
			}
			a--			
		}		
		isi = ""
		i--
	}
}

func getFilm (input []int, a, i int) string {
	result := strconv.Itoa(input [a-1]) + " dan " + strconv.Itoa(input [i])
	return result
}