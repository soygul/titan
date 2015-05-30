package devastator

import (
	"fmt"
	"log"
	"strconv"

	"github.com/nbusy/devastator/neptulon"
	"github.com/nbusy/devastator/neptulon/jsonrpc"
)

// Token is an encrypted identifier for connecting devices.
type Token struct {
	ID uint32
	IV []byte
}

func authMiddleware(conn *neptulon.Conn, session *neptulon.Session, msg *jsonrpc.Message) {
	peerCerts := conn.ConnectionState().PeerCertificates

	// client certificate authorization: certificate is verified by the TLS listener instance so we trust it
	if len(peerCerts) > 0 {
		idstr := peerCerts[0].Subject.CommonName
		uid64, err := strconv.ParseUint(idstr, 10, 32)
		if err != nil {
			session.Error = fmt.Errorf("Cannot parse client message or method mismatched: %v", err)
			return
		}
		userID := uint32(uid64)
		log.Printf("Client connected with client certificate subject: %+v", peerCerts[0].Subject)
		session.UserID = userID
	}
}
