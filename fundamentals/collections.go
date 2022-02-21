package fundamentals

import "fmt"

func Collections() {
	fmt.Println("******* Collections ******")
	
	// array
	var arrVerbose [3]int
	arrVerbose[0] = 123
	fmt.Println(arrVerbose, arrVerbose[0])

	arr := [3]int{1, 2, 3}
	fmt.Println(arr)

	// slice
	// point to the array, any change in the slice affect the array an viseversa
	slice := arr[:]
	fmt.Println(slice)

	arr[1] = 11111
	slice[2] = 22222
	fmt.Println("slice after mutation from array and viseversa", slice)

	// slice does not have fixed size
	sliceOne := []int{1, 2, 3}
	fmt.Println(sliceOne)
	sliceOne = append(sliceOne, 4, 5, 6)
	fmt.Println(sliceOne)

	fmt.Println(sliceOne[1:], sliceOne[:2], sliceOne[1:2])

	// map
	// declare map with implicit initialization
	m := map[string]int{"one": 1, "two": 2}
	fmt.Println(m, m["one"], m["four"]) //map[one:1 two:2] 1 0

	delete(m, "one")
	fmt.Println(m) //map[two:2]

	// struct -> associate any type of data
	type song struct {
		ID   int
		Name string
	}

	// create song struct
	var songMacarena song
	fmt.Println(songMacarena) // {0 } as songMacarena fields are not defined the get the default values

	songMacarena.ID = 1
	songMacarena.Name = "Macarena"
	fmt.Println(songMacarena)

	// create song short syntax
	songBB := song{2, "Back in black"}
	fmt.Println(songBB)

	songHH := song{Name: "Highway to hell"}
	fmt.Println(songHH)

	songMultiline := song{
		ID:   4,
		Name: "Multiline song",
	}
	fmt.Println(songMultiline)
}
