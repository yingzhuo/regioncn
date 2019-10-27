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
	@echo " build          | 构建以上三者"
	@echo " clean          | 清理构建产物"
	@echo " release        | 发布"
	@echo " github         | 将代码推送到Github"
	@echo "------------------------------------------"

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
	docker image build -t registry.cn-shanghai.aliyuncs.com/yingzhor/$(NAME):$(VERSION) --build-arg VERSION=$(VERSION) $(CURDIR)/_bin
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
