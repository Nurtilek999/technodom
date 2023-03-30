package handler

import (
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
)

func respond(writer http.ResponseWriter, req *http.Request, code int, responseStruct interface{}) {
	writer.Header().Set("Content-Type", "application/json")

	writer.WriteHeader(code)
	if responseStruct != nil {
		err := json.NewEncoder(writer).Encode(responseStruct)
		if err != nil {
			logrus.WithError(err).Errorf("cannot send response: %+v", responseStruct)
			writer.WriteHeader(http.StatusBadRequest)
		}
	}
}
