package cli

import (
	"fmt"

	"github.com/woong-s/mnemoniv/rest"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	Mode       = kingpin.Flag("mode", "set html, rest, make").Short('m').Required().String()
	Passphrase = kingpin.Flag("passphrase", "set the port").Short('p').Int()
	Ent        = kingpin.Flag("entropy", "set 256, 158").Short('e').Int()
	Port       = kingpin.Flag("entBit", "set the port num").Default("6060").Int()
)

func Start() {

	kingpin.Parse()
	fmt.Println(*Mode, *Ent, *Port, *Passphrase)

	switch *Mode {
	case "rest":
		fmt.Printf("rest api start, port %d", *Port)
		rest.Start(*Port)
	case "html":

	case "make":
	}
}
