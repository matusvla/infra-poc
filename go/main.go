package main

import (
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

const serviceAddr = ":8081"

type LogHandler struct{}

func (l LogHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	body, _ := ioutil.ReadAll(request.Body)
	log.WithField("req_body", body).WithField("req_header", request.Header).Info("received a request")
	writer.WriteHeader(http.StatusAccepted)
	writer.Write([]byte("request accepted"))
}

func main() {
	if err := http.ListenAndServe(serviceAddr, LogHandler{}); err != nil {
		log.WithError(err).Info("server stopped")
		return
	}
}
