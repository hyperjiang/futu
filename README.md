# Futu open api golang client

[![GoDoc](https://pkg.go.dev/badge/github.com/hyperjiang/futu)](https://pkg.go.dev/github.com/hyperjiang/futu)
[![CI](https://github.com/hyperjiang/futu/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/hyperjiang/futu/actions/workflows/ci.yml)
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
- protoid.QotPushIndicatorCalc // 3261
- protoid.QotUpdateOptionEvent // 3310
- protoid.QotUpdateEventContractOrderBook // 3450
- protoid.QotUpdateEventContractKline // 3451
- protoid.QotUpdateEventContractTicker // 3452

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
- [x] QotGetFinancialsEarningsPriceMove    = 3225 // 获取财报盈利预测变动
- [x] QotGetFinancialsEarningsPriceHistory = 3226 // 获取财报盈利预测历史
- [x] QotGetFinancialsStatements           = 3227 // 获取财务报表
- [x] QotGetFinancialsRevenueBreakdown     = 3228 // 获取主营构成
- [x] QotGetResearchAnalystConsensus       = 3229 // 获取分析师评级汇总
- [x] QotGetResearchRatingSummary          = 3230 // 获取分析师评级详情
- [x] QotGetResearchMorningstarReport      = 3231 // 获取晨星研报
- [x] QotGetValuationDetail                = 3232 // 获取估值详情
- [x] QotGetValuationPlateStockList        = 3233 // 获取板块估值股票列表
- [x] QotGetCorporateActionsDividends      = 3234 // 获取除权除息
- [x] QotGetCorporateActionsBuybacks       = 3235 // 获取回购
- [x] QotGetCorporateActionsStockSplits    = 3236 // 获取拆合股
- [x] QotGetShareholdersOverview           = 3237 // 获取股东概要
- [x] QotGetShareholdersHoldingChanges     = 3238 // 获取持股变动
- [x] QotGetShareholdersHolderDetail       = 3239 // 获取股东持仓明细
- [x] QotGetShareholdersInstitutional      = 3240 // 获取机构持仓
- [x] QotGetInsiderHolderList              = 3241 // 获取内部持有人列表
- [x] QotGetInsiderTradeList               = 3242 // 获取内部人交易列表
- [x] QotGetCompanyProfile                 = 3243 // 获取公司资料
- [x] QotGetCompanyExecutives              = 3244 // 获取公司高管
- [x] QotGetCompanyExecutiveBackground     = 3245 // 获取高管背景
- [x] QotGetCompanyOperationalEfficiency   = 3246 // 获取营运效率
- [x] QotGetTopTenBuySellBrokers           = 3247 // 获取十大经纪商
- [x] QotGetDailyShortVolume               = 3248 // 获取每日做空量
- [x] QotGetShortInterest                  = 3249 // 获取做空比例
- [x] QotGetOptionVolatility               = 3250 // 获取期权波动率
- [x] QotGetOptionExerciseProbability      = 3251 // 获取期权行权概率
- [x] QotStockScreen                       = 3252 // 条件选股
- [x] QotOptionScreen                      = 3253 // 期权筛选
- [x] QotWarrantScreen                     = 3254 // 窝轮筛选
- [x] QotGetOptionQuote                    = 3255 // 获取期权行情
- [x] QotGetOptionStrategy                 = 3256 // 获取期权策略
- [x] QotGetOptionStrategyAnalysis         = 3257 // 获取期权策略分析
- [x] QotGetOptionStrategySpread           = 3258 // 获取期权策略价差
- [x] QotGetIndicatorList                  = 3259 // 获取指标列表
- [x] QotRequestIndicatorCalc              = 3260 // 异步发起指标计算
- [x] QotPushIndicatorCalc                 = 3261 // 指标异步计算结果推送
- [x] QotGetSearchQuote                    = 3262 // 搜索行情
- [x] QotGetSearchNews                     = 3263 // 搜索资讯
- [x] QotGetOptionMarketStatistic          = 3301 // 获取期权市场统计
- [x] QotGetOptionUnderlyingHisStatistic   = 3302 // 获取期权标的历史统计
- [x] QotGetOptionUnderlyingOverview       = 3303 // 获取批量标的最新数据
- [x] QotGetOptionUnderlyingHisVolatility  = 3304 // 获取历史波动率
- [x] QotGetOptionUnderlyingRank           = 3305 // 获取标的排行
- [x] QotGetOptionRank                     = 3306 // 获取期权合约排行
- [x] QotGetOptionEvent                    = 3307 // 获取期权异动列表
- [x] QotGetOptionEventAlert               = 3308 // 获取期权异动告警设置
- [x] QotSetOptionEventAlert               = 3309 // 修改期权异动告警条件
- [x] QotUpdateOptionEvent                 = 3310 // 期权异动推送
- [x] QotGetOptionZeroDteScreener          = 3311 // 获取末日期权标的列表
- [x] QotGetOptionZeroDteContract          = 3312 // 获取末日期权合约列表
- [x] QotGetOptionEarningsScreener         = 3313 // 获取财报期权标的列表
- [x] QotGetOptionSellerScreener           = 3314 // 获取期权卖方策略列表
- [x] QotGetEarningsCalendar               = 3401 // 获取财报日历
- [x] QotGetMacroIndicatorList             = 3402 // 获取宏观指标列表
- [x] QotGetMacroIndicatorHistory          = 3403 // 获取宏观指标历史数据
- [x] QotGetFedWatchTargetRate             = 3404 // 获取美联储利率预测
- [x] QotGetFedWatchDotPlot                = 3405 // 获取CME利率点阵图
- [x] QotGetEarningsBeatRank               = 3406 // 获取盈利超预期排行
- [x] QotGetDividendRank                   = 3407 // 获取股息排行
- [x] QotGetDividendCalendar               = 3408 // 获取派息日历
- [x] QotGetEconomicCalendar               = 3409 // 获取财经日历
- [x] QotGetUSPreMarketRank                = 3410 // 获取盘前榜
- [x] QotGetUSAfterHoursRank               = 3411 // 获取盘后榜
- [x] QotGetUSOvernightRank                = 3412 // 获取夜盘榜
- [x] QotGetTopMoversRank                  = 3413 // 获取领涨领跌榜
- [x] QotGetHotList                        = 3414 // 获取热议榜
- [x] QotGetShortSellingRank               = 3415 // 获取卖空异动榜
- [x] QotGetPeriodChangeRank               = 3416 // 获取区间涨跌幅
- [x] QotGetHighDividendSOERank            = 3417 // 获取破净高股息国央企
- [x] QotGetInstitutionList                = 3418 // 获取机构持仓列表
- [x] QotGetInstitutionProfile             = 3419 // 获取机构概况
- [x] QotGetInstitutionDistribution        = 3420 // 获取机构持仓行业分布
- [x] QotGetInstitutionHoldingChange       = 3421 // 获取机构持仓变动
- [x] QotGetInstitutionHoldingList         = 3422 // 获取机构持股列表
- [x] QotGetArkFundHolding                 = 3423 // 获取ARK基金持仓
- [x] QotGetArkStockDynamic                = 3424 // 获取ARK个股交易动态
- [x] QotGetArkActiveTransaction           = 3425 // 获取ARK主动交易聚合
- [x] QotGetRatingChange                   = 3426 // 获取评级变动
- [x] QotGetIndustrialChainList            = 3427 // 获取产业链列表
- [x] QotGetIndustrialChainDetail          = 3428 // 获取产业链详情
- [x] QotGetIndustrialChainByPlate         = 3429 // 获取板块关联产业链
- [x] QotGetIndustrialPlateInfo            = 3430 // 获取产业板块信息
- [x] QotGetIndustrialPlateStock           = 3431 // 获取产业板块成分股
- [x] QotGetHeatMapData                    = 3432 // 获取热力图数据
- [x] QotGetRiseFallDistribution           = 3433 // 获取涨跌分布
- [x] QotGetEventContractCategory          = 3434 // 获取事件合约分类
- [x] QotFilterCompetition                 = 3435 // 获取赛事筛选选项
- [x] QotGetEventContractSeriesList        = 3436 // 获取事件合约Series列表
- [x] QotGetEventContractEventList         = 3437 // 获取事件合约Event列表
- [x] QotGetEventContract                  = 3438 // 获取事件合约
- [x] QotGetEventContractMilestoneList     = 3439 // 获取事件合约里程碑列表
- [x] QotGetEventContractSnapshot          = 3445 // 获取事件合约快照
- [x] QotGetEventContractOrderBook         = 3446 // 获取事件合约摆盘
- [x] QotGetEventContractKline             = 3447 // 获取事件合约K线
- [x] QotGetEventContractTicker            = 3448 // 获取事件合约逐笔
- [x] QotUpdateEventContractOrderBook      = 3450 // 事件合约摆盘推送
- [x] QotUpdateEventContractKline          = 3451 // 事件合约K线推送
- [x] QotUpdateEventContractTicker         = 3452 // 事件合约逐笔推送
- [x] QotGetEventContractComboList         = 3453 // 获取事件合约可Combo列表
- [x] QotGetEventContractComboRfq          = 3454 // 事件合约Combo询价
- [x] QotSubEventContract                  = 3455 // 事件合约订阅/反订阅
- [x] QotRequestHistoryEventContractKL     = 3456 // 拉取事件合约历史K线

### 交易接口
- [x] TrdGetAccList              = 2001 // 获取交易业务账户列表
- [x] TrdUnlockTrade             = 2005 // 解锁或锁定交易
- [x] TrdSubAccPush              = 2008 // 订阅业务账户的交易推送数据
- [x] TrdGetFunds                = 2101 // 获取账户资金
- [x] TrdGetPositionList         = 2102 // 获取账户持仓
- [x] TrdGetMaxTrdQtys           = 2111 // 获取最大交易数量
- [x] TrdGetComboMaxTrdQtys      = 2112 // 获取组合的可买卖信息
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
- [x] TrdPlaceComboOrder         = 2227 // 组合期权下单
