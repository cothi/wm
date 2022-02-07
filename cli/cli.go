package cli

import "gopkg.in/alecthomas/kingpin.v2"

var (
	mode       = kingpin.Arg("mode", "set html or rest").Required().String()
	ent        = kingpin.Arg("entropy", "set 256, 158").Required().Int()
	passphrase = kingpin.Arg("passphrase", "").Required().Int()
)

func Start() {
	kingpin.Parse()

	switch *mode {
	case "rest":

	}

}
