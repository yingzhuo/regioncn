TIMESTAMP             	:= $(shell /bin/date "+%F %T")
NAME					:= regioncn
VERSION					:= 1.0.0

usage:
	@echo "------------------------------------------"
	@echo " 目标           | 功能"
	@echo "------------------------------------------"
	@echo " usage          | 显示本菜单"
	@echo " fmt            | 格式化代码"
	@echo " protoc         | 编译protobuf文件"
	@echo " build-linux    | 构建 (linux-amd64)"
	@echo " build-darwin   | 构建 (darwin-amd64)"
	@echo " build-windows  | 构建 (windows-amd64)"
	@echo " build-all      | 构建以上三者"
	@echo " clean          | 清理构建产物"
	@echo " release        | 发布"
	@echo " github         | 将代码推送到Github"
	@echo "------------------------------------------"

clean:
	rm -rf $(CURDIR)/bin/regioncn-*
	docker image prune -f

fmt:
	go fmt $(CURDIR)/...

protoc:
	protoc -I=$(CURDIR)/protobuf/ --go_out=$(CURDIR)/protobuf/ $(CURDIR)/protobuf/regioncn.proto

build-linux: protoc
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $(CURDIR)/bin/$(NAME)-linux-amd64-$(VERSION)

build-darwin: protoc
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o $(CURDIR)/bin/$(NAME)-darwin-amd64-$(VERSION)

build-windows: protoc
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o $(CURDIR)/bin/$(NAME)-windows-amd64-$(VERSION).exe

build-all: build-linux build-darwin build-windows

release: build-linux
	docker image build -t quay.io/yingzhuo/$(NAME):$(VERSION) --build-arg VERSION=$(VERSION) $(CURDIR)/bin
	docker login --username=yingzhuo --password="${QUAY_PASSWORD}" quay.io &> /dev/null
	docker image push quay.io/yingzhuo/$(NAME):$(VERSION)
	docker image tag  quay.io/yingzhuo/$(NAME):$(VERSION) quay.io/yingzhuo/$(NAME):latest
	docker image push quay.io/yingzhuo/$(NAME):latest
	docker logout quay.io &> /dev/null

github: clean fmt protoc
	git add .
	git commit -m "$(TIMESTAMP)"
	git push

.PHONY: usage clean fmt protoc build-linux build-darwin build-linux build-all release github
