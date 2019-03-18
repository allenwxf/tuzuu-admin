git clone https://github.com/googleapis/google-cloud-go $GOPATH/src/cloud.google.com

go get -u github.com/gin-gonic/gin

go get github.com/gin-contrib/cors

go get github.com/go-sql-driver/mysql

go get github.com/mgutz/ansi

mkdir -p src/golang.org/x/

git clone https://github.com/grpc/grpc-go.git $GOPATH/src/google.golang.org/grpc
git clone https://github.com/golang/net.git $GOPATH/src/golang.org/x/net
git clone https://github.com/golang/text.git $GOPATH/src/golang.org/x/text
go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
git clone https://github.com/google/go-genproto.git $GOPATH/src/google.golang.org/genproto

cd $GOPATH/src/golang.org/x/; git clone https://github.com/golang/crypto.git; git clone https://github.com/golang/sys.git

go get github.com/NebulousLabs/fastrand

go get github.com/sirupsen/logrus