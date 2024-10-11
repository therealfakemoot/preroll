# Example usage
# ./build.sh && ./dist/demo
go generate ./...
rm -rf ./dist/*
find cmd/* -type d -exec go build -o dist/ "./"{} \;
