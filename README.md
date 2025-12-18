# Futu open api golang client

[![GoDoc](https://pkg.go.dev/badge/github.com/hyperjiang/futu)](https://pkg.go.dev/github.com/hyperjiang/futu)
[![CI](https://github.com/hyperjiang/futu/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/hyperjiang/futu/actions/workflows/ci.yml)
[![](https://goreportcard.com/badge/github.com/hyperjiang/futu)](https://goreportcard.com/report/github.com/hyperjiang/futu)
[![codecov](https://codecov.io/gh/hyperjiang/futu/graph/badge.svg?token=iI7hyTEenz)](https://codecov.io/gh/hyperjiang/futu)
[![Release](https://img.shields.io/github/release/hyperjiang/futu.svg)](https://github.com/hyperjiang/futu/releases)

富途牛牛 OpenAPI Golang 客户端，要求Golang版本 >= 1.21

Futu Open API 官方文档: https://openapi.futunn.com/futu-api-doc/

## 代码目录说明

- `根目录`: 提供用户友好的客户端SDK，对底层的`client`包做了一层用法包装
- `client`: 基础客户端，拥有所有功能，需要用pb定义的结构体传参，可以直接使用，但是使用起来略繁琐
- `.proto`: protobuf 定义文件
- `adapt`: protobuf 结构体和普通类型的适配层
- `infra`: 底层支持库，使用者无需关心
- `pb`: 基于 protobuf 文件生成的 golang 代码
- `protoid`: 接口ID常量列表

## 使用说明

本SDK跟`FutuOpenD`的通信协议格式是protobuf，使用`FutuOpenD`的默认配置即可，不要改成json格式。

具体用法可以参考单元测试，里面每个接口都有用例。

方法设计的原则是，必填参数显式要求传递，可选参数放在可变参数`adapt.Option`传递。但有部分必填参数实际上可以有默认值，这种情况不会显式要求传递。

每个接口都有一对方法，其中结尾有`WithContext`的方法是方便让使用者可以自己通过context设置超时时间，不带`WithContext`的方法使用了默认超时时间5s。

以下是一个简单示例：

```go
import "github.com/hyperjiang/futu"

sdk, err := futu.NewSDK()
if err != nil {
    log.Fatal(err)
}

res, err := sdk.GetGlobalState()
fmt.Println(res)
```

对于系统推送过来的数据，需要调用`RegisterHandler(protoID uint32, h Handler)`来注册自己的处理逻辑。
如果没有设置，SDK会使用默认的Handler，只打印收到的消息到日志。
可以设置推送Handler的协议ID如下:

- protoid.Notify // 1003
- protoid.TrdUpdateOrder // 2208
- protoid.TrdUpdateOrderFill // 2218
- protoid.QotUpdateBasicQot // 3005
- protoid.QotUpdateKL // 3007
- protoid.QotUpdateRT // 3009
- protoid.QotUpdateTicker // 3011
- protoid.QotUpdateOrderBook // 3013
- protoid.QotUpdateBroker // 3015
- protoid.QotUpdatePriceReminder // 3019

设置其他ID没有任何作用，因为永远不可能触发到。

## 支持的功能

### 基础功能（用户无需调用）
- [x] 支持使用RSA和AES加解密
- [x] 初始化连接
- [x] 获取全局状态
- [x] 事件通知推送
- [x] 保活心跳

### 行情接口
- [x] QotSub                     = 3001 // 订阅或者反订阅
- [x] QotGetSubInfo              = 3003 // 获取订阅信息
- [x] QotGetBasicQot             = 3004 // 获取股票基本报价
- [x] QotUpdateBasicQot          = 3005 // 推送股票基本报价
- [x] QotGetKL                   = 3006 // 获取K线
- [x] QotUpdateKL                = 3007 // 推送K线
- [x] QotGetRT                   = 3008 // 获取分时
- [x] QotUpdateRT                = 3009 // 推送分时
- [x] QotGetTicker               = 3010 // 获取逐笔
- [x] QotUpdateTicker            = 3011 // 推送逐笔
- [x] QotGetOrderBook            = 3012 // 获取买卖盘
- [x] QotUpdateOrderBook         = 3013 // 推送买卖盘
- [x] QotGetBroker               = 3014 // 获取经纪队列
- [x] QotUpdateBroker            = 3015 // 推送经纪队列
- [x] QotUpdatePriceReminder     = 3019 // 到价提醒通知
- [x] QotRequestHistoryKL        = 3103 // 在线获取单只股票一段历史K线
- [x] QotRequestHistoryKLQuota   = 3104 // 获取历史K线额度
- [x] QotRequestRehab            = 3105 // 在线获取单只股票复权信息
- [x] QotGetStaticInfo           = 3202 // 获取股票静态信息
- [x] QotGetSecuritySnapshot     = 3203 // 获取股票快照
- [x] QotGetPlateSet             = 3204 // 获取板块集合下的板块
- [x] QotGetPlateSecurity        = 3205 // 获取板块下的股票
- [x] QotGetReference            = 3206 // 获取正股相关股票
- [x] QotGetOwnerPlate           = 3207 // 获取股票所属板块
- [x] QotGetOptionChain          = 3209 // 获取期权链
- [x] QotGetWarrant              = 3210 // 获取窝轮
- [x] QotGetCapitalFlow          = 3211 // 获取资金流向
- [x] QotGetCapitalDistribution  = 3212 // 获取资金分布
- [x] QotGetUserSecurity         = 3213 // 获取自选股分组下的股票
- [x] QotModifyUserSecurity      = 3214 // 修改自选股分组下的股票
- [x] QotStockFilter             = 3215 // 获取条件选股
- [x] QotGetIpoList              = 3217 // 获取新股
- [x] QotGetFutureInfo           = 3218 // 获取期货合约资料
- [x] QotRequestTradeDate        = 3219 // 获取市场交易日，在线拉取不在本地计算
- [x] QotSetPriceReminder        = 3220 // 设置到价提醒
- [x] QotGetPriceReminder        = 3221 // 获取到价提醒
- [x] QotGetUserSecurityGroup    = 3222 // 获取自选股分组列表
- [x] QotGetMarketState          = 3223 // 获取指定品种的市场状态
- [x] QotGetOptionExpirationDate = 3224 // 获取期权到期日

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
- [x] TrdUpdateOrder             = 2208 // 推送订单状态变动通知
- [x] TrdGetOrderFillList        = 2211 // 获取成交列表
- [x] TrdUpdateOrderFill         = 2218 // 推送成交通知
- [x] TrdGetHistoryOrderList     = 2221 // 获取历史订单列表
- [x] TrdGetHistoryOrderFillList = 2222 // 获取历史成交列表
- [x] TrdGetMarginRatio          = 2223 // 获取融资融券数据
- [x] TrdGetOrderFee             = 2225 // 获取订单费用
- [x] TrdFlowSummary             = 2226 // 查询账户现金流水
