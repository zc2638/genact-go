build:
	@CGO_ENABLED=0 go build  -ldflags="-s -w" -installsuffix cgo -o /usr/local/bin/genact github.com/zc2638/genact-go/cmd/genact

docker:
	@docker build -t zc2638/genact-go -f build/Dockerfile .
