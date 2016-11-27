default: yacc
	@go run main.go lexer.go

test:
	@echo "Testing lexer:"
	@go test -v main.go lexer.go *_test.go

yacc:
	@go tool yacc -o main.go -p Tdoc tdoc.y
