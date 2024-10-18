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
	TrdUpdateOrderFill         = 2218 // 推送成交通知
	TrdGetHistoryOrderList     = 2221 // 获取历史订单列表
	TrdGetHistoryOrderFillList = 2222 // 获取历史成交列表
	TrdGetMarginRatio          = 2223 // 获取融资融券数据
	TrdGetOrderFee             = 2225 // 获取订单费用

	// 行情接口
	QotSub                     = 3001 // 订阅或者反订阅
	QotGetSubInfo              = 3003 // 获取订阅信息
	QotGetBasicQot             = 3004 // 获取股票基本报价
	QotUpdateBasicQot          = 3005 // 推送股票基本报价
	QotGetKL                   = 3006 // 获取 K 线
	QotUpdateKL                = 3007 // 推送 K 线
	QotGetRT                   = 3008 // 获取分时
	QotUpdateRT                = 3009 // 推送分时
	QotGetTicker               = 3010 // 获取逐笔
	QotUpdateTicker            = 3011 // 推送逐笔
	QotGetOrderBook            = 3012 // 获取买卖盘
	QotUpdateOrderBook         = 3013 // 推送买卖盘
	QotGetBroker               = 3014 // 获取经纪队列
	QotUpdateBroker            = 3015 // 推送经纪队列
	QotUpdatePriceReminder     = 3019 // 到价提醒通知
	QotRequestHistoryKL        = 3103 // 在线获取单只股票一段历史 K 线
	QotRequestHistoryKLQuota   = 3104 // 获取历史 K 线额度
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
)
