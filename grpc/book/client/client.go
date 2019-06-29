package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/gorilla/mux"
	pb "github.com/yiv/go-lab/grpc/book/editorpb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	//"os"
	//"os/signal"
	//"syscall"
	//"time"
)

var (
	serverAddr = flag.String("server_addr", "127.0.0.1:11000", "The server address in the format of host:port")
	httpAddr   = flag.String("http.addr", ":11001", "Address for HTTP (JSON) server")
	runTime    float64
)

func main() {
	flag.Parse()

	// Logging domain.
	var logger log.Logger
	{
		logLevel := level.AllowInfo()
		logLevel = level.AllowDebug()
		logger = log.NewLogfmtLogger(os.Stdout)
		logger = level.NewFilter(logger, logLevel)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	conn, err := grpc.Dial(*serverAddr, grpc.WithInsecure())
	if err != nil {
		os.Exit(1)
	}
	defer conn.Close()
	var clients []pb.EditorClient
	for i := 0; i < 10; i++ {
		client := pb.NewEditorClient(conn)
		clients = append(clients, client)
	}
	rand.Seed(time.Now().UnixNano())
	// server routes.
	r := mux.NewRouter()
	{
		index := rand.Int31n(10)
		client := clients[index]
		r.PathPrefix("/app/editor/characters").Handler(http.StripPrefix("/app/editor/characters", HandleCharacters(logger, client)))
		r.PathPrefix("/app/user/edit-profile").Handler(http.StripPrefix("/app/user/edit-profile", HandleEditProfile(logger, client)))
	}
	errc := make(chan error)
	go func() {
		logger.Log("transport", "HTTP", "addr", *httpAddr)
		errc <- http.ListenAndServe(*httpAddr, r)
	}()
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errc <- fmt.Errorf("%s", <-c)
	}()
	logger.Log("exit", <-errc)
	level.Info(logger).Log("client-server", runTime)
}

func HandleCharacters(logger log.Logger, client pb.EditorClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		req := &pb.CharactersReq{BookId: 1500086}
		jsRes := map[string]interface{}{"msg": "操作成功", "code": 200}
		begin := time.Now()
		if res, err := client.Characters(context.Background(), req); err != nil {
			w.WriteHeader(500)
		} else {
			w.WriteHeader(200)
			jsRes["data"] = res
		}
		runTime += time.Since(begin).Seconds()
		rb, _ := json.Marshal(jsRes)
		w.Write(rb)
	})
}

func HandleEditProfile(logger log.Logger, client pb.EditorClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		desc := fmt.Sprintf("%v", time.Now().UnixNano())
		req := &pb.EditProfileReq{Uid: 221, Fields: []string{"desc"}, Profile: &pb.AppUserInfo{Desc: desc}}
		jsRes := map[string]interface{}{"msg": "操作成功", "code": 200}
		begin := time.Now()
		if res, err := client.EditProfile(context.Background(), req); err != nil {
			w.WriteHeader(500)
		} else {
			w.WriteHeader(200)
			jsRes["data"] = res
		}
		runTime += time.Since(begin).Seconds()
		rb, _ := json.Marshal(jsRes)
		w.Write(rb)
	})
}
