README.md: main.go releases.json README-template.md
	go run main.go > README.md

lint-go:
	golangci-lint run

lint-json:
	jq . releases.json > /dev/null

lint: lint-go lint-json