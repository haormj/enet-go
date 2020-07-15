package main

import (
	"log"
	"time"

	"github.com/haormj/enet-go"
)

func main() {
	if enet.Enet_initialize() != 0 {
		log.Println("enet init failed")
		return
	}
	defer enet.Enet_deinitialize()

	cli := enet.Enet_host_create(enet.NewENetAddress(), 1, 2, enet.NewEnetUint32(0), enet.NewEnetUint32(0))
	if cli == nil {
		log.Println("client create failed")
		return
	}
	defer enet.Enet_host_destroy(cli)

	addr := enet.NewENetAddress()
	enet.Enet_address_set_host(addr, "127.0.0.1")
	addr.SetPort(enet.NewEnetUint32(18756))
	peer := enet.Enet_host_connect(cli, addr, 2, enet.NewEnetUint32(0))
	if peer == nil {
		log.Println("enet host connect error")
		return
	}
	connCh := make(chan struct{})
	go func() {
		<-connCh
		for {
			buff := []byte("hello world, I'm client")
			buffPtr, buffLen := enet.BytesToUintptr(buff)
			// packet := enet.Enet_packet_create(buffPtr, int64(buffLen),
			// 	enet.NewEnetUint32(uint32(enet.ENET_PACKET_FLAG_RELIABLE)))
			packet := enet.NewENetPacket()
			packet.SetData(enet.SwigcptrEnet_uint8(buffPtr))
			packet.SetDataLength(int64(buffLen))

			flags := []uint32{uint32(enet.ENET_PACKET_FLAG_RELIABLE)}
			flagsPtr, _ := enet.Uint32BytesToUintptr(flags)
			packet.SetFlags(enet.SwigcptrEnet_uint32(flagsPtr))
			if packet == nil {
				log.Println("enet packet create failed")
				continue
			}
			if ret := enet.Enet_peer_send(peer, enet.NewEnetUint8(0), packet); ret != 0 {
				log.Println("enet peer send failed")
			}
			enet.DeleteENetPacket(packet)
			time.Sleep(time.Second)
		}
	}()
	event := enet.NewENetEvent()
	for {
		if enet.Enet_host_service(cli, event, enet.NewEnetUint32(1000)) > 0 {
			switch event.GetXtype() {
			case enet.ENET_EVENT_TYPE_NONE:
				log.Println("none")
			case enet.ENET_EVENT_TYPE_CONNECT:
				select {
				case connCh <- struct{}{}:
				default:
				}
				log.Println("connect")
			case enet.ENET_EVENT_TYPE_RECEIVE:
				data := enet.UintptrToBytes(event.GetPacket().GetData().Swigcptr(),
					int(event.GetPacket().GetDataLength()))
				log.Println("receive: ", string(data))
			case enet.ENET_EVENT_TYPE_DISCONNECT:
				enet.Enet_peer_disconnect(peer, enet.NewEnetUint32(0))
				log.Println("disconnect")
			}
		}
	}
}
