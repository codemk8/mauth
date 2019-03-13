ifndef $(tag)
	tag=latest
endif

build: cmd/*/*.go
	go build -o bin/mauth ./cmd/mauth.go
	#GOOS=linux go build -o bin/mauth ./cmd/mauth

test: pkg/*/*.go
	go test github.com/codemk8/mauth/pkg

clean:
	-rm -rf bin/*
