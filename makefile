all: test

update:
	go get -u ./...
	go mod tidy
	go mod vendor

install-tools:
	go install github.com/google/go-licenses@latest
	go install github.com/olekukonko/tablewriter/csv2table@latest

generate-licenses: install-tools
	go-licenses csv github.com/elgohr/semv | csv2table -p=true -h=false > cmd/licenses.txt

test:
	go test -race ./...

secure:
	docker run --rm -e GO111MODULE=on -w /go/scan/ -v "$(shell pwd):/go/scan" securego/gosec:latest -exclude-dir=systemtest /go/scan/...

install-semv:
	go install github.com/elgohr/semv@latest

newPatchVersion = $(shell git describe --tags --abbrev=0 | semv increment --patch)
new-patch-release: install-semv
	echo $(newPatchVersion) > cmd/version.txt ; git add . ; git commit -m ":tada: Releasing $(newPatchVersion)" ; git push origin main ; git tag $(newPatchVersion); git push origin $(newPatchVersion)

newMinorVersion = $(shell git describe --tags --abbrev=0 | semv increment --minor)
new-minor-release: install-semv
	echo $(newMinorVersion) > cmd/version.txt ; git add . ; git commit -m ":tada: Releasing $(newMinorVersion)" ; git push origin main ; git tag $(newMinorVersion); git push origin $(newMinorVersion)

newMajorVersion = $(shell git describe --tags --abbrev=0 | semv increment --major)
new-major-release: install-semv
	echo $(newMajorVersion) > cmd/version.txt ; git add . ; git commit -m ":tada: Releasing $(newMajorVersion)" ; git push origin main ; git tag $(newMajorVersion); git push origin $(newMajorVersion)
