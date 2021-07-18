go test ./... -coverprofile=coverage.out
go tool cover -html=coverage.out
go test -coverprofile fmtcoverage.html fmt