# release - Display GhostBSD release information

This is a suggested workaround for GhostBSD users to view release information. To be updated for each release.

## Usage
```
  release [global options]
```

## Version
```  
  18.12
```  
## Commands
```
  help, h  Shows a list of commands
```  
## Global Options
```
  --help, -h  shows help
  --version, -v print the version
```

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
sudo cp release /usr/local/bin
 ./release
```
