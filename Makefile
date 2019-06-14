.PHONY: 执行go的顺序1：go build -o 生成命令目录 main文件
.PHONY: 执行go的顺序2：./bin/执行命令
.PHONY: 执行go的顺序3：测试：直接 go run main文件执行


.PHONY: default
default: chan

.PHONY: chan
chan:
*****go build -o ./bin/chan-bin ./pkg/main/chan.go

