package console

import (
	"context"
	"errors"
	"net"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zeebo/errs"
	"golang.org/x/sync/errgroup"

	"awesomeProject/console/controllers"
	"awesomeProject/users"
)

// Config contains server required data.
type Config struct {
	Address string
}

// Server console web server implementation.
type Server struct {
	server   http.Server
	listener net.Listener
	config   Config
	users    *users.Service
}

// New console server constructor.
func New(listener net.Listener, config Config, users *users.Service) *Server {
	server := &Server{listener: listener, config: config}

	router := mux.NewRouter()
	usersRouter := router.PathPrefix("/users").Subrouter()
	usersController := controllers.NewUsers(users)

	usersRouter.HandleFunc("/create", usersController.Create).Methods(http.MethodPost)
	usersRouter.HandleFunc("/get", usersController.Get).Methods(http.MethodGet)
	usersRouter.HandleFunc("/delete", usersController.Delete).Methods(http.MethodDelete)
	usersRouter.HandleFunc("/update", usersController.Update).Methods(http.MethodPut)

	server.server = http.Server{
		Handler: router,
	}

	return server
}

// Run starts the server that host webapp and api endpoint.
func (server *Server) Run(ctx context.Context) (err error) {
	ctx, cancel := context.WithCancel(ctx)
	var group errgroup.Group
	group.Go(func() error {
		<-ctx.Done()
		return server.server.Shutdown(context.Background())
	})
	group.Go(func() error {
		defer cancel()
		err := server.server.Serve(server.listener)
		isCancelled := errs.IsFunc(err, func(err error) bool { return errors.Is(err, context.Canceled) })
		if isCancelled || errors.Is(err, http.ErrServerClosed) {
			err = nil
		}
		return err
	})

	return group.Wait()
}

// Close closes server and underlying listener.
func (server *Server) Close() error {
	return server.server.Close()
}
