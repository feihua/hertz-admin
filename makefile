# 定义变量
ifndef GOPATH
	GOPATH := $(shell go env GOPATH)
endif

GOBIN=$(GOPATH)/bin
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOGET=$(GOCMD) mod tidy

MYSQL_INFO=root:r-wz9wop62956dh5k9ed@tcp(110.41.179.89:30395)/gozero

all: deps build ## 默认的构建目标

clean: ## 清理目标
	$(GOCLEAN)
	rm -rf target

deps: ## 安装依赖目标
	@export GOPROXY=https://goproxy.cn,direct
	$(GOGET) -v

build: ## 构建目标
	$(GOBUILD) -o target/hertz-admin

start: build stop ## 运行目标
	nohup ./target/hertz-admin > /dev/null 2>&1 &

stop: ## 停止目标
	-pkill -f hertz-admin


restart: stop start ## 重启项目

.DEFAULT_GOAL := all ## 默认构建目标是

HZ=$(GOBIN)/hz

gen:	## 生成代码
	$(HZ) update -I idl -idl idl/sys/sys.proto



image: ## 构建docker镜像
	docker build -t hertz-admin:0.0.1 -f Dockerfile .

run: image ## 启动docker容器
	docker run -itd --net=host --name=hertz-admin hertz-admin:0.0.1; \


kubectl: image ## 部署k8s容器
	kubectl apply -f script/hertz-admin.yaml; \

help: ## show help message
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m\033[0m\n"} /^[$$()% 0-9a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)
