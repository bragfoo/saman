rm -rf build
echo 'removed previous build'
cd api
echo 'building linux64'
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ../build/saman_linux_amd64
tar czf ../build/saman_linux_amd64.tar.gz ../build/saman_linux_amd64
echo 'building linux86'
CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -o ../build/saman_linux_386
tar czf ../build/saman_linux_386.tar.gz ../build/saman_linux_386
echo 'building linuxarm'
CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -o ../build/saman_linux_arm
tar czf ../build/saman_linux_arm.tar.gz ../build/saman_linux_arm
echo 'building MacOS64'
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o ../build/saman_darwin_amd64
tar czf ../build/saman_darwin_amd64.tar.gz ../build/saman_darwin_amd64
echo 'building MacOS86'
CGO_ENABLED=0 GOOS=darwin GOARCH=386 go build -o ../build/saman_darwin_386
tar czf ../build/saman_darwin_386.tar.gz ../build/saman_darwin_386
echo 'building freebsd'
CGO_ENABLED=0 GOOS=freebsd GOARCH=amd64 go build -o ../build/saman_freebsd_amd64
tar czf ../build/saman_freebsd_amd64.tar.gz ../build/saman_freebsd_amd64
CGO_ENABLED=0 GOOS=freebsd GOARCH=386 go build -o ../build/saman_freebsd_386
tar czf ../build/saman_freebsd_386.tar.gz ../build/saman_freebsd_386
echo 'building windows'
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ../build/saman_windows_amd64.exe
tar czf ../build/saman_windows_amd64.tar.gz ../build/saman_windows_amd64.exe
CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -o ../build/saman_windows_386.exe
tar czf ../build/saman_windows_386.tar.gz ../build/saman_windows_386.exe
echo 'build success'
