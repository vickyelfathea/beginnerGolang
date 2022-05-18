package main

import ("fmt")

func printSegitiga (s int){
	var bintang string
	var spasi string

	for i := s; i > 0; {
		for a := i*2-1; a > 0; {
			bintang += "*"
			a --
		}
		spasi += " "
		fmt.Printf("%v%v\n", spasi, bintang)
		bintang = ""
		i--
	}
}

func main(){ 
	printSegitiga(7)
}