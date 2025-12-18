# Futu open api protobuf definitions

The source files are downloaded from https://openapi.futunn.com/futu-api-doc/quick/demo.html

Current version is `v9.4.5418`

## How to generate go codes on MacOS

### 1. Install `protoc`

```bash
brew install protobuf
```

### 2. Install the compiler plugin `protoc-gen-go`

```bash
brew install protoc-gen-go
```

### 3. Rename package

Replace `github.com/hyperjiang/futu` with `github.com/hyperjiang/futu` in all the proto files.

Replace `github.com/hyperjiang/futu/pb/qotgetoptionexpirationdate` with `github.com/hyperjiang/futu/pb/qotgetoptionexpirationdate`.

### 4. Generate the codes

```bash
protoc -I=./.proto --go_out=/tmp ./.proto/*.proto;
cp -rf /tmp/github.com/hyperjiang/futu/pb/* ./pb
```
