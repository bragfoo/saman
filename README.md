# saman
Saman is a growth platform

```
.
├── LICENSE
├── README.md
├── backend   
│   ├── api ------------ # api server source code
│   └── master --------- # master server source code, read es and write influxdb
├── docs --------------- # documents
├── plugin ------------- # plugin for master
├── res ---------------- # resource file
├── util --------------- # util package
├── conf --------------- # default config
└── web ---------------- # front web site
```
## Start

1. Install govendor

    ```shell
    go get -u github.com/kardianos/govendor 
    ```
2. Get Source Code

    ```shell
    go get github.com/bragfoo/saman
    ```

3. Install deps

    ```shell
    cd $GOPATH/src/github.com/bragfoo/saman
    govendor sync
    ```

4. start backend

    ```shell
    go run api/main.go
    ```
5. start fron

