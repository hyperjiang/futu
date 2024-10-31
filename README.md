# Futu open api golang client

[![GoDoc](https://pkg.go.dev/badge/github.com/hyperjiang/futu)](https://pkg.go.dev/github.com/hyperjiang/futu)
[![CI](https://github.com/hyperjiang/futu/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/hyperjiang/futu/actions/workflows/ci.yml)
[![](https://goreportcard.com/badge/github.com/hyperjiang/futu)](https://goreportcard.com/report/github.com/hyperjiang/futu)
[![codecov](https://codecov.io/gh/hyperjiang/futu/graph/badge.svg?token=iI7hyTEenz)](https://codecov.io/gh/hyperjiang/futu)
[![Release](https://img.shields.io/github/release/hyperjiang/futu.svg)](https://github.com/hyperjiang/futu/releases)

Futu open api golang client. Require go version >= 1.21.

API doc: https://openapi.futunn.com/futu-api-doc/

做最好用的富途牛牛 OpenAPI Golang 客户端。

## 代码目录说明

- 根目录: 提供用户友好的客户端
- `client`: 基础客户端，拥有所有功能，需要用pb定义的结构体传参，可以直接使用，但是使用起来略繁琐
- `.proto`: protobuf 定义文件
- `infra`: 底层支持库，使用者无需关心
- `pb`: 基于 protobuf 文件生成的 golang 代码
- `protoid`: 接口ID常量列表

## Usage

```bash
go get -u github.com/hyperjiang/futu
```

```go
import "github.com/hyperjiang/futu"

client, err := futu.NewClient()
if err != nil {
    log.Fatal(err)
}

ctx, cancel := context.WithTimeout(context.Background(), time.Second)
defer cancel()
res, err := client.GetGlobalState(ctx)
fmt.Println(res)
```

## Features

### 基础功能
- [x] 支持使用RSA加解密
- [x] 初始化连接
- [x] 获取全局状态
- [x] 事件通知推送
- [x] 保活心跳

### 行情接口
- [x] QotSub                     = 3001 // 订阅或者反订阅
- [x] QotGetSubInfo              = 3003 // 获取订阅信息
- [x] QotGetBasicQot             = 3004 // 获取股票基本报价
- [ ] QotUpdateBasicQot          = 3005 // 推送股票基本报价
- [x] QotGetKL                   = 3006 // 获取K线
- [ ] QotUpdateKL                = 3007 // 推送K线
- [ ] QotGetRT                   = 3008 // 获取分时
- [ ] QotUpdateRT                = 3009 // 推送分时
- [ ] QotGetTicker               = 3010 // 获取逐笔
- [ ] QotUpdateTicker            = 3011 // 推送逐笔
- [ ] QotGetOrderBook            = 3012 // 获取买卖盘
- [ ] QotUpdateOrderBook         = 3013 // 推送买卖盘
- [ ] QotGetBroker               = 3014 // 获取经纪队列
- [ ] QotUpdateBroker            = 3015 // 推送经纪队列
- [ ] QotUpdatePriceReminder     = 3019 // 到价提醒通知
- [x] QotRequestHistoryKL        = 3103 // 在线获取单只股票一段历史K线
- [ ] QotRequestHistoryKLQuota   = 3104 // 获取历史K线额度
- [ ] QotRequestRehab            = 3105 // 在线获取单只股票复权信息
- [ ] QotGetStaticInfo           = 3202 // 获取股票静态信息
- [x] QotGetSecuritySnapshot     = 3203 // 获取股票快照
- [ ] QotGetPlateSet             = 3204 // 获取板块集合下的板块
- [ ] QotGetPlateSecurity        = 3205 // 获取板块下的股票
- [ ] QotGetReference            = 3206 // 获取正股相关股票
- [ ] QotGetOwnerPlate           = 3207 // 获取股票所属板块
- [ ] QotGetOptionChain          = 3209 // 获取期权链
- [ ] QotGetWarrant              = 3210 // 获取窝轮
- [ ] QotGetCapitalFlow          = 3211 // 获取资金流向
- [ ] QotGetCapitalDistribution  = 3212 // 获取资金分布
- [ ] QotGetUserSecurity         = 3213 // 获取自选股分组下的股票
- [ ] QotModifyUserSecurity      = 3214 // 修改自选股分组下的股票
- [x] QotStockFilter             = 3215 // 获取条件选股
- [ ] QotGetIpoList              = 3217 // 获取新股
- [ ] QotGetFutureInfo           = 3218 // 获取期货合约资料
- [ ] QotRequestTradeDate        = 3219 // 获取市场交易日，在线拉取不在本地计算
- [ ] QotSetPriceReminder        = 3220 // 设置到价提醒
- [ ] QotGetPriceReminder        = 3221 // 获取到价提醒
- [ ] QotGetUserSecurityGroup    = 3222 // 获取自选股分组列表
- [ ] QotGetMarketState          = 3223 // 获取指定品种的市场状态
- [ ] QotGetOptionExpirationDate = 3224 // 获取期权到期日

### 交易接口
- [x] TrdGetAccList              = 2001 // 获取交易业务账户列表
- [x] TrdUnlockTrade             = 2005 // 解锁或锁定交易
- [x] TrdSubAccPush              = 2008 // 订阅业务账户的交易推送数据
- [x] TrdGetFunds                = 2101 // 获取账户资金
- [x] TrdGetPositionList         = 2102 // 获取账户持仓
- [x] TrdGetMaxTrdQtys           = 2111 // 获取最大交易数量
- [x] TrdGetOrderList            = 2201 // 获取订单列表
- [x] TrdPlaceOrder              = 2202 // 下单
- [x] TrdModifyOrder             = 2205 // 修改订单
- [ ] TrdUpdateOrder             = 2208 // 推送订单状态变动通知
- [x] TrdGetOrderFillList        = 2211 // 获取成交列表
- [ ] TrdUpdateOrderFill         = 2218 // 推送成交通知
- [x] TrdGetHistoryOrderList     = 2221 // 获取历史订单列表
- [x] TrdGetHistoryOrderFillList = 2222 // 获取历史成交列表
- [x] TrdGetMarginRatio          = 2223 // 获取融资融券数据
- [x] TrdGetOrderFee             = 2225 // 获取订单费用

