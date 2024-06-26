
SRC	=	src/main.go \
		src/banner.go \
		src/flag.go \
		src/file.go \

NAME	=	discoDir

all: build

hello:
	echo "Hello"

build:
	go build -o $(NAME) $(SRC)

run:
	go run $(SRC)

compile:
	echo "Compiling for every OS and Platform"
	GOOS=linux GOARCH=arm go build -o bin/main-linux-arm main.go
	GOOS=linux GOARCH=arm64 go build -o bin/main-linux-arm64 main.go
	GOOS=freebsd GOARCH=386 go build -o bin/main-freebsd-386 main.go

clean:	rm $(NAME)

.PHONY:	hello build run compile clean
