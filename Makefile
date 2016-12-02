default: yacc
	@go run main.go

test:
	@go test ./... -cover -v

yacc:
	@go tool yacc -o parser/parser.go -p Tdoc parser/tdoc.y

cover:
	@go tool cover -html=coverage.out
