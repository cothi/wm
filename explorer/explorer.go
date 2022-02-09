package explorer

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/tetgo/wm/mnemonicre"
)

type homeData struct {
	PageTitle string
	Success   bool
}
type url string

var templates *template.Template
var templateDir string = "explorer/templates/"

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
	PageTitle  string `json:"pageTitle"`
	Mnemonic   string `json:"mnemonic"`
	PrivateKey string `json:"privateKey"`
	PublicKey  string `json:"publicKey"`
	Success    bool   `json:"success"`
}

func home(rw http.ResponseWriter, r *http.Request) {
	data := homeData{"home", false}

	switch r.Method {

	case "GET":
		templates.ExecuteTemplate(rw, "home", data)

	case "POST":

		entropyBit, _ := strconv.Atoi(r.FormValue("entropyBit"))
		passphrase := r.FormValue("passphrase")

		a, b, c := mnemonicre.Create(entropyBit, passphrase)
		mnemonicBody := Mnemonic{
			PageTitle:  "home",
			Mnemonic:   a,
			PrivateKey: b,
			PublicKey:  c,
			Success:    true,
		}
		templates.ExecuteTemplate(rw, "home", mnemonicBody)
	}
}

func Start(port int) {
	handler := http.NewServeMux()
	handler.HandleFunc("/", home)
	templates = template.Must(template.ParseGlob(templateDir + "pages/*.gohtml"))
	templates = template.Must(templates.ParseGlob(templateDir + "partials/*.gohtml"))
	fmt.Printf("Listening on localhost:%d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), handler))
}
