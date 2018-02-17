package main

import (
	"flag"
	"log"
	"context"
	"bufio"
	"os"
	"strings"
	"strconv"
	"io"

	"google.golang.org/grpc"
	pb "github.com/thara/applereversi-gameserver"
)

var (
	serverAddr         = flag.String("server_addr", "127.0.0.1:10000", "The server address in the format of host:port")
	serverHostOverride = flag.String("server_host_override", "example.com", "The server name use to verify the hostname returned by TLS handshake")
)

func main() {
	flag.Parse()
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())

	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	client := pb.NewReversiAIClient(conn)

	config := pb.GameConfig{Color: pb.GameConfig_BLACK}
	client.Init(context.Background(), &config)

	stream, err := client.SelectMove(context.Background())
	if err != nil {
		log.Fatalf("%v.RouteChat(_) = _, %v", client, err)
	}

	waitc := make(chan struct{})
	go func() {
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				// read done.
				close(waitc)
				return
			}
			if err != nil {
				log.Fatalf("Failed to receive a move : %v", err)
			}
			log.Printf("Got message move(%d, %d)", in.Row, in.Column)
		}
	}()

	stdin := bufio.NewScanner(os.Stdin)
	for stdin.Scan(){
		text := stdin.Text()
		if text == "exit" {
			log.Print("Bye.")
			break
		}

		data := strings.Split(text, ",")
		if len(data) != 2 {
			log.Print("Input {row},{column}")
		} else {
			row, err := strconv.ParseInt(data[0], 10, 32)
			if err != nil {
				log.Print("Input row as number")
				break
			}
			col, err := strconv.ParseInt(data[1], 10, 32)
			if err != nil {
				log.Print("Input column as number")
				break
			}
			move := pb.Move{Row: int32(row), Column: int32(col)}
			if err := stream.Send(&move); err != nil {
				log.Fatalf("Failed to send a note: %v", err)
			}
		}
	}

	stream.CloseSend()
	<-waitc
}
