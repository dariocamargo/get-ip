package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "CLI App"
	app.Usage = "Returns IP and Server Name from URL"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "github.com/dariocamargo",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Gets IP from URL",
			Flags:  flags,
			Action: getIPs,
		},
		{
			Name:   "servers",
			Usage:  "Gets servers from URL",
			Flags:  flags,
			Action: getServers,
		},
	}

	return app
}

func getIPs(c *cli.Context) {
	host := c.String("host")
	ips, err := net.LookupIP(host)
	if err != nil {
		log.Fatal(err)
	}
	for _, ip := range ips {
		fmt.Println(ip)
	}

}


func getServers(c *cli.Context) {
	host:= c.String("host")
	servers, err := net.LookupNS(host)
	if err != nil {
		log.Fatal(err)
	}
	for _, server := range servers {
		fmt.Println(server.Host)
	}

}