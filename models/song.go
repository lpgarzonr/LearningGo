package models

type Song struct {
	ID   int
	Name string
}

var (
	songs  []*Song
	nextID = 1
)

func getSongs() []*Song {
	return songs
}

func AddSong(song Song) (Song, error) {
	song.ID = nextID
	nextID++
	songs = append(songs, &song)
	return song, nil
}
