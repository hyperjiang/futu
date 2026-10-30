# Futu open api protobuf definitions

The source files are downloaded from https://openapi.futunn.com/futu-api-doc/quick/demo.html

Current version is `v10.12.7208`

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

Add `option go_package = "github.com/hyperjiang/futu/pb/qotoptioncommon";` to `Qot_OptionCommon.proto`

Add `option go_package = "github.com/hyperjiang/futu/pb/qotgetearningscalendar";` to `Qot_GetEarningsCalendar.proto`

Add `option go_package = "github.com/hyperjiang/futu/pb/qotgetoptionearningsscreener";` to `Qot_GetOptionEarningsScreener.proto`

Add `option go_package = "github.com/hyperjiang/futu/pb/qotgetoptionevent";` to `Qot_GetOptionEvent.proto`

Add `option go_package = "github.com/hyperjiang/futu/pb/qotgetoptioneventalert";` to `Qot_GetOptionEventAlert.proto`

Add `option go_package = "github.com/hyperjiang/futu/pb/qotgetoptionmarketstatistic";` to `Qot_GetOptionMarketStatistic.proto`

Add `option go_package = "github.com/hyperjiang/futu/pb/qotgetoptionrank";` to `Qot_GetOptionRank.proto`

Add `option go_package = "github.com/hyperjiang/futu/pb/qotgetoptionsellerscreener";` to `Qot_GetOptionSellerScreener.proto`

Add `option go_package = "github.com/hyperjiang/futu/pb/qotgetoptionunderlyinghisstatistic";` to `Qot_GetOptionUnderlyingHisStatistic.proto`

Add `option go_package = "github.com/hyperjiang/futu/pb/qotgetoptionunderlyinghisvolatility";` to `Qot_GetOptionUnderlyingHisVolatility.proto`

Add `option go_package = "github.com/hyperjiang/futu/pb/qotgetoptionunderlyingoverview";` to  `Qot_GetOptionUnderlyingOverview.proto`

Add `option go_package = "github.com/hyperjiang/futu/pb/qotgetoptionunderlyingrank";` to `Qot_GetOptionUnderlyingRank.proto`

Add `option go_package = "github.com/hyperjiang/futu/pb/qotgetoptionzerodtecontract";` to `Qot_GetOptionZeroDteContract.proto`

Add `option go_package = "github.com/hyperjiang/futu/pb/qotgetoptionzerodtescreener";` to `Qot_GetOptionZeroDteScreener.proto`

Add `option go_package = "github.com/hyperjiang/futu/pb/qotsetoptioneventalert";` to `Qot_SetOptionEventAlert.proto`

Add `option go_package = "github.com/hyperjiang/futu/pb/qotupdateoptionevent";` to `Qot_UpdateOptionEvent.proto`

### 4. Generate the codes

```bash
protoc -I=./.proto --go_out=/tmp ./.proto/*.proto;
cp -rf /tmp/github.com/hyperjiang/futu/pb/* ./pb
```
