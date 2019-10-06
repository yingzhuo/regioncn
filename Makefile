TIMESTAMP             	:= $(shell /bin/date "+%F %T")
NAME					:= regioncn
VERSION					:= 1.0.0

clean:
	rm -rf $(CURDIR)/bin/regioncn-*
	docker image prune -f

fmt:
	go fmt $(CURDIR)/...

build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $(CURDIR)/bin/$(NAME)-linux-amd64-$(VERSION)

build-darwin:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o $(CURDIR)/bin/$(NAME)-darwin-amd64-$(VERSION)

build-windows:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o $(CURDIR)/bin/$(NAME)-windows-amd64-$(VERSION).exe

build-all: build-linux build-darwin build-windows

release: build-linux
	docker image build -t quay.io/yingzhuo/$(NAME):$(VERSION) --build-arg VERSION=$(VERSION) $(CURDIR)/bin
	docker login --username=yingzhuo --password="${QUAY_PASSWORD}" quay.io &> /dev/null
	docker image push quay.io/yingzhuo/$(NAME):$(VERSION)
	docker image tag  quay.io/yingzhuo/$(NAME):$(VERSION) quay.io/yingzhuo/$(NAME):latest
	docker image push quay.io/yingzhuo/$(NAME):latest
	docker logout quay.io &> /dev/null

.PHONY: clean fmt build-linux build-darwin build-linux build-all release
