package services_errors

import "net/http"

func HTTPErrorHandler(err error, w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(err.Error()))
}
