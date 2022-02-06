package routing

type DecoratedHandler struct {
	domainName      string
	originalToShort map[string]string
	shortToOriginal map[string]string
}

//Config выполняет первоначальную конфигурацию сервиса и возвращает - имя домена, соответствие  ключа к оригинальному URL
func Config() (DecoratedHandler, error) {
	domainName := "http://localhost:8080"
	originalToShort := make(map[string]string)
	shortToOriginal := make(map[string]string)
	handler := DecoratedHandler{domainName: domainName, originalToShort: originalToShort, shortToOriginal: shortToOriginal}

	return handler, nil
}
