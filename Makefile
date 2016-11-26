default:
	@go tool yacc -o main.go -p Tdoc tdoc.y
	@go run main.go lexer.go

test:
	@echo "Testing lexer:"
	@go test main.go lexer.go *_test.go
