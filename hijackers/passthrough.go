package hijackers

import (
	"fmt"
	"net"
	"net/http"
)

type passThroughHijacker struct {
	dialer Dialer
}

func NewPassThroughHijacker(dialer Dialer) Hijacker {
	return &passThroughHijacker{
		dialer: dialer,
	}
}

func (h *passThroughHijacker) GetConns(url *http.Request, clientRaw net.Conn, _ Logger) (net.Conn, net.Conn, error) {
	remoteConn, err := h.dialer.Dial("tcp", url.Host)
	if err != nil {
		return nil, nil, err
	}
	_, err = clientRaw.Write([]byte(fmt.Sprintf("%s 200 OK\r\n\r\n", url.Proto)))
	return clientRaw, remoteConn, err
}
