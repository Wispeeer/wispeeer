# wispeeer
我的静态博客生成器

## 依赖 (Require)
### 编译依赖 (Build Require)
> GNU Make 4.3 (make --version)
>
> go version go1.16 (go version)

## Build & Run
```
make build
./bin/wispeeer_$(go env GOOS)_$(go env GOARCH) -h
```
more build option
```bash
make help
```

## Usage
```bash
wispeeer -h
```
