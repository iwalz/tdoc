default: yacc
	@go run main.go

test:
	@echo "Testing lexer:"
	@go test ./lexer -cover -v

yacc:
	@go tool yacc -o lexer/parser.go -p Tdoc lexer/tdoc.y

cover:
	@go tool cover -html=coverage.out
