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
	"math/rand"
	"time"

	"golang.org/x/net/context"
	_ "google.golang.org/grpc"

	_ "github.com/golang/protobuf/proto"
	"github.com/pkg/errors"

	"google.golang.org/grpc/metadata"
	"strconv"
)

type Server struct {
	mu    sync.Mutex
	games map[int64]*Room
}

type CellState uint8

const (
	cellStateEmpty CellState = iota
	cellStateBlack
	cellStateWhite
)

func OpponentColor(c CellState) CellState {
	switch c {
	case cellStateBlack:
		return cellStateWhite
	case cellStateWhite:
		return cellStateBlack
	default:
		return c
	}
}

type Room struct {
	gameId  int64
	states  map[int64]CellState
	members map[CellState]int64
	host    int64
	guest   int64
	ch      map[int64]chan Move
	mu      sync.Mutex
}

func (r *Room) opponent(playerId int64) int64 {
	if r.host == playerId {
		return r.guest
	} else {
		return r.host
	}
}

func (s *Server) CreateGame(ctx context.Context, config *GameConfig) (*GameJoined, error) {
	room := s.newRoom(toCellState(config.Color))
	room.ch = make(map[int64]chan Move)
	room.ch[room.host] = make(chan Move)
	return &GameJoined{GameId: room.gameId, PlayerId: room.host, Color: config.Color}, nil
}

func (s *Server) JoinGame(ctx context.Context, game *Game) (*GameJoined, error) {
	room, ok := s.games[game.GameId]
	if !ok {
		return nil, errors.New("game id does not exists")
	}
	c := room.states[room.guest]
	room.ch[room.guest] = make(chan Move)
	return &GameJoined{GameId: room.gameId, PlayerId: room.guest, Color: toColor(c)}, nil
}

func (s *Server) SelectMove(stream Reversi_SelectMoveServer) error {
	md, ok := metadata.FromIncomingContext(stream.Context())
	if !ok {
		return errors.New("Can not get metadata from stream")
	}
	log.Printf("%v", md)
	if len(md["game-id"]) == 0 && len(md["player-id"]) == 0 {
		return errors.New("Invalid metadata : required gameId & playerId")
	}

	gameId, err := strconv.Atoi(md["game-id"][0])
	if err != nil {
		return errors.New("Invalid metadata : game id must be numeric")
	}
	playerId, err := strconv.ParseInt(md["player-id"][0], 10, 64)
	if err != nil {
		return errors.New("Invalid metadata : game id must be numeric")
	}

	room, ok := s.games[int64(gameId)]
	if !ok {
		return errors.New("game id does not exists")
	}

	go func() {
		for {
			select {
			case mv := <-room.ch[playerId]:
				if err := stream.Send(&mv); err != nil {
					log.Print("Send failed")
				}
			case <-stream.Context().Done():
				log.Print("Consumer end")
				return
			}
		}
	}()

	for {
		move, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		log.Printf("Player move : %s", move)

		s.mu.Lock()
		room.ch[room.opponent(playerId)] <- *move
		s.mu.Unlock()
	}
}

func NewServer() *Server {
	rand.Seed(time.Now().UnixNano())
	return &Server{games: make(map[int64]*Room, 100)}
}

func (s *Server) newRoom(state CellState) *Room {
	gameId := rand.Int63()

	r := &Room{}
	r.gameId = gameId

	r.states = make(map[int64]CellState, 2)
	r.states[rand.Int63()] = cellStateBlack
	r.states[rand.Int63()] = cellStateWhite

	r.members = make(map[CellState]int64, 2)

	for k := range r.states {
		r.members[r.states[k]] = k
	}

	r.host = r.members[state]
	r.guest = r.members[OpponentColor(state)]

	s.games[gameId] = r
	return r
}

func toCellState(c Color) CellState {
	switch c {
	case Color_BLACK:
		return cellStateBlack
	case Color_WHITE:
		return cellStateWhite
	}
	return cellStateEmpty
}

func toColor(state CellState) Color {
	switch state {
	case cellStateBlack:
		return Color_BLACK
	case cellStateWhite:
		return Color_WHITE
	}
	return Color_BLACK
}
