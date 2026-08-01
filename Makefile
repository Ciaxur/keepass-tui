BIN := bin/combine_dirpaths
SRC := combine_dirpaths/main.go

.PHONY: all test clean

all: $(BIN)

$(BIN): $(SRC)
	mkdir -p bin
	cd combine_dirpaths && go build -o ../$(BIN) .

test:
	cd combine_dirpaths && go test ./...

clean:
	rm -rf bin
