default: yacc
	@go run main.go

test:
	@go test ./... -v

yacc:
	@go tool yacc -o parser/parser.go -p Tdoc parser/tdoc.y

cover:
	@go tool cover -html=gover.coverprofile
