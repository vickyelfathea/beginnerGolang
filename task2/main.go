package main

import ( "strings" 
"fmt"
"math/rand"
"time"
"strconv"
 )

 var runes = []rune("~!@#$%^&*()_+=`:;,.?/") 

func main() {
    var input string 
	fmt.Printf("Input: ")
	fmt.Scan(&input)
	
	var level string 
	fmt.Printf("Level: ")
	fmt.Scan(&level)

	if len(input) < 4 {
		fmt.Println("input harus lebih dari 4 karakter!")
	} else {
		output := genPass(input, level)
		fmt.Printf("Output: %v", output)
	}	
}

func genPass(input string, level string) string{
	hasil:= ""
	switch level {
	
	case "low": 
	hasil = replaceAtIndex(input)
	
	case "med": 
		rand.Seed(time.Now().UnixNano())
    	num := rand.Intn(900) + 100
			if (num<0){
				num *= -1
			}
		hasil = replaceAtIndex(input) + strconv.Itoa(num)

	case "strong":
		rand.Seed(time.Now().UnixNano())
		num := rand.Intn(900) + 100
			if (num<0){
				num *= -1
			}
		hasil = replaceAtIndex(input) + generateRandomRune() + strconv.Itoa(num)

	default: hasil = "masukan level low, med atau strong!"
	
	}

	return hasil
}

func generateRandomRune() string {
	randRune := runes[rand.Intn(len(runes))]
	return string(randRune)
}

func replaceAtIndex(input string) string {
	rand.Seed(time.Now().UnixNano())
    ind := rand.Intn(100)
    n := 0
    hasil := input

    for i := 0; i < ind; {
        rand.Seed(time.Now().UnixNano())
        n = rand.Intn(getLength(input))
		hasil = strings.Replace(hasil, hasil[n:n+1], strings.ToUpper(hasil[n:n+1]) , 1)
        i++
    }   
    return hasil
}

func getLength(input string) int {
    return len(input)-1
}