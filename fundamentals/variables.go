package fundamentals

import "fmt"

func Variables() {
	fmt.Println("******* Variables ******")

	// Value types

	// Declare var and assign it latter
	var variableInt int
	variableInt = 1
	fmt.Println(variableInt)

	// Declare var and assigned in one line
	var variableString string = "dale a tu cuerpo alegria macarena"
	fmt.Println(variableString)

	// Implicit initialization: Go implicitly assign the type by using (:+)
	implicitFloat := 3.1416
	fmt.Println(implicitFloat)

	implicitBool := true
	fmt.Println(implicitBool)

	// Complex numbers have an specific type
	complexNumber := complex(3, 4)
	fmt.Println(complexNumber)

	// Multiple declarations un one line
	real, imag := real(complexNumber), imag(complexNumber)
	fmt.Println(real, imag)

	// Declare a constant
	const piConstant = 3.1416
	fmt.Println(piConstant)
}
