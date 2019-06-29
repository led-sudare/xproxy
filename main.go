package main

import (
	log "github.com/cihub/seelog"
	zmq "github.com/zeromq/goczmq"
)


func main() {


	xsub := "tcp://*:5510"
	xpub := "tcp://*:5511"
	log.Info("New Sub: ", xsub)
	log.Info("New Pub: ", xpub)


	xpubSock:= zmq.NewSock(zmq.Pub)

	xsubSock := zmq.NewSock(zmq.Sub)

	var err error
	_,err = xpubSock.Bind(xpub)
	if err != nil {
		panic(err)
	}
	xsubSock.SetSubscribe("")
	_,err = xsubSock.Bind(xsub)
	if err != nil {
		panic(err)
	}

	defer xsubSock.Destroy()
	defer xpubSock.Destroy()

	for{
		data, _, err := xsubSock.RecvFrame()
		log.Info("RecvFrame.", len(data))
		if err != nil{
			panic(err)
		}
		xpubSock.SendFrame(data, zmq.FlagNone)
	}

}
