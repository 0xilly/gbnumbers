Name = gbnumbers

compile: compileLinux compileWin32

compileLinux:
	CGO_ENABLED=1 CC=gcc GOOS=linux GOARCH=amd64 go build -tags static -ldflags "-s -w" -o gbnumbers cmd/main.go

compileWin32:
	CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc GOOS=windows GOARCH=amd64 go build -tags static -ldflags "-s -w" -o gbnumbers.exe cmd/main.go