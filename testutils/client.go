package testutils

import (
	"errors"
	"github.com/datumchi/protocol-tests/generated/protocol"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func CreateIdentityServicesClient(identityServiceUrl string) (protocol.IdentityServicesClient, error) {

	// Establish connection
	var opts []grpc.DialOption
	tlsCredentials := credentials.NewTLS(nil)
	opts = append(opts, grpc.WithTransportCredentials(tlsCredentials))
	conn, err := grpc.Dial(identityServiceUrl, opts...)
	if err != nil {
		return nil, errors.New("Client connection error:  " + err.Error())
	}

	client := protocol.NewIdentityServicesClient(conn)

	return client, nil
}
