# AGENTS.md

Go SDK for [Futu OpenAPI](https://openapi.futunn.com/futu-api-doc/) — programmatic access to Futu's stock trading and market data platform (Hong Kong, US, China, Singapore, Japan markets). Communicates with `FutuOpenD` gateway over TCP using protobuf.

## Commands

```bash
go test ./...                    # run all tests (requires FutuOpenD running locally)
go test ./adapt/...              # unit tests only (no external dependency)
go test ./infra/...              # unit tests only
go test -run TestNewSecurity     # single test
go vet ./...                     # static analysis
```

**Tests require a live FutuOpenD instance** on `localhost:11111` with RSA keys configured. The `adapt/` and `infra/` packages have pure unit tests that run without connectivity. All other test suites (`client/`, root `sdk_test.go`) are integration tests that will skip or fail without a running gateway.

## Architecture

Three-layer design with clear separation of concerns:

```
Root package (futu/)     ← User-facing SDK, friendly types, 5s default timeout
    ↓ delegates to
client/                  ← Low-level protobuf client, raw C2S/S2C structs
    ↓ communicates with
FutuOpenD (TCP:11111)    ← Binary protocol: 'FT' header + protobuf body + AES-CBC encryption
```

**Supporting packages:**
- `adapt/` — Adapter/bridge layer: converts user-friendly types (string codes like `"HK.00700"`) to protobuf structs, re-exports protobuf enum constants as plain `int32`, provides the `Options` map pattern
- `pb/` — Auto-generated protobuf Go code (do NOT edit manually)
- `protoid/` — Protocol ID constants (e.g., `QotGetKL = 3006`)
- `infra/` — Internal plumbing: `Crypto` (AES-CBC), `DispatcherHub` (routes responses by protoID+serialNo), `ProtobufChan` (type-safe generic channel via reflection)

### Data/Control Flow

1. SDK method builds `adapt.Options` map → converts to protobuf C2S via `json.Marshal`/`json.Unmarshal` (`adapt/option.go:88`)
2. Client wraps C2S in a `Request` proto, encrypts body (RSA for InitConnect, AES-CBC for all else), writes binary header + body to TCP conn
3. Response read loop (`infiniteRead`) decodes header, decrypts, dispatches to registered channel via `DispatcherHub` keyed on `(protoID, serialNo)`
4. Push notifications use `serialNo = 0` and are routed to handler functions registered via `RegisterHandler`

## Key Patterns

### Dual-method convention
Every API has two variants:
- `MethodName()` — uses 5-second default timeout (`sdk.go`)
- `MethodNameWithContext(ctx)` — caller-controlled timeout (`sdk_ctx.go`)

When adding a new API, implement the `WithContext` variant in `sdk_ctx.go` first, then add the convenience wrapper in `sdk.go`.

### Options pattern (`adapt.Option`)
Optional parameters use `adapt.With("key", value)` functional options. Internally these populate a `map[string]any` which is serialized to protobuf via JSON round-trip:

```go
sdk.GetKL("HK.09988", adapt.KLType_Day,
    adapt.With("rehabType", adapt.RehabType_Forward),
    adapt.With("reqNum", 3),
)
```

**Gotcha**: The JSON round-trip (`Options.ToProto`) means option keys must match protobuf JSON field names exactly (camelCase). Mismatched keys are silently ignored.

### Security codes
Always `"MARKET.CODE"` format: `"HK.00700"`, `"US.AAPL"`, `"SH.600000"`. The `adapt.NewSecurity()` helper splits on `.` — returns `nil` if no dot found.

### Trade headers
Trading APIs require a `*trdcommon.TrdHeader` specifying environment (real vs simulation), account ID, and market. Use `adapt.NewTradeHeader()` or `adapt.NewSimulationTradeHeader()`.

### Constants re-export
The `adapt` package re-exports all protobuf enum values as plain `int32` constants (e.g., `adapt.TrdMarket_HK`, `adapt.KLType_Day`, `adapt.OrderType_Normal`). Use these rather than referencing protobuf enum types directly.

## Gotchas

- **Options map key naming**: Keys in `adapt.With()` must use the protobuf JSON name (camelCase). E.g., `adapt.With("filterStatusList", ...)` works; `adapt.With("FilterStatusList", ...)` does not.
- **CodeForTrade splitting**: `Options.SetCodeForTrade()` splits `"US.AAPL"` into `secMarket` + `code` fields. If the code has no dot, it's passed as-is — this may cause unexpected behavior for malformed codes.
- **Default values in WithContext methods**: Some methods set defaults (e.g., `rehabType`, `reqNum`, `maxRetNum`, `num`, `begin`) when the option key is missing. Check `sdk_ctx.go` for each method's defaults.
- **Simulated accounts**: Several trade APIs don't work with simulation accounts (order fills, order fees, flow summary). Tests assert `should.Error(err)` for these cases.
- **Push handler registration**: Only specific protoIDs can be registered as push handlers (see README list). Registering other IDs has no effect.
- **Proto regeneration**: After updating `.proto` files, run `protoc -I=./.proto --go_out=/tmp ./.proto/*.proto` then `cp -rf /tmp/github.com/hyperjiang/futu/pb/* ./pb`. See `.proto/README.md`.
- **Encryption**: InitConnect uses RSA; all subsequent messages use AES-CBC with key/IV from the InitConnect response. If no RSA keys are provided, communication is unencrypted.
- **`adapt.Options.ToProto` uses JSON**: This is the bridge between the map-based options and protobuf structs. Fields not present in the JSON (or with wrong types) are silently zeroed.

## Code Style

- Standard Go conventions, `gofmt`-compatible
- Error returns: `(result, error)` pairs; methods returning only `error` for void operations
- Protobuf field access via generated getters (`GetXxx()`)
- Logging via `zerolog` (`github.com/rs/zerolog/log`)
- Tests use `testify/require` with `should := require.New(t)` idiom (not `assert`)
- Test suites use `testify/suite` for integration tests requiring setup/teardown
- Chinese comments in `adapt/` constants and `client/` methods (these are the official Futu API descriptions)
