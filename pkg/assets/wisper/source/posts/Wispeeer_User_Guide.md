------
title: Wispeeer User Guide
posted: 2021-03-29 3:20:37
tags: Wispeeer
categories: Wispeeer
------

Usage: wispeeer -[gsdhv] [-i &lt;alias&gt;] [-n &lt;title&gt;] 

If you are using it for the first time,
first execute the command "wispeeer init &lt;Blog directory&gt;" to initialize the blog.

<!-- more -->

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

![screenshot](screenshot.png)

## wispeeer init

```bash
wispeeer init ${USER}.github.io
cd ${USER}.github.io
```

# Reference
- [github.com/ka1i/wispeeer](https://github.com/ka1i/wispeeer)
- [github.com/Wispeeer/wisper](https://github.com/Wispeeer/wisper)
