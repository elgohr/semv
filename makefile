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

install-semv:
	go install github.com/elgohr/semv@latest

new-patch-release: install-semv
	newVersion=$(semv increment --patch $(git describe --tags --abbrev=0)) ; echo ${newVersion} > cmd/version.txt ; git add . ; git commit -m ":tada: Releasing version ${newVersion}" ; git push origin main ; git tag ${newVersion}; git push origin ${newVersion}

new-minor-release: install-semv
	newVersion=$(semv increment --minor $(git describe --tags --abbrev=0)) ; echo ${newVersion} > cmd/version.txt ; git add . ; git commit -m ":tada: Releasing version ${newVersion}" ; git push origin main ; git tag ${newVersion}; git push origin ${newVersion}

new-major-release: install-semv
	newVersion=$(semv increment --major $(git describe --tags --abbrev=0)) ; echo ${newVersion} > cmd/version.txt ; git add . ; git commit -m ":tada: Releasing version ${newVersion}" ; git push origin main ; git tag ${newVersion}; git push origin ${newVersion}
