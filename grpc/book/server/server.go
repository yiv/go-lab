package main

import (
	"avg_service/user/pkg/user"
	"avg_service/x/mysql"
	"flag"
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net"
	"net/http"
	"net/http/pprof"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pb "github.com/yiv/go-lab/grpc/book/editorpb"
)

var (
	port      = flag.Int("port", 11000, "The server port")
	debugAddr = flag.String("debug.addr", ":11100", "Debug and metrics listen address")
	runTime   float64
)

type UserServer struct {
	Repo   *mysql.MysqlRepo
	Logger log.Logger
}

func NewUserServer(logger log.Logger) UserServer {
	us := UserServer{Logger: logger}
	mysqlOps := mysql.MysqlOptions{DBAddr: "10.72.12.19", DBPort: "3306", DBUser: "root", DBPwd: "root"}
	us.Repo, _ = mysql.NewMysqlRepo(mysqlOps, logger)
	return us
}
func (u UserServer) Characters(context context.Context, req *pb.CharactersReq) (res *pb.CharactersRes, err error) {
	begin := time.Now()
	if resp, err := u.Repo.GetCharacters(req.BookId); err != nil {
		level.Error(u.Logger).Log("err", err.Error())
		return nil, err
	} else {
		//level.Info(u.Logger).Log("server-mysql", time.Since(begin).Seconds())
		runTime += time.Since(begin).Seconds()
		var characs []*pb.AppEditorCharacter
		for _, v := range resp {
			c := &pb.AppEditorCharacter{
				Id:            v.Id,
				BookId:        v.BookId,
				NickName:      v.NickName,
				CharacterType: v.CharacterType,
				Align:         v.Align,
				Icon:          v.Icon,
				Created:       v.Created,
				Updated:       v.Updated,
			}
			characs = append(characs, c)
		}
		res = &pb.CharactersRes{Data: characs}
	}
	return
}

func (u UserServer) EditProfile(context context.Context, req *pb.EditProfileReq) (res *pb.EditProfileRes, err error) {
	begin := time.Now()
	profile := user.AppUserInfo{
		ID:        req.Profile.ID,
		UserID:    req.Profile.UserID,
		NickName:  req.Profile.NickName,
		Desc:      req.Profile.Desc,
		Email:     req.Profile.Email,
		Gender:    req.Profile.Gender,
		Age:       req.Profile.Age,
		Avatar:    req.Profile.Avatar,
		BgImg:     req.Profile.BgImg,
		Signature: req.Profile.Signature,
		Birthday:  req.Profile.Birthday,
		Nation:    req.Profile.Nation,
		Province:  req.Profile.Province,
		City:      req.Profile.City,
		County:    req.Profile.County,
		Location:  req.Profile.Location,
		RegIP:     req.Profile.RegIP,
		Created:   req.Profile.Created,
		Updated:   req.Profile.Updated,
	}

	if err := u.Repo.EditProfile(req.Uid, req.Fields, profile); err != nil {
		level.Error(u.Logger).Log("err", err.Error())
		return nil, err
	} else {
		//level.Info(u.Logger).Log("server-mysql", time.Since(begin).Seconds())
		runTime += time.Since(begin).Seconds()
		res = &pb.EditProfileRes{Err: ""}
	}
	return
}

func main() {
	flag.Parse()

	var logger log.Logger
	{
		logLevel := level.AllowInfo()
		logLevel = level.AllowDebug()
		logger = log.NewLogfmtLogger(os.Stdout)
		logger = level.NewFilter(logger, logLevel)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		level.Error(logger).Log("err", err.Error())
		os.Exit(0)
	}

	errc := make(chan error)
	go func() {
		grpcServer := grpc.NewServer()
		grpcHandler := NewUserServer(logger)
		pb.RegisterEditorServer(grpcServer, grpcHandler)
		errc <- grpcServer.Serve(lis)
	}()

	go func() {
		logger := log.With(logger, "transport", "debug")
		m := http.NewServeMux()
		m.Handle("/debug/pprof/", http.HandlerFunc(pprof.Index))
		m.Handle("/debug/pprof/cmdline", http.HandlerFunc(pprof.Cmdline))
		m.Handle("/debug/pprof/profile", http.HandlerFunc(pprof.Profile))
		m.Handle("/debug/pprof/symbol", http.HandlerFunc(pprof.Symbol))
		m.Handle("/debug/pprof/trace", http.HandlerFunc(pprof.Trace))
		m.Handle("/metrics", promhttp.Handler())

		logger.Log("addr", *debugAddr)
		errc <- http.ListenAndServe(*debugAddr, m)
	}()

	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errc <- fmt.Errorf("%s", <-c)
	}()

	logger.Log("exit", <-errc)
	level.Info(logger).Log("server-mysql", runTime)
}
