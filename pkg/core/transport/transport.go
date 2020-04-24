package transport

import "net/http"

type Transport interface {
	Connect(url string) (conn Connection, err error)
	HandleConnection(w http.ResponseWriter, r *http.Request) (conn Connection, err error)
	Serve(w http.ResponseWriter, r *http.Request)
}
