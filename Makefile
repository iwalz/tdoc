default: yacc
	@go run main.go

test:
	@for pkg in `go list ./...`; do \
		sub_pkg=$$(echo "$$pkg" | cut -d '/' -f 4) ; \
		if [ $$sub_pkg ]; then \
			richgo test -v $$pkg -coverprofile="$$sub_pkg".coverprofile ; \
		fi \
	done

yacc:
	@go tool yacc -o parser/parser.go -p Tdoc parser/tdoc.y

cover:
	@go tool cover -html=gover.coverprofile

deps:
	@go get github.com/stretchr/testify/assert
	@go get github.com/davecgh/go-spew/spew
	@go get github.com/dnephin/cobra/cobra
	@go get github.com/ajstarks/svgo
	@go get github.com/kyoh86/richgo
	@go get github.com/Sirupsen/logrus
	@go get github.com/spf13/afero
	@go get github.com/mattn/goveralls
	@go get golang.org/x/tools/cmd/cover
	@go get github.com/modocache/gover
