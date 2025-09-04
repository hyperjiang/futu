# Futu open api protobuf definitions

The source files are downloaded from https://openapi.futunn.com/futu-api-doc/quick/demo.html

Current version is `v9.3.5308`

## How to generate go codes on MacOS

### 1. Install `protoc`

```bash
brew install protobuf
```

### 2. Install the compiler plugin `protoc-gen-go`

```bash
brew install protc-gen-go
```

### 3. Rename package

Replace `github.com/futuopen/ftapi4go` with `github.com/hyperjiang/futu` in all the proto files.

Replace `github.com/hyperjiang/futu/pb/getoptionexpirationdate` with `github.com/hyperjiang/futu/pb/qotgetoptionexpirationdate`.

### 4. Generate the codes

```bash
protoc -I=./.proto --go_out=/tmp ./.proto/*.proto;
cp -rf /tmp/github.com/hyperjiang/futu/pb/* ./pb
```
