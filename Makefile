.PHONY: all build run gotool clean help

APP=syslog


BINARYWIN="./Releases/${APP}-win"
BINARYMAC="./Releases/${APP}-mac"
BINARYLINUX="./Releases/${APP}-linux"


all: gotool build
build:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ${BINARYWIN}
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o ${BINARYMAC}
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${BINARYLINUX}
run:
	@go run ./
gotool:
	go fmt ./
	go vet ./
clean:
	@if [ -f ${BINARYWIN} ] ; then rm ${BINARYWIN} ; fi
	@if [ -f ${BINARYMAC} ] ; then rm ${BINARYMAC} ; fi
	@if [ -f ${BINARYLINUX} ] ; then rm ${BINARYLINUX} ; fi
help:
	@echo "make - 格式化 Go 代码， 并编译生成二进制文件"
	@echo "make build - 编译 Go 代码，生成二进制文件"
	@echo "make run - 直接运行 Go 代码"
	@echo "make clean - 移除二进制文件和 vim swap files"
	@echo "make gotool - 运行 Go 工具 'fmt' and 'vet'"