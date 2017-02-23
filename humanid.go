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

func New() (id string) {
	rand.Seed(time.Now().UnixNano())
	var vals []rune
	for i := 0; i < 8; i++ {
		c := chars[rand.Intn(len(chars))]
		if i == 4 {
			vals = append(vals, '-')
		}
		vals = append(vals, c)
	}
	id = string(vals)
	return
}
