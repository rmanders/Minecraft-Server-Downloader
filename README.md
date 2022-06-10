# Minecraft-Server-Downloader
A little tool written in go for fetching the most recent Minecraft Java Edition server jar

![main build](https://github.com/rmanders/minecraft-server-downloader/actions/workflows/go.yml/badge.svg?branch=main)

## Usage
Simply run the ```mc-server-downloader``` file and it will fetch the most recent Minecraft Java Edition server release version to the current directory. 

## Building

```
go build ./cmd/mc-server-downloader
```

### OS Specific builds  
<br />

#### Windows
```
GOOS=windows GOARCH=amd64 go build ./cmd/mc-server-downloader
```

#### Mac M1 
```
GOOS=darwin GOARCH=arm64 go build ./cmd/mc-server-downloader
```

#### Linux
```
GOOS=linux GOARCH=amd64 go build ./cmd/mc-server-downloader
```