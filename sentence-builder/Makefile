build:
	dep ensure
	env GOOS=linux go build -ldflags="-s -w" -o bin/begin cmd/begin/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/stepfunction_phrase cmd/stepfunctions/phrase/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/stepfunction_subject cmd/stepfunctions/subject/main.go
	