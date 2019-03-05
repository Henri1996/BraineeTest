package models

var brainees []Brainee

type Brainee struct {
	id     int
	text   string
	author string
	brand  string
}

func POST(id int, text string, author string, brand string) {
	b1 := Brainee{id, text, author, brand}
	brainees = append(brainees, b1)
}

func GET(id int) Brainee {
	for i := 0; i < len(brainees); i++ {
		if brainees[i].id == id {
			return brainees[i]
			break
		}
	}
	b := Brainee{}
	return b
}
