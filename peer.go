package awesomeProject

import (
	"context"
	"errors"
	"net"

	"github.com/zeebo/errs"
	"golang.org/x/sync/errgroup"

	"awesomeProject/console"
	"awesomeProject/users"
)

// DB provides access to all databases and database related functionality.
type DB interface {
	Users() users.DB
	Close()
	CreateSchema(ctx context.Context) error
}

// Config is the global configuration for peer.
type Config struct {
	Console struct {
		Server console.Config `json:"server"`
	} `json:"console"`
}

// Peer is the representation of a peer.
type Peer struct {
	Database DB
	Config   Config

	Users struct {
		Service *users.Service
	}

	Console struct {
		Listener net.Listener
		Endpoint *console.Server
	}
}

// New is a constructor for Peer.
func New(config Config, db DB) (peer *Peer, err error) {
	peer = &Peer{
		Database: db,
		Config:   config,
	}

	// setup Users
	{
		peer.Users.Service = users.New(db.Users())
	}

	// setup Console
	{
		peer.Console.Listener, err = net.Listen("tcp", config.Console.Server.Address)
		if err != nil {
			return nil, err
		}
		peer.Console.Endpoint = console.New(peer.Console.Listener, peer.Config.Console.Server, peer.Users.Service)
	}

	return peer, nil
}

// Run runs peer until it's either closed or it errors.
func (peer *Peer) Run(ctx context.Context) error {
	group, ctx := errgroup.WithContext(ctx)

	group.Go(func() error {
		return ignoreCancel(peer.Console.Endpoint.Run(ctx))
	})

	return group.Wait()
}

// Close closes all the resources.
func (peer *Peer) Close() error {
	var errlist errs.Group

	errlist.Add(peer.Console.Endpoint.Close())

	return errlist.Err()
}

// we ignore cancellation and stopping errors since they are expected.
func ignoreCancel(err error) error {
	if errors.Is(err, context.Canceled) {
		return nil
	}

	return err
}
