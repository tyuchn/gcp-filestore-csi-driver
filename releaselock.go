package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"math/big"
	"net"
	"net/rpc"
	"strconv"

	"github.com/prashanthpai/sunrpc"
)

func IP4toInt(IPv4Address net.IP) int64 {
	IPv4Int := big.NewInt(0)
	IPv4Int.SetBytes(IPv4Address.To4())
	return IPv4Int.Int64()
}

//similar to Python's socket.inet_aton() function
//https://docs.python.org/3/library/socket.html#socket.inet_aton

func Pack32BinaryIP4(ip4Address string) []byte {
	ipv4Decimal := IP4toInt(net.ParseIP(ip4Address))

	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.BigEndian, uint32(ipv4Decimal))

	if err != nil {
		fmt.Println("Unable to write to buffer:", err)
	}

	// // present in hexadecimal format
	// result := fmt.Sprintf("%x", buf.Bytes())
	return buf.Bytes()
}

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

	err = client.Call("IN_BAND_PROPRIETARY_LOCK_OPS_PROG.IN_BAND_PROPRIETARY_LOCK_OPS_V1", Pack32BinaryIP4(clientNodeIp), nil)
	if err != nil {
		log.Print("client.Call() failed: ", err)
	}
	log.Printf("Lock sucessfully released for node ip %q", clientNodeIp)
}
