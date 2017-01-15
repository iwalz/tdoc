default: yacc
	@go run main.go

test:
	@for pkg in `go list ./...`; do \
		sub_pkg=$$(echo "$$pkg" | cut -d '/' -f 4) ; \
		if [ $$sub_pkg ]; then \
			go test -v $$pkg -coverprofile="$$sub_pkg".coverprofile ; \
		fi \
	done

yacc:
	@go tool yacc -o parser/parser.go -p Tdoc parser/tdoc.y

cover:
	@go tool cover -html=gover.coverprofile
