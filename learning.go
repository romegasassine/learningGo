package main

import "fmt"
import "time"



func main(){

	fmt.Println("Hello from Rachid's local folder")
	fmt.Println("The time is currently ", time.Now())

	var age int = 25

	fmt.Println("I am ",age, " years old")

	fmt.Println("25 / 5 = ", 25 / 5)

	const lastName = "Sassine"

	fmt.Println("My last name is ", lastName)

	i := 3
	for i >= 25/5{
		fmt.Println(lastName)
	}

	for i := range 10{
		fmt.Println(i)
	}


	if (10 % 8 >= 2) {
	   fmt.Println("Even or greater")
	} else{
	   fmt.Println("Less than or odd this doestnmake any sense")
   	}


	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Its the weekend !")
	case time.Friday, time.Thursday:
		fmt.Println("Getting There!")
	default:
		fmt.Println("its a weekday...")
	}
	

	
	

	//array  now
	var arr [5]int

	arr[4]=3
	arr[1]=64
	arr[3]=0
	arr[2]=25
	arr[0]=700

	for i:=0;i<len(arr); i++ {
		fmt.Println("Array slot : ", i, " element stored : ", arr[i])
		
	}
	//slice that is nil
	var s []string

	//slice with a pre-defined set of numbers using make()
	//var slice3 = make([]string, 3)
	
	s = append(s, "r", "a", "c", "h","i","d")

	fmt.Println(s)
	fmt.Println("Length of S", len(s))
	
	leftSideofS := s[:3]
	rightSideofS := s[3:6]

	fmt.Println("Left side of slice : ", leftSideofS , "\n", "Right side of slice : ", rightSideofS)
	


	//maps use map = make(map[key-type]value-type)
	//for example a black book that holds first names and phone numbers, you can use a key as string, and value type as integer:
	blackBook := make(map[string]int)

	//setting map key value pairs manually
	blackBook["rachid"] = 9176601691
	blackBook["random"] = 123467890

	fmt.Println("Black Book : ", blackBook)
	
	mynumber := blackBook["rachid"]

	fmt.Println("My number should be ", mynumber)
	 //maps also have delete(mapName, keytobedeleted)'
	 //& they have a clear(mapName) method

	 /*in Go, when looking up a key and its value in a Map you can get two results
	   First, the value
	   Second, whether the key actually exists, false if it doesnt
	   so blackBook["nichola"] will return 0, but instead of 0 we can use second placeholder for value, which triggers the second variable which returns a bool whether the key is available or not. If we just want to check if a key is a available or not, and we dont care about storing the value, we can use the "_" placeholder for the value, and store the bool in a present variable or ok variable
*/
	_, isMyGfThere := blackBook["nichola"]
	fmt.Println("Is my gfs number jotted down?", isMyGfThere)
	
	if isMyGfThere == false{
		blackBook["nichola"] = 3471111111
		_, isMyGfThere := blackBook["nichola"]
		
		fmt.Println("Is my gfs number jotted down now?", isMyGfThere)
	}



	
	
}
