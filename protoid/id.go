package protoid

const (
	// 基础功能
	InitConnect    = 1001 // 初始化连接
	GetGlobalState = 1002 // 获取全局状态
	Notify         = 1003 // 事件通知推送
	KeepAlive      = 1004 // 保活心跳

	// 交易接口
	TrdGetAccList              = 2001 // 获取交易业务账户列表
	TrdUnlockTrade             = 2005 // 解锁或锁定交易
	TrdSubAccPush              = 2008 // 订阅业务账户的交易推送数据
	TrdGetFunds                = 2101 // 获取账户资金
	TrdGetPositionList         = 2102 // 获取账户持仓
	TrdGetMaxTrdQtys           = 2111 // 获取最大交易数量
	TrdGetOrderList            = 2201 // 获取订单列表
	TrdPlaceOrder              = 2202 // 下单
	TrdModifyOrder             = 2205 // 修改订单
	TrdUpdateOrder             = 2208 // 推送订单状态变动通知
	TrdGetOrderFillList        = 2211 // 获取成交列表
	TrdUpdateOrderFill         = 2218 // 推送成交通知。该接口只支持实盘交易，不支持模拟交易。
	TrdGetHistoryOrderList     = 2221 // 获取历史订单列表
	TrdGetHistoryOrderFillList = 2222 // 获取历史成交列表
	TrdGetMarginRatio          = 2223 // 获取融资融券数据
	TrdGetOrderFee             = 2225 // 获取订单费用
	TrdFlowSummary             = 2226 // 查询账户现金流水

	// 行情接口
	QotSub                     = 3001 // 订阅或者反订阅
	QotGetSubInfo              = 3003 // 获取订阅信息
	QotGetBasicQot             = 3004 // 获取股票基本报价
	QotUpdateBasicQot          = 3005 // 推送股票基本报价。实时报价回调，异步处理已订阅股票的实时报价推送。
	QotGetKL                   = 3006 // 获取K线
	QotUpdateKL                = 3007 // 推送K线。实时K线回调，异步处理已订阅股票的实时K线推送。
	QotGetRT                   = 3008 // 获取分时
	QotUpdateRT                = 3009 // 推送分时。实时分时回调，异步处理已订阅股票的实时分时推送。
	QotGetTicker               = 3010 // 获取逐笔
	QotUpdateTicker            = 3011 // 推送逐笔。实时逐笔回调，异步处理已订阅股票的实时逐笔推送。
	QotGetOrderBook            = 3012 // 获取买卖盘
	QotUpdateOrderBook         = 3013 // 推送买卖盘。实时摆盘回调，异步处理已订阅股票的实时摆盘推送。
	QotGetBroker               = 3014 // 获取经纪队列
	QotUpdateBroker            = 3015 // 推送经纪队列。实时经纪队列回调，异步处理已订阅股票的实时经纪队列推送。
	QotUpdatePriceReminder     = 3019 // 到价提醒通知。到价提醒通知回调，异步处理已设置到价提醒的通知推送。
	QotRequestHistoryKL        = 3103 // 在线获取单只股票一段历史K线
	QotRequestHistoryKLQuota   = 3104 // 获取历史K线额度
	QotRequestRehab            = 3105 // 在线获取单只股票复权信息
	QotGetStaticInfo           = 3202 // 获取股票静态信息
	QotGetSecuritySnapshot     = 3203 // 获取股票快照
	QotGetPlateSet             = 3204 // 获取板块集合下的板块
	QotGetPlateSecurity        = 3205 // 获取板块下的股票
	QotGetReference            = 3206 // 获取正股相关股票
	QotGetOwnerPlate           = 3207 // 获取股票所属板块
	QotGetOptionChain          = 3209 // 获取期权链
	QotGetWarrant              = 3210 // 获取窝轮
	QotGetCapitalFlow          = 3211 // 获取资金流向
	QotGetCapitalDistribution  = 3212 // 获取资金分布
	QotGetUserSecurity         = 3213 // 获取自选股分组下的股票
	QotModifyUserSecurity      = 3214 // 修改自选股分组下的股票
	QotStockFilter             = 3215 // 获取条件选股
	QotGetIpoList              = 3217 // 获取新股
	QotGetFutureInfo           = 3218 // 获取期货合约资料
	QotRequestTradeDate        = 3219 // 获取市场交易日，在线拉取不在本地计算
	QotSetPriceReminder        = 3220 // 设置到价提醒
	QotGetPriceReminder        = 3221 // 获取到价提醒
	QotGetUserSecurityGroup    = 3222 // 获取自选股分组列表
	QotGetMarketState          = 3223 // 获取指定品种的市场状态
	QotGetOptionExpirationDate = 3224 // 获取期权到期日

	// v10.6 新增行情接口
	QotGetFinancialsEarningsPriceMove    = 3225 // 获取财报盈利预测变动
	QotGetFinancialsEarningsPriceHistory = 3226 // 获取财报盈利预测历史
	QotGetFinancialsStatements           = 3227 // 获取财务报表
	QotGetFinancialsRevenueBreakdown     = 3228 // 获取主营构成
	QotGetResearchAnalystConsensus       = 3229 // 获取分析师评级汇总
	QotGetResearchRatingSummary          = 3230 // 获取分析师评级详情
	QotGetResearchMorningstarReport      = 3231 // 获取晨星研报
	QotGetValuationDetail                = 3232 // 获取估值详情
	QotGetValuationPlateStockList        = 3233 // 获取板块估值股票列表
	QotGetCorporateActionsDividends      = 3234 // 获取除权除息
	QotGetCorporateActionsBuybacks       = 3235 // 获取回购
	QotGetCorporateActionsStockSplits    = 3236 // 获取拆合股
	QotGetShareholdersOverview           = 3237 // 获取股东概要
	QotGetShareholdersHoldingChanges     = 3238 // 获取持股变动
	QotGetShareholdersHolderDetail       = 3239 // 获取股东持仓明细
	QotGetShareholdersInstitutional      = 3240 // 获取机构持仓
	QotGetInsiderHolderList              = 3241 // 获取内部持有人列表
	QotGetInsiderTradeList               = 3242 // 获取内部人交易列表
	QotGetCompanyProfile                 = 3243 // 获取公司资料
	QotGetCompanyExecutives              = 3244 // 获取公司高管
	QotGetCompanyExecutiveBackground     = 3245 // 获取高管背景
	QotGetCompanyOperationalEfficiency   = 3246 // 获取营运效率
	QotGetTopTenBuySellBrokers           = 3247 // 获取十大经纪商
	QotGetDailyShortVolume               = 3248 // 获取每日做空量
	QotGetShortInterest                  = 3249 // 获取做空比例
	QotGetOptionVolatility               = 3250 // 获取期权波动率
	QotGetOptionExerciseProbability      = 3251 // 获取期权行权概率
	QotStockScreen                       = 3252 // 条件选股
	QotOptionScreen                      = 3253 // 期权筛选
	QotWarrantScreen                     = 3254 // 窝轮筛选
)
