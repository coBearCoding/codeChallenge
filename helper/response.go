package helper

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Status        int         `json:"status"`
	Data          interface{} `json:"data"`
	Error         string      `json:"error"`
	contentType   string
	controlOrigin string
	allowHeaders  string
	responseWrite http.ResponseWriter
}

func createDefaultResponse(w http.ResponseWriter) Response {

	return Response{
		Status:        http.StatusOK,
		responseWrite: w,
		contentType:   "application/json",
		controlOrigin: "*",
		allowHeaders: "Content-Type",
	}
}

func (response *Response) errorParameter(error string) {
	response.Status = http.StatusBadRequest
	response.Error = error
}

func (response *Response) send() {
	response.responseWrite.Header().Set("Content-Type", response.contentType)
	response.responseWrite.Header().Set("Access-Control-Allow-Origin", response.controlOrigin)
	response.responseWrite.Header().Set("Access-Control-Allow-Headers", response.allowHeaders)
	//response.responseWrite.WriteHeader(response.Status)

	output, _ := json.Marshal(&response)
	fmt.Fprintln(response.responseWrite, string(output))
}

func SendError(w http.ResponseWriter, error string) {
	response := createDefaultResponse(w)
	response.errorParameter(error)
	response.send()
}

func SendData(w http.ResponseWriter, data interface{}) {
	response := createDefaultResponse(w)
	response.Data = data
	response.send()
}
