package source

import (
	"devops-dns-server/config"
	"log"
	"strings"
)

var (
	conf = config.GetConfig()
)

func GetIP(hostname string) string {
	var (
		order   string
		orderS  []string
		address string
	)

	hostname = strings.TrimRight(hostname, ".")

	order = conf.String("source::order")
	orderS = strings.Split(order, ",")

	for _, src := range orderS {

		switch src {
		case "fromAPI":
			address = FromAPI(hostname)
		case "fromFile":
			address = FromFile(hostname)
		default:
			log.Println("not support data source: ", src)
			continue
		}

		if address != "" {
			return address
		}

	}

	return address
}
