package main

import (
	"log"

	"github.com/haormj/enet-go"
)

func main() {
	if enet.Enet_initialize() != 0 {
		log.Fatalln("enet init failed")
	}
	defer enet.Enet_deinitialize()
	addr := enet.NewENetAddress()
	enet.Enet_address_set_host(addr, "0.0.0.0")
	addr.SetPort(enet.NewEnetUint16(18756))
	srv := enet.Enet_host_create(addr, 100, 2, enet.NewEnetUint32(0), enet.NewEnetUint32(0))
	if srv == nil {
		log.Fatalln("server create failed")
	}
	defer enet.Enet_host_destroy(srv)
	for {
		event := enet.NewENetEvent()
		if enet.Enet_host_service(srv, event, enet.NewEnetUint32(1000)) > 0 {
			switch event.GetXtype() {
			case enet.ENET_EVENT_TYPE_NONE:
				log.Println("none")
			case enet.ENET_EVENT_TYPE_CONNECT:
				log.Println("connect")
			case enet.ENET_EVENT_TYPE_RECEIVE:
				log.Println("receive")
			case enet.ENET_EVENT_TYPE_DISCONNECT:
				log.Println("disconnect")
			default:
				log.Println("hello")
			}
		}
	}

}
