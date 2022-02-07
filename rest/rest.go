package rest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/woong-s/mnemoniv/mnemonicre"
	"github.com/woong-s/mnemoniv/utils"
)

var port string

type (
	url string
)

type mnemonicBody struct {
	EntropyBit string
	Passphrase string
}

type urlDescription struct {
	URL         url    `json:"url"`         // url
	Method      string `json:"method"`      // json
	Description string `json:"description"` // Description
	Payload     string `json:"payload"`     // payload
}
type Mnemonic struct {
	Mnemonic   string `json:"mnemonic"`
	PrivateKey string `json:"privateKey"`
	PublicKey  string `json:"publicKey"`
}

func documentation(rw http.ResponseWriter, r *http.Request) {
	data := []urlDescription{
		{
			URL:         url("/"),
			Method:      "GET",
			Description: "See Documentation",
		},
		{
			URL:         url("/mnemonic"),
			Method:      "POST",
			Description: "make mnemonic, entropyBit:string,passphrase:string",
		},
	}
	json.NewEncoder(rw).Encode(data)
}

func mnemonic(rw http.ResponseWriter, r *http.Request) {

	var mnemonicBody mnemonicBody
	utils.HandleErr(json.NewDecoder(r.Body).Decode(&mnemonicBody))

	entBit, _ := strconv.Atoi(mnemonicBody.EntropyBit)
	a, b, c := mnemonicre.Create(entBit, mnemonicBody.Passphrase)

	data := Mnemonic{
		Mnemonic:   a,
		PrivateKey: b,
		PublicKey:  c,
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusCreated)
	json.NewEncoder(rw).Encode(data)

}

func Start(aPort int) {
	port = fmt.Sprintf(":%d", aPort)
	router := mux.NewRouter()
	router.HandleFunc("/", documentation).Methods("GET")
	router.HandleFunc("/mnemonic", mnemonic).Methods("POST")
	log.Fatal(http.ListenAndServe(port, router))
}
