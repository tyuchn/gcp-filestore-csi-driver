package main

import (
	"io"
	"log"
	"net"
	"net/rpc"
	"strconv"

	"github.com/prashanthpai/sunrpc"
)

func main() {

	programNumber := uint32(200002)
	programVersion := uint32(1)

	// Filestore instance Ip
	host := "10.224.239.130"
	// GKE Node Ip
	clientNodeIp := "10.128.0.92"

	_ = sunrpc.RegisterProcedure(sunrpc.Procedure{
		ID:   sunrpc.ProcedureID{ProgramNumber: programNumber, ProgramVersion: programVersion, ProcedureNumber: uint32(1)},
		Name: "IN_BAND_PROPRIETARY_LOCK_OPS_PROG.IN_BAND_PROPRIETARY_LOCK_OPS_V1"}, true)

	sunrpc.DumpProcedureRegistry()

	// // Get port from portmapper
	// port, err := sunrpc.PmapGetPort(host, 100021, 4, sunrpc.IPProtoTCP)
	// if err != nil {
	// 	log.Fatalf("sunrpc.PmapGetPort() failed to get port %d: %q", port, err)
	// }

	// Connect to server
	conn, err := net.Dial("tcp", host+":"+strconv.Itoa(int(4045)))
	if err != nil {
		log.Fatal("net.Dial() failed: ", err)
	}

	// Get notified on server closes the connection
	notifyClose := make(chan io.ReadWriteCloser, 5)
	go func() {
		for rwc := range notifyClose {
			conn := rwc.(net.Conn)
			log.Printf("Server %s disconnected", conn.RemoteAddr().String())
		}
	}()

	// Create client using sunrpc codec
	client := rpc.NewClientWithCodec(sunrpc.NewClientCodec(conn, notifyClose))

	log.Printf("Calling IN_BAND_PROPRIETARY_LOCK_OPS_PROG.IN_BAND_PROPRIETARY_LOCK_OPS_V1 to release lock for node ip %q", clientNodeIp)

	err = client.Call("IN_BAND_PROPRIETARY_LOCK_OPS_PROG.IN_BAND_PROPRIETARY_LOCK_OPS_V1", clientNodeIp, nil)
	if err != nil {
		log.Print("client.Call() failed: ", err)
	}
	log.Printf("Lock sucessfully released for node ip %q", clientNodeIp)
}
