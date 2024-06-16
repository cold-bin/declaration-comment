.PHONY: test install
all: test install
install:
	go install ./cmd/declarationcomment
test:
	go test -v ./pkg/analyzer/