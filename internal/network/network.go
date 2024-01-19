package network

import (
	"fmt"
	"net"
)

func NSLookup(host string) error {
	ns, err := net.LookupNS(host)
	if err != nil {
		fmt.Println(err)
		return err
	}
	for i := 0; i < len(ns); i++ {
		fmt.Println(ns[i].Host)
	}
	return nil
}

func IPLookup(host string) error {
	IP, err := net.LookupIP(host)
	if err != nil {
		fmt.Println(err)
		return err
	}

	for i := 0; i < len(IP); i++ {
		fmt.Println(IP[i])
	}
	return nil
}

func CnameLookup(host string) error {
	cname, err := net.LookupCNAME(host)
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println(cname)

	return nil
}

func MXLookup(host string) error {
	mx, err := net.LookupMX(host)
	if err != nil {
		fmt.Println(err)
		return err
	}

	for i := 0; i < len(mx); i++ {
		fmt.Println(mx[i].Host, mx[i].Pref)
	}
	return nil
}
