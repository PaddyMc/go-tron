# go-tron

<h3>
	Running the docker file
</h3>

<code>docker build .</code><br/>
<code>docker run -p 50051:50051 -it "--imagename--" /bin/bash -c "cd build/libs;java -jar java-tron.jar"</code><br/>

This will start the tron node, and expose port 50051 on localhost

<h3>
	Dependencies
</h3>

go get github.com/golang/protobuf <br/>
go get google.golang.org/grpc <br/>
go get github.com/shengdoushi/base58 <br/>

<h3>
	Running go-tron
</h3>

go run start.go -grpcAddress localhost:50051<br/>

<h3>
	Generating protobuffers
</h3>

TODO: List protobuffer build steps<br/>
protoc --go_out=plugins=grpc:. api/api.proto<br/>
