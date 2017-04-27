
export GOPATH=`pwd`
go get github.com/jfelten/hook_manager
CGO_ENABLED=0 GOOS=linux go build -a -tags netgo -ldflags '-w' github.com/jfelten/hook_manager/
docker build . -t jfelten/hook_manager
docker push jfelten/hook_manager
rm -rf src pkg bin