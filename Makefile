BINARY=mymaths
BIN_DIR=bin
build:
		go build -o ${BIN_DIR}/${BINARY} main.go
run:
		go run main.go
run_binary:
		if [ -f "./${BIN_DIR}/${BINARY}" ]; then chmod a+x ./${BIN_DIR}/${BINARY}; ./${BIN_DIR}/${BINARY}; fi
test:
		if [ -d "./Maths" ]; then cd ./Maths ; go test -v ; fi
		if [ -d "./Accounts" ]; then cd ./Accounts ; go test -v ; fi
rm:
		if [ -d "./${BIN_DIR}" ]; then rm -Rf ./${BIN_DIR} ; fi
format:
		if [ -d "./Maths" ]; then cd ./Maths ; go fmt *.go ; fi
		if [ -d "./Accounts" ]; then cd ./Accounts ; go fmt *.go ; fi
