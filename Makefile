TIMESTAMP  	:= $(shell /bin/date "+%F %T")
VERSION		:= 1.0.0
NAME		:= regioncn
LDFLAGS		:= -s -w \
			   -X 'main.BuildVersion=$(VERSION)' \
			   -X 'main.BuildGitBranch=$(shell git describe --all)' \
			   -X 'main.BuildGitRev=$(shell git rev-list --count HEAD)' \
			   -X 'main.BuildGitCommit=$(shell git rev-parse HEAD)' \
			   -X 'main.BuildDate=$(shell /bin/date "+%F %T")'

usage:
	@echo "no default target"; false

clean:
	rm -rf $(CURDIR)/_bin/regioncn-*
	docker image rm registry.cn-shanghai.aliyuncs.com/yingzhor/$(NAME):$(VERSION) &> /dev/null || true
	docker image rm registry.cn-shanghai.aliyuncs.com/yingzhor/$(NAME):latest &> /dev/null || true
	docker image prune -f

fmt:
	go fmt $(CURDIR)/...

protoc:
	protoc -I=$(CURDIR)/protobuf/ --go_out=$(CURDIR)/protobuf/ $(CURDIR)/protobuf/regioncn.proto

build: protoc
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $(CURDIR)/_bin/$(NAME)-linux-amd64-$(VERSION)

release: build
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags "$(LDFLAGS)" -o $(CURDIR)/_bin/$(NAME)-linux-amd64-$(VERSION)
	docker image tag      registry.cn-shanghai.aliyuncs.com/yingzhor/$(NAME):$(VERSION) registry.cn-shanghai.aliyuncs.com/yingzhor/$(NAME):latest
	docker login --username=yingzhor@gmail.com --password="${ALIYUN_PASSWORD}" registry.cn-shanghai.aliyuncs.com
	docker image push registry.cn-shanghai.aliyuncs.com/yingzhor/$(NAME):$(VERSION)
	docker image push registry.cn-shanghai.aliyuncs.com/yingzhor/$(NAME):latest
	docker logout registry.cn-shanghai.aliyuncs.com &> /dev/null

github: clean fmt protoc
	git add .
	git commit -m "$(TIMESTAMP)"
	git push

.PHONY: usage clean fmt protoc build-linux build-darwin build-linux build-all release github
