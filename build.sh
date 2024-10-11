# Example usage
# ./build.sh && ./dist/demo
go generate ./...
find cmd/* -type d -exec go build -o dist/ "./"{} \;
