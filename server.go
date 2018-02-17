package applereversi

import (
	_ "fmt"
	"io"
	_ "io/ioutil"
	"log"
	_ "math"
	_ "net"
	"sync"
	_ "time"

	"golang.org/x/net/context"
	_ "google.golang.org/grpc"

	_ "github.com/golang/protobuf/proto"
)

type AppleReversiServer struct {
	mu sync.Mutex
	cpuColor GameConfig_Color
}

func (s *AppleReversiServer) Init(ctx context.Context, config *GameConfig) (*Empty, error) {
	s.cpuColor = config.Color
	return &Empty{}, nil
}

func (s *AppleReversiServer) SelectMove(stream ReversiAI_SelectMoveServer) error {
	for {
		pMove, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		s.mu.Lock()

		// TODO do move
		log.Printf("Player move : %s", pMove)

		cMoves := make([]*Move, 0)

		s.mu.Unlock()

		for _, m := range cMoves {
			if err := stream.Send(m); err != nil {
				return err
			}
		}
	}
}

func NewServer() *AppleReversiServer {
	return &AppleReversiServer{}
}