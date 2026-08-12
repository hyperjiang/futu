# Futu open api protobuf definitions

The source files are downloaded from https://openapi.futunn.com/futu-api-doc/quick/demo.html

Current version is `v10.9.6908`

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

Replace `github.com/futuopen/ftapi4go` with `github.com/hyperjiang/futu` in all the proto files.

Replace `github.com/hyperjiang/futu/pb/getoptionexpirationdate` with `github.com/hyperjiang/futu/pb/qotgetoptionexpirationdate`.

`Trd_PlaceComboOrder.proto` and `Trd_GetComboMaxTrdQtys.proto` share the go package with `Trd_PlaceOrder.proto` and
`Trd_GetMaxTrdQtys.proto`, which causes duplicated `C2S`/`S2C` declarations, so replace their go packages with
`github.com/hyperjiang/futu/pb/trdplacecomboorder` and `github.com/hyperjiang/futu/pb/trdgetcombomaxtrdqtys`.

### 4. Generate the codes

```bash
protoc -I=./.proto --go_out=/tmp ./.proto/*.proto;
cp -rf /tmp/github.com/hyperjiang/futu/pb/* ./pb
```
