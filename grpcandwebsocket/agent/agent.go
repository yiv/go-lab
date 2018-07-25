package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/satori/go.uuid"
	"google.golang.org/grpc"

	"github.com/yiv/go-lab/grpcandwebsocket/pb"
)

type sid string

type Session struct {
	id            sid
	conn          *websocket.Conn
	rpcStream     pb.GMService_StreamClient
	toServiceChan chan sid
}

type AgentService struct {
	RPCClient   pb.GMServiceClient
	SessDieChan chan sid
	Sessions    map[sid]*Session
}

func NewAgentService(rpcClient pb.GMServiceClient) AgentService {
	agent := AgentService{
		RPCClient: rpcClient,
		Sessions:  make(map[sid]*Session),
	}
	return agent
}
func (s *Session) ForwardToClient() {
	fmt.Println("start to Forward To Client......")
	for {
		frame, err := s.rpcStream.Recv()
		if err == io.EOF || err != nil {
			s.toServiceChan <- s.id
			s.rpcStream.CloseSend()
			return
		}

		err = s.conn.WriteMessage(websocket.BinaryMessage, frame.Payload)
		if err != nil {
			s.toServiceChan <- s.id
			s.conn.Close()
			return
		}
	}
}
func (s *Session) ForwardToServer() {
	fmt.Println("start to Forward To Server......")
	//s.conn.SetReadDeadline(time.Now().Add(time.Duration(WebSocketReadDeadline) * time.Second))
	for {
		_, bytes, err := s.conn.ReadMessage()

		if err != nil {
			s.toServiceChan <- s.id
			s.conn.Close()
			return
		}
		err = s.rpcStream.Send(&pb.Frame{Payload: bytes})
		if err != nil {
			fmt.Println("err on stream  websocket")
			s.rpcStream.CloseSend()
			s.toServiceChan <- s.id
			return
		}
	}
}
func (a AgentService) WebSocketServer(w http.ResponseWriter, r *http.Request) {
	var upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true },
	}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("err on create websocket")
		return
	}
	stream, err := a.RPCClient.Stream(context.Background())
	if err != nil {
		fmt.Println("err on create stream")
		return
	}

	id := sid(uuid.NewV4().String())
	a.Sessions[id] = NewSession(id, conn, stream, a.SessDieChan)
	go a.Sessions[id].ForwardToClient()
	go a.Sessions[id].ForwardToServer()
	fmt.Println("start new session")
}

func NewSession(id sid, conn *websocket.Conn, rpcStream pb.GMService_StreamClient, toServiceChan chan sid) *Session {
	s := &Session{
		id:            id,
		conn:          conn,
		rpcStream:     rpcStream,
		toServiceChan: toServiceChan,
	}
	return s
}
func (a AgentService) goDaemon() {
	for {
		select {
		case id := <-a.SessDieChan:
			a.closeSession(id)
		}
	}
}
func (a AgentService) closeSession(id sid) {
	a.Sessions[id].rpcStream.CloseSend()
	a.Sessions[id].conn.Close()
	delete(a.Sessions, id)
}

var (
	webSocketAddr = flag.String("websocket.addr", "192.168.1.51:7788", "game agent webSocket address")
	gameServer    = flag.String("grpc.gameserver", "192.168.1.51:7799", "game server gRPC server address")
)

func main() {
	flag.Parse()
	// gRPC transport.
	conn, err := grpc.Dial(*gameServer, grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	client := pb.NewGMServiceClient(conn)

	// Business domain.
	agentService := NewAgentService(client)

	m := http.NewServeMux()
	m.HandleFunc("/", agentService.WebSocketServer)
	http.ListenAndServe(*webSocketAddr, m)
	select {}
}
