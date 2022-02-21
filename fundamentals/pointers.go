package fundamentals

import "fmt"

func Pointers() {
	fmt.Println("******* Pointers ******")

	// Declare a pointer var by using pointer operator
	var pointerInt *int
	fmt.Println("This is the place pointerInt is poiting:", pointerInt)

	var pointerString *string = new(string)
	// Dereference operation achieved by using dereference operator *
	*pointerString = "Leidy"
	fmt.Println("This is the place pointerString is poiting:", pointerString)
	fmt.Println("This is the value pointerString is poiting:", *pointerString)

	var anotherPointerString = pointerString
	fmt.Println("This is the value anotherPointerString is poiting:", anotherPointerString)
	*anotherPointerString = "Patricia"

	fmt.Println("This is the value pointerString is poiting:", *pointerString)

	// Create pointer using addressOf operator &
	variableString := "dale a tu cuerpo alegria macarena"
	fmt.Println(variableString)

	pointerFromVarAddress := &variableString
	fmt.Println(pointerFromVarAddress, *pointerFromVarAddress)

	variableString = "que tu cuerpo es pa darle alegria cosa buena"
	fmt.Println(pointerFromVarAddress, *pointerFromVarAddress) 
	// Same memory, new value!! -> 0x1400009e080 que tu cuerpo es pa darle alegria cosa buena

}
