all:
	if [ ! -d bin ]; then mkdir bin; fi
	go build -o bin/game -ldflags -s main.go

clean:
	if [ -d bin ]; then rm -rf bin; fi
