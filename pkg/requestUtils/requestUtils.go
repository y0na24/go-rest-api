package requestUtils

import (
	"encoding/json"
	resUtils "go-rest-api/pkg/responseUtils"
	"io"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func HandleBody[Payload any](w http.ResponseWriter, req *http.Request) (*Payload, error) {
	payload, err := Decode[Payload](req.Body)

	if err != nil {
		resUtils.JsonResponse(w, err.Error(), 402)
		return nil, err
	}

	err = isPayloadValid(payload)

	if err != nil {
		resUtils.JsonResponse(w, err.Error(), 402)
		return nil, err
	}

	return &payload, nil
}

func Decode[T any](body io.ReadCloser) (T, error) {
	var payload T

	err := json.NewDecoder(body).Decode(&payload)

	if err != nil {
		return payload, err
	}

	return payload, nil
}

func isPayloadValid[T any](payload T) error {
	validate := validator.New()
	err := validate.Struct(payload)
	return err
}
