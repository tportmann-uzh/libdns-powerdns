package main

import (
	"context"

	"github.com/libdns/libdns"

	"github.com/tportmann-uzh/libdns-powerdns"
)

func main() {
	p := &powerdns.Provider{
		ServerURL: "http://localhost:8081", // required
		ServerID:  "localhost",        // if left empty, defaults to localhost.
		APIToken:  "test",     // required
	}

	_, err := p.AppendRecords(context.Background(), "example.org.", []libdns.Record{
		{
			Name:  "_acme_whatever",
			Type:  "TXT",
			Value: "123456",
			TTL:   300,
		},
	})
	if err != nil {
		panic(err)
	}

}
