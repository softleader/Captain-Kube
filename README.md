# Captain Kube

> A command line tool for Kubernetes hepls SoftLeader client deploying SoftLeader products

## Install

```shell
$ docker run -it --rm -v $PATH:/data softleader/captain-kube GOOS=$GOOS GOARCH=$GOARCH
```

- `$PATH` - 安裝的目錄, 如 `/usr/local/bin`
- `$GOOS` - 預設為 `linux`, 若輸入 `macos` 則自動轉換為 `darwin`
- `$GOARCH` - 預設為 `amd64`

> `GOOS` 及 `$GOARCH` 可參考 [environment variables](https://golang.org/doc/install/source#environment) 變更

### Example: 

```shell
$ docker run -it --rm -v /usr/local/bin:/data softleader/captain-kube GOOS=macos
```