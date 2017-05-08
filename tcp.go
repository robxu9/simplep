package main

import (
	"io"
	"log"
	"net"
	"sync"
)

func makeTCPTunnel(listen, target string) {
	listener, err := net.Listen("tcp", listen)
	if err != nil {
		errAndExit(err)
	}

	go func() {
		for {
			listenerConn, err := listener.Accept()
			if err != nil {
				log.Println(err)
			} else {
				go func() {
					defer listenerConn.Close()
					targetConn, err := net.Dial("tcp", target)
					if err != nil {
						log.Println(err)
						return
					}
					defer targetConn.Close()

					wg := &sync.WaitGroup{}
					wg.Add(2)

					// input
					go func() {
						defer wg.Done()
						_, inErr := io.Copy(targetConn, listenerConn)
						if inErr != nil {
							log.Println(inErr)
						}
					}()

					// output
					go func() {
						defer wg.Done()
						_, outErr := io.Copy(listenerConn, targetConn)
						if outErr != nil {
							log.Println(outErr)
						}
					}()

					wg.Wait()
				}()
			}
		}
	}()
}
