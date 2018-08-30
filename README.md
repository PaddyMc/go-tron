# go-tron

<h1>
	Running the docker file
</h1>
<code>docker build .</code>
<code>docker run -p 50051:50051 -it "--imagename--" /bin/bash -c "cd build/libs;java -jar java-tron.jar"</code>
This will start the tron node, and allow go-tron to calculate circulating supply

<br/>protoc --go_out=plugins=grpc:. api/api.proto
<!-- <br/>go get github.com/tronprotocol/grpc-gateway -->
<br/>go get google.golang.org/grpc
<br/>go get github.com/shengdoushi/base58

<br/>go run start.go -grpcAddress localhost:50051