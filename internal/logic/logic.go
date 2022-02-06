package logic

import (
	"math/rand"
	"time"
)

/* Каждому длинному URL-адресу присваивается ключ,
который добавляется после http://domain.tld/. К примеру, http://tinyurl.com/m3q2xt имеет ключ m3q2xt.
*/

var letters = []rune("abcdefghijklmnopqrstuvwxyz")

func randSeq(n int) (string, error) {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b), nil
}

// ReturnShortKey возвращает рандомный ключ (используем для генерации коротких URL)
func ReturnShortKey() (string, error) {
	rand.Seed(time.Now().UnixNano())
	return randSeq(5)
}
