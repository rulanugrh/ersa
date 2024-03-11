package main

import (
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/jedib0t/go-pretty/v6/table"
)

type DNS struct {
	args string
}

func (d *DNS) gotIP() []string {
	var response []string
	lookup, err := net.LookupIP(d.args)
	if err != nil {
		return nil
	}

	for _, data := range lookup {
		result := data.To4()

		response = append(response, result.String())
	}

	return response
}

func (d *DNS) lookupAddr() string {
	ip := d.gotIP()

	var response []string
	for _, dt := range ip {
		lookup, err := net.LookupAddr(dt)
		if err != nil {
			continue
		}

		if lookup == nil {
			return "nil"
		}

		response = append(response, lookup...)
	}

	return strings.Join(response, "\n  ")
}

func (d *DNS) lookupTXT() string {
	lookup, err := net.LookupTXT(d.args)
	if err != nil {
		return "invalid host"
	}

	if lookup == nil {
		return "nil"
	}

	return strings.Join(lookup, "\n  ")
}

func (d *DNS) lookupNS() string {
	lookup, err := net.LookupNS(d.args)
	if err != nil {
		return "invalid host"
	}

	if lookup == nil {
		return "nil"
	}

	var response []string
	for _, dt := range lookup {
		result := dt.Host
		response = append(response, result)
	}

	return strings.Join(response, "\n  ")
}

func (d *DNS) lookupCNAME() string {
	lookup, err := net.LookupCNAME(d.args)
	if err != nil {
		return "invaid host"
	}

	if lookup == "" {
		return "nil"
	}

	return lookup
}

func (d *DNS) lookupMX() string {
	lookup, err := net.LookupMX(d.args)
	if err != nil {
		return "invalid"
	}

	if lookup == nil {
		return "nil"
	}
	var response []string
	for _, dt := range lookup {
		response = append(response, dt.Host)
	}

	return strings.Join(response, "\n  ")
}

func printout() {
	dns := DNS{
		args: os.Args[1],
	}

	t := table.NewWriter()
	t.AppendHeader(table.Row{"", ""})
	t.AppendRow(table.Row{"  IP Address   ", fmt.Sprint("  " + strings.Join(dns.gotIP(), "\n  "))})
	t.AppendRow(table.Row{"  Reverse  ", fmt.Sprint("  " + dns.lookupAddr())})
	t.AppendRow(table.Row{"  TXT   ", fmt.Sprint("  " + dns.lookupTXT())})
	t.AppendRow(table.Row{"  NS    ", fmt.Sprint("  " + dns.lookupNS())})
	t.AppendRow(table.Row{"  CNAME ", fmt.Sprint("  " + dns.lookupCNAME())})
	t.AppendRow(table.Row{"  MX    ", fmt.Sprint("  " + dns.lookupMX())})
	t.AppendFooter(table.Row{"", ""})

	t.SetStyle(table.Style{
		Name: "newStyle",
		Options: table.Options{
			DrawBorder:      false,
			SeparateColumns: false,
			SeparateFooter:  false,
			SeparateHeader:  false,
			SeparateRows:    false,
		},
	})

	fmt.Println(t.Render())
}

func main() {
	printout()
}
