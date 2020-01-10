# build cross compatible binaries to publish

# create a bin directory if one does not already exist
if [ ! -d "bin" ]; then
  echo "Directory created for binaries: ./bin/"
  mkdir bin
fi

# build for osx
env GOOS=darwin GOARCH=amd64 go build main.go
mv main bin/darwin-amd64
echo "Binary created for OSX: ./bin/darwin-amd64"

# build for linux
env GOOS=linux GOARCH=amd64 go build main.go
mv main bin/linux-amd64
echo "Binary created for Linux: ./bin/linux-amd64"

# build for windows
env GOOS=windows GOARCH=amd64 go build main.go
mv main.exe bin/windows-amd64.exe
echo "Binary created for Windows: ./bin/windows-amd64.exe"
