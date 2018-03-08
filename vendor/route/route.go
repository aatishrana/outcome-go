package route

import (
	"net/http"
	"github.com/gorilla/context"
	"route/middleware/logrequest"
	"router"
	"route/middleware/jsonresponse"
)

func LoadHTTPS() http.Handler {
	return middleware(router.Instance())
}

func LoadHTTP() http.Handler {
	return middleware(router.Instance())
}

func middleware(h http.Handler) http.Handler {

	h = logrequest.Handler(h)

	h = context.ClearHandler(h)

	h = jsonresponse.Handler(h)

	return h
}
