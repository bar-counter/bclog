
# for what

- this project used to gin logs
- pkg name `github.com/bar-counter/bclog` to target

## dependInfo

| lib | url | version |
|:-----|:-----|:-----|
| go sdk | https://golang.org/ | 1.12+ |
| gin | https://github.com/gin-gonic/gin | 1.4.0 |

# demo

```bash
make init
# ensure right then
make dev
# run as docker contant
$ make dockerRun
# stop or remove docker
$ make dockerStop
$ make dockerRemove
```

# just-use lib

- zap https://github.com/uber-go/zap
- lumberjack https://github.com/natefinch/lumberjack

# benchmarking

- logrus https://github.com/sirupsen/logrus

# import plugin

```bash
# above go 1.12
GOPROXY=https://mirrors.aliyun.com/goproxy/ GO111MODULE=on go mod edit -require=github.com/json-iterator/go@v1.1.7
make dep
```

```go
import "github.com/json-iterator/go"

var json = jsoniter.ConfigCompatibleWithStandardLibrary
```
