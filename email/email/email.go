package email

import (
	"bytes"
	"emailservice/mailjetmail"
	"emailservice/models"
	"encoding/json"
	"net/http"
)

func SendOne(w http.ResponseWriter, r *http.Request) {

	var req models.From

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp, err := mailjetmail.SendEmailWithMailjet(r.Context(), req)
	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(200)
	w.Write(reqBodyBytes.Bytes())
}
