package server

import (
	"sync"

	"github.com/osamikoyo/hrm-notify/internal/reciewer"
	"github.com/osamikoyo/hrm-notify/pkg/config"
	"github.com/osamikoyo/hrm-notify/pkg/loger"
)

type Server struct{
	Reciewer *reciewer.Reciewer
	Logger loger.Logger
}

func New() (*Server, error) {
	cfg, err := config.LoadConfig()
	if err != nil{
		return nil, err
	}

	rec, err := reciewer.Init(&cfg)
	if err != nil{
		return nil, err
	}

	return &Server{
		Reciewer: rec,
		Logger: loger.New(),
	}, nil
}

func (s *Server) Run() {
	s.Logger.Info().Msg("Starting service")

	var wg *sync.WaitGroup
	wg.Add(1)


	go func() {
		for {
			s.Reciewer.WaitMessages()
		}
	}()

	wg.Wait()
}