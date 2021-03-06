package humanid

import (
	"math/rand"
	"time"
)

var chars = []rune{
	/*"0"*/ '1', '2', '3', '4', '5', '6', '7', '8', '9',
	'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H' /*'I',*/, 'J', 'K',
	/*'L',*/ 'M', 'N' /*'O'*/, 'P', 'Q', 'R', 'S', 'T', 'U', 'V',
	'W', 'X', 'Y', 'Z',
}

var numchars = []rune{
	'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
}

func New() (id string) {
	rand.Seed(time.Now().UnixNano())
	var vals = make([]rune, 9)
	for i := 0; i < 9; i++ {
		c := chars[rand.Intn(len(chars))]
		if i == 4 {
			vals[i] = '-'
		} else {
			vals[i] = c
		}
	}
	id = string(vals)
	return
}

func NewNumber() (id string) {
	rand.Seed(time.Now().UnixNano())
	var vals = make([]rune, 8)
	vals[0] = numchars[rand.Intn(len(numchars)-1)+1] // no zero for first.
	for i := 1; i < 8; i++ {
		c := numchars[rand.Intn(len(numchars))]
		vals[i] = c
	}
	id = string(vals)
	return
}
