





generate-protocol:
	mkdir -p ./generated/protocol
	protoc -I=../protocol/protobuf --go_out=plugins=grpc:./generated/protocol ../protocol/protobuf/*.proto


generate-tls:

	# Create Dirs
	mkdir -p ./security/ca
	mkdir -p ./security/alphaidentity.datumchi.com

	# Generate Root Key
	openssl genrsa -out ./security/ca/ca.key 4096

	# Self Signed Root Cert
	openssl req -new -x509 -key ./security/ca/ca.key -sha256 -subj "/C=US/ST=WA/O=DatumChi CA" -days 365 -out ./security/ca/ca.crt

	# Creating Alpha Identity Key
	openssl genrsa -out ./security/alphaidentity.datumchi.com/alphaidentity.datumchi.com.key 4096

	# Create CSR for Alpha Identity
	openssl req -new -key ./security/alphaidentity.datumchi.com/alphaidentity.datumchi.com.key -out ./security/alphaidentity.datumchi.com/alphaidentity.datumchi.com.csr -config ./security/alphaidentity.datumchi.com/certificate.conf

	# Generate Certificate
	openssl x509 -req -in ./security/alphaidentity.datumchi.com/alphaidentity.datumchi.com.csr -CA ./security/ca/ca.crt -CAkey ./security/ca/ca.key -CAcreateserial -out ./security/alphaidentity.datumchi.com/alphaidentity.datumchi.com.pem -days 365 -sha256 -extfile ./security/alphaidentity.datumchi.com/certificate.conf -extensions req_ext
