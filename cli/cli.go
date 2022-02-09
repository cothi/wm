package cli

import (
	"fmt"
	"os"

	"github.com/tetgo/wm/explorer"
	"github.com/tetgo/wm/mnemonicre"
	"github.com/tetgo/wm/rest"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	app      = kingpin.New("chat", "A command-line mnemonic application.")
	serverIP = app.Flag("server", "Server address").Default("127.0.0.1").IP()

	just  = app.Command("just", "Just make wallet")
	hPass = just.Flag("justPassphrase", "set the port").Short('p').Required().String()
	Ent   = just.Flag("entropy", "set 256, 158").Short('e').Required().Int()

	html  = app.Command("html", "Start html")
	hPort = html.Flag("entBit", "set the port num").Default("6060").Int()

	srest = app.Command("rest", "Start rest")
	rPort = srest.Flag("entBit", "set the port num").Default("6060").Int()
)

func Start() {

	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	case "rest":
		fmt.Printf("rest api start, port %d", *rPort)
		rest.Start(*rPort)

	case "html":
		explorer.Start(*hPort)

	case "just":
		mnemonicre.Create(*Ent, *hPass)

	}

}
