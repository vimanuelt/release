# release - Display GhostBSD release information

This is a suggested workaround for GhostBSD users to view release information. To be updated for each release.

## Dependencies
```
 sudo pkg install go
 go get github.com/urfave/cli
```

## Build
```
 git clone https://github.com/vimanuelt/release.git
 cd release
 go build release
 ```
 ## Install
 ```
 cp release /usr/local/bin
 ./release
```
