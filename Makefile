build:
	GOOS=windows GOARCH=amd64 go build -o bin/passowrdGenerator.exe
	GOOS=linux GOARCH=amd64 go build -o bin/passwordGenerator 
