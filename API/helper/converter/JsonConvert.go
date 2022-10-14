package converter

import (
	"encoding/json"
	"github.com/guancang10/BookStore/API/helper"
	"net/http"
)

func EncoderToResponse(response http.ResponseWriter, param interface{}) {
	response.Header().Add("content-type", "application/json")
	err := json.NewEncoder(response).Encode(param)
	helper.CheckError(err)
}

func DecoderFromRequest(request *http.Request, param interface{}) {
	err := json.NewDecoder(request.Body).Decode(param)
	helper.CheckError(err)
}
