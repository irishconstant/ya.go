package routing

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	logic "github.com/irishconstant/ya.go/internal/logic"
)

/*
Сервис для сокращения длинных URL. Требования:
Сервер должен быть доступен по адресу: http://localhost:8080.
Сервер должен предоставлять два эндпоинта: POST / и GET /{id}.
Эндпоинт POST / принимает в теле запроса строку URL для сокращения и возвращает ответ с кодом 201 и сокращённым URL в виде текстовой строки в теле.
Эндпоинт GET /{id} принимает в качестве URL-параметра идентификатор сокращённого URL и возвращает ответ с кодом 307 и оригинальным URL в HTTP-заголовке Location.
Нужно учесть некорректные запросы и возвращать для них ответ с кодом 400.
*/

func RouterStart() {
	handler, err := Config()
	if err != nil {
		log.Fatal("Configuation: ", err)
	}
	http.HandleFunc("/", handler.returnURLHandler)
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func (h DecoratedHandler) returnURLHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		//params := make(map[string]string)
		bodyBytes, _ := ioutil.ReadAll(r.Body)
		bodyString := string(bodyBytes)
		fmt.Println(bodyString)
		/*
			fmt.Println(r.FormValue("url"))
			fmt.Println(r.FormValue(""))
			params["url"] = r.FormValue("url")
			params[""] = r.FormValue("")
		*/
		shortURLKey, _ := h.returnShortURL(bodyString)
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, h.domainName+"/"+shortURLKey)
	case "GET":
		currentID := r.URL.String()
		originalURL, isExist, _ := h.returnOriginalURL(currentID)
		fmt.Println("Оригинальный УРЛ, который вернулся: ", originalURL)
		if isExist {
			//w.Header().Set("Content-Type", "text/html; charset=utf-8")
			w.Header().Set("Location", "http://"+originalURL)
			fmt.Println("Header: ", w.Header())
			w.WriteHeader(http.StatusTemporaryRedirect)
			fmt.Fprintf(w, "%s", originalURL)
		} else {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "%s", "Для данного URL не найден оригинальный URL")
		}
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}

func (h DecoratedHandler) returnShortURL(URL string) (string, error) {
	clearURL := clearURL(URL)
	value, isExist := h.originalToShort[clearURL]
	if isExist {
		return value, nil
	} else {
		key, _ := logic.ReturnShortKey()

		h.originalToShort[clearURL] = key
		h.shortToOriginal[key] = clearURL
		return key, nil
	}
}

func (h DecoratedHandler) returnOriginalURL(shortURL string) (string, bool, error) {
	shortURL = strings.ToLower(shortURL)
	shortURL = strings.ReplaceAll(shortURL, "/", "")
	value, isExist := h.shortToOriginal[shortURL]
	return value, isExist, nil
}

//clearURL очищает URL от http, https, // и т.п.
func clearURL(s string) string {
	s = strings.ReplaceAll(s, "https", "")
	s = strings.ReplaceAll(s, "http", "")
	s = strings.ReplaceAll(s, "/", "")
	s = strings.ReplaceAll(s, ":", "")
	s = strings.ReplaceAll(s, ":", "")
	s = strings.ToLower(s)
	return s
}
