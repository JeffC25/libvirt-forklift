package main

import (
	"fmt"

	"libvirt.org/go/libvirt"
)

func main() {
	conn, err := libvirt.NewConnect("qemu:///session")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	doms, err := conn.ListAllDomains(libvirt.CONNECT_LIST_DOMAINS_ACTIVE)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%d running domains:\n", len(doms))
	for _, dom := range doms {
		name, err := dom.GetName()
		if err == nil {
			fmt.Printf("  %s\n", name)
		}
		dom.Free()
	}

	fmt.Println(libvirt.VERSION_NUMBER)
}
