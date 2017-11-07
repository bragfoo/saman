rm -rf build
cd api
echo 'building linux64'
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ../build/saman_amd64_linux
echo 'building linux86'
CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -o ../api/build/saman_386_linux
echo 'building linuxarm'
CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -o ../build/saman_arm_linux
echo 'building MacOS64'
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o ../build/saman_amd64_darwin
echo 'building MacOS86'
CGO_ENABLED=0 GOOS=darwin GOARCH=386 go build -o ../build/saman_386_darwin
echo 'building freebsd'
CGO_ENABLED=0 GOOS=freebsd GOARCH=amd64 go build -o ../build/saman_amd64_freebsd
CGO_ENABLED=0 GOOS=freebsd GOARCH=386 go build -o ../build/saman_386_freebsd
echo 'building windows'
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ../build/saman_amd64_windows.exe
CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -o ../docs/build/saman_386_windows.exe
echo 'build success'
