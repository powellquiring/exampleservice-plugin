# print the checksums for each binary file
# these are used when publishling plugins

declare -a binaries=(
  "./bin/darwin-amd64"
  "./bin/linux-amd64"
  "./bin/windows-amd64.exe"
)

for binary in "${binaries[@]}"
do
  if [ -f $binary ]; then
    shasum -a 1 $binary
  else
    echo "File not found: $binary"
  fi
done
