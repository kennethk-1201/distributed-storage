package main

import (
	"github.com/kennethk-1201/distributed-storage/p2p"
	"log"
)

func OnPeer(peer p2p.Peer) error {
	peer.Close()
	return nil
}

func main() {
	tcpTransportOpts := p2p.TCPTransportOpts{
		ListenAddr:    ":3000",
		HandshakeFunc: p2p.NOPHandshakeFunc,
		Decoder:       p2p.DefaultDecoder{},
		// TODO： onPeer Func
	}

	tcpTransport := p2p.NewTCPTransport(tcpTransportOpts)

	fileServerOpts := FileServerOpts{
		StorageRoot:       "3000_network",
		PathTransformFunc: CASPathTransformFunc,
		Transport:         tcpTransport,
		BootstrapNodes:    []string{":4000"},
	}
	s := NewFileServer(fileServerOpts)
	/*
		go func() {
			time.Sleep(time.Second * 3)
			s.Stop()
		}()
	*/
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
