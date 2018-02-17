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

	metadata "google.golang.org/grpc/metadata"
	"strings"
	"strconv"
)

type AppleReversiServer struct {
	mu       sync.Mutex
	games	 map[int64]*Room
}

type Room struct {
	gameId  int64
	board   *Board
	states  map[int64]CellState
	members map[CellState]int64
	host    int64
	guest   int64
	changes map[CellState]chan CellChangeSet
}

func (s *AppleReversiServer) CreateGame(ctx context.Context, config *GameConfig) (*GameJoined, error) {
	room := s.newRoom(toCellState(config.Color))
	return &GameJoined{GameId: room.gameId, PlayerId: room.host, Color: config.Color}, nil
}

func (s *AppleReversiServer) JoinGame(ctx context.Context, game *Game) (*GameJoined, error) {
	room, ok := s.games[game.GameId]
	if !ok {
		return nil, errors.New("game id does not exists")
	}
	c := room.states[room.guest]

	return &GameJoined{GameId: room.gameId, PlayerId: room.guest, Color: toColor(c)}, nil
}

func (s *AppleReversiServer) SelectMove(stream Reversi_SelectMoveServer) error {
	md, ok := metadata.FromIncomingContext(stream.Context())
	if !ok {
		return errors.New("Can not get metadata from stream")
	}
	if len(md["gameId"]) == 0 && len(md["playerId"]) == 0 {
		return errors.New("Invalid metadata : required gameId & playerId")
	}

	gameId, err := strconv.Atoi(md["gameId"][0])
	if err != nil {
		return errors.New("Invalid metadata : game id must be numeric")
	}
	playerId, err := strconv.Atoi(md["playerId"][0])
	if err != nil {
		return errors.New("Invalid metadata : game id must be numeric")
	}

	room, ok := s.games[int64(gameId)]
	if !ok {
		return errors.New("game id does not exists")
	}
	state, ok := room.states[int64(playerId)]
	if !ok {
		return errors.New("player id does not exists")
	}

	go func(){
		c := <- room.changes[state]
		if err := stream.Send(&c); err != nil {
			log.Print("Send failed")
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

		mv := &BoardMove{
			color: state,
			row: int(move.Row),
			column: int(move.Column),
		}

		if room.board.CanPlace(mv) {
			s.mu.Lock()
			result := room.board.MakeMove(mv)
			s.mu.Unlock()

			changes := make([]*CellChanged, 0)
			for i := range result {
				c := result[i]
				changes = append(changes, &CellChanged{Row: int32(c.row), Column: int32(c.column)})
			}
			res := CellChangeSet{Move: move, Cells: changes, GameFinished: room.board.HasGameFinished()}

			for k := range room.changes {
				room.changes[k] <- res
			}
			//if err := stream.Send(res); err != nil {
			//	return err
			//}
		}
	}
}

func NewServer() *AppleReversiServer {
	rand.Seed(time.Now().UnixNano())
	return &AppleReversiServer{games: make(map[int64]*Room, 100)}
}

func (s *AppleReversiServer) newRoom(state CellState) *Room {
	gameId := rand.Int63()

	r := &Room{}
	r.gameId = gameId
	r.board = NewBoard()

	r.states = make(map[int64]CellState, 2)
	r.states[rand.Int63()] = cellStateBlack
	r.states[rand.Int63()] = cellStateWhite

	r.members = make(map[CellState]int64, 2)

	for k := range r.states {
		r.members[r.states[k]] = k
	}

	r.host = r.members[state]
	r.guest = r.members[OppnentColor(state)]

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
}

func toColor(state CellState) Color {
	switch state {
	case cellStateBlack:
		return Color_BLACK
	case cellStateWhite:
		return Color_WHITE
	}
}