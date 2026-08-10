package drivermongo

import (
	"context"
	"net"
	"time"

	"golang.org/x/crypto/ssh"
)

// Implementamos SetDeadline para que no haga nada y no devuelva error
type sshConn struct{ net.Conn }

func (c *sshConn) SetDeadline(t time.Time) error      { return nil }
func (c *sshConn) SetReadDeadline(t time.Time) error  { return nil }
func (c *sshConn) SetWriteDeadline(t time.Time) error { return nil }

// 1. Definimos un struct que guardará nuestro cliente SSH
type sshDialer struct {
	client *ssh.Client
}

// 2. Implementamos el método DialContext para cumplir con la interfaz options.ContextDialer
func (s *sshDialer) DialContext(ctx context.Context, network, addr string) (net.Conn, error) {
	conn, err := s.client.Dial(network, addr)
	if err != nil {
		return nil, err
	}
	// Devolvemos nuestra conexión envuelta que ignora los deadlines
	return &sshConn{conn}, nil
}
