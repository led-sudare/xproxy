package main

import (
	"time"
	"xproxy/lib/util"

	log "github.com/cihub/seelog"
	zmq "github.com/zeromq/goczmq"
)

func main() {

	xsub := "tcp://*:5510"
	xpub := "tcp://*:5511"
	log.Info("New Sub: ", xsub)
	log.Info("New Pub: ", xpub)

	xpubSock := zmq.NewSock(zmq.Pub)

	xsubSock := zmq.NewSock(zmq.Sub)

	var err error
	_, err = xpubSock.Bind(xpub)
	if err != nil {
		panic(err)
	}
	xsubSock.SetSubscribe("")
	_, err = xsubSock.Bind(xsub)
	if err != nil {
		panic(err)
	}

	defer xsubSock.Destroy()
	defer xpubSock.Destroy()

	ticker := util.NewInlineTicker(2 * time.Second)

	for {
		data, _, err := xsubSock.RecvFrame()
		ticker.DoIfFire(func() {
			log.Info("RecvFrame. last data len: ", len(data))
		})
		if err != nil {
			panic(err)
		}
		xpubSock.SendFrame(data, zmq.FlagNone)
	}

}
