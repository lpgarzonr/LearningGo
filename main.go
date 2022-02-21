package main

import (
	"fmt"
	"net/http"

	"github.com/lpgarzonr/LearningGo/controllers"
	"github.com/lpgarzonr/LearningGo/fundamentals"
)

func main() {
	fmt.Println("************ Fundamentals ************\n**************************************")
	fundamentals.Variables()
	fundamentals.Pointers()
	fundamentals.Constants()
	fundamentals.Collections()
	fundamentals.Functions()

	fmt.Println("************ Controller ************\n************************")
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
