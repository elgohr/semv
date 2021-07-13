all: test

install-tools:
	go install github.com/google/go-licenses
	go install github.com/olekukonko/tablewriter/csv2table

tidy:
	go mod tidy

generate-licenses: install-tools tidy
	go-licenses csv github.com/elgohr/semv | csv2table -p=true > cmd/licenses.txt

test:
	go test -race ./...

secure:
	docker run --rm -e GO111MODULE=on -w /go/scan/ -v "$(shell pwd):/go/scan" securego/gosec:latest -exclude-dir=systemtest /go/scan/...
