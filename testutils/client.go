package testutils

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"github.com/datumchi/protocol-tests/generated/protocol"
	"github.com/datumchi/protocol-tests/testutils/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)


func CreateIdentityServicesClient(identityServiceUrl string) (protocol.IdentityServicesClient, error) {

	// CA roots
	roots := x509.NewCertPool()
	if !roots.AppendCertsFromPEM([]byte(TLS_CA)) {
		return nil, errors.New("Cannot append CA Root cert")
	}

	// Establish connection
	var opts []grpc.DialOption
	tlsCredentials := credentials.NewTLS(&tls.Config{RootCAs: roots})
	opts = append(opts, grpc.WithTransportCredentials(tlsCredentials))
	conn, err := grpc.Dial(identityServiceUrl, opts...)
	if err != nil {
		return nil, errors.New("Client connection error:  " + err.Error())
	}

	logger.Infof("Connecting:  %s", identityServiceUrl)
	client := protocol.NewIdentityServicesClient(conn)

	return client, nil
}
