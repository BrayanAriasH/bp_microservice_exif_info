
rm main.zip > /dev/null
export GOOS=linux
export CGO_ENABLED=0
export GOARCH=amd64
go build ./src/main.go
zip main.zip main
rm main