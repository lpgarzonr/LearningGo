package fundamentals

import (
	"errors"
	"fmt"

	"github.com/lpgarzonr/LearningGo/models"
)

func Functions() {
	fmt.Println("******* Functions ******")

	macarena := models.Song{ID: 1, Name: "Macarena"}
	fmt.Println(macarena)

	dance()
	danceSong(macarena)
	danceTwoSong(macarena, macarena, 1, 2, 3)
	getBoolean()

	// function returning error
	error := startToDance(macarena)
	fmt.Println(error)

	s, error := startToDanceMultipleReturn(macarena)
	fmt.Println("multiple return", s, error)

	_, e := startToDanceMultipleReturn(macarena)
	fmt.Println("ignore first return", e)

	t, _ := startToDanceMultipleReturn(macarena)
	fmt.Println("ignore second return", t)
}

func dance() {
	fmt.Println("Dancing")
}

func danceSong(song models.Song) {
	fmt.Println("Dancing ", song.Name)
}

// Specify param type just once bcs is the same type for all params
func danceTwoSong(songOne, songTwo models.Song, times, hours, steps int) {
	fmt.Println("Dancing ", songOne.Name, "and", songTwo.Name)
}

func getBoolean() bool {
	return true
}

func startToDance(song models.Song) error { // error pointer data type
	fmt.Println("Starting dance ", song.Name)
	return errors.New("Error: I am tired, cant dance")
}

func startToDanceMultipleReturn(song models.Song) (string, error) {
	fmt.Println("Starting dance ", song.Name)
	return song.Name, nil
}
