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
	"google.golang.org/grpc/metadata"
)

var (
	serverAddr         = flag.String("server_addr", "127.0.0.1:10000", "The server address in the format of host:port")
	serverHostOverride = flag.String("server_host_override", "example.com", "The server name use to verify the hostname returned by TLS handshake")

	guest = flag.Bool("guest", false, "")
	gameId = flag.Int64("gameId", 0, "")
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

	client := pb.NewReversiClient(conn)

	var joined *pb.GameJoined
	if *guest {
		log.Print("Guest : JoinGame")
		game := pb.Game{ GameId: *gameId}
		joined, err = client.JoinGame(context.Background(), &game)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		log.Print("Host : CreateGame")
		config := pb.GameConfig{Color: pb.Color_BLACK}
		joined, err = client.CreateGame(context.Background(), &config)
		if err != nil {
			log.Fatal(err)
		}
	}
	log.Printf("playerId: %d, gameId : %d", joined.PlayerId, joined.GameId)

	md := metadata.Pairs("player-id", strconv.FormatInt(joined.PlayerId, 10), "game-id", strconv.FormatInt(joined.GameId, 10))
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	stream, err := client.SelectMove(ctx)
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
