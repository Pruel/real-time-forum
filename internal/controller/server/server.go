package server

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/Pruel/real-time-forum/internal/controller/router"
	"github.com/Pruel/real-time-forum/pkg/cstructs"
	"github.com/Pruel/real-time-forum/pkg/serror"
)

type Server struct {
	httpServer *http.Server
	// tcpServer for example
	// udpServer for example
	// smtpServer for example
	// wsServer for example
}

func New(cfg *cstructs.Config, handler *router.Router) (*Server, error) {
	if cfg == nil {
		return nil, serror.ErrNilConfigStruct
	}

	Addr := strings.Join([]string{cfg.HTTPServer.Host, cfg.HTTPServer.Port}, ":")

	return &Server{
		httpServer: &http.Server{
			Addr:           Addr,
			IdleTimeout:    cfg.HTTPServer.IdleTimeout * time.Second,
			WriteTimeout:   cfg.HTTPServer.WriteTimeout * time.Second,
			ReadTimeout:    cfg.HTTPServer.ReadTimeout * time.Second,
			MaxHeaderBytes: cfg.HTTPServer.MaxHeaderMB << 20,
			Handler:        handler.Mux,
		},
	}, nil
}

// TODO:  run server
func (s *Server) RunServer() error {
	// if err := s.httpServer.ListenAndServe(); err != nil {
	// 	return err
	// }
	// return nil
	// То что выше, тоже правильная реализация, но сам по себе метод httpServer в стоке возвращает ошибку и можно не писать проверку и всё будет окей. Т.е код выше
	// Равнозначен коду нижу

	return s.httpServer.ListenAndServe()
}

// TODO: stop server
func (s *Server) StopServer(ctx context.Context) error {
	// s := &Server{} Ресивер в функции позволяет нам не инициализировать в теле функции структуру а уже сходу работать над ней.
	return s.httpServer.Shutdown(ctx) // Этот код так-же можно сократить до 1 строчки.
}
