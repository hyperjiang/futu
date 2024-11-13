package adapt

import (
	"strings"

	"github.com/hyperjiang/futu/pb/trdcommon"
)

// 交易市场，主要用于交易协议公共参数头和交易业务账户结构
const (
	// TrdMarket_Unknown 未知市场
	TrdMarket_Unknown = int32(trdcommon.TrdMarket_TrdMarket_Unknown)

	// TrdMarket_HK 香港市场
	TrdMarket_HK = int32(trdcommon.TrdMarket_TrdMarket_HK)

	// TrdMarket_US 美国市场
	TrdMarket_US = int32(trdcommon.TrdMarket_TrdMarket_US)

	// TrdMarket_CN 大陆市场
	TrdMarket_CN = int32(trdcommon.TrdMarket_TrdMarket_CN)

	// TrdMarket_HKCC 香港A股通市场
	TrdMarket_HKCC = int32(trdcommon.TrdMarket_TrdMarket_HKCC)

	// TrdMarket_Futures 期货市场
	TrdMarket_Futures = int32(trdcommon.TrdMarket_TrdMarket_Futures)

	// TrdMarket_SG 新加坡市场
	TrdMarket_SG = int32(trdcommon.TrdMarket_TrdMarket_SG)

	// TrdMarket_Futures_Simulate_HK 香港模拟期货市场
	TrdMarket_Futures_Simulate_HK = int32(trdcommon.TrdMarket_TrdMarket_Futures_Simulate_HK)

	// TrdMarket_Futures_Simulate_US 美国模拟期货市场
	TrdMarket_Futures_Simulate_US = int32(trdcommon.TrdMarket_TrdMarket_Futures_Simulate_US)

	// TrdMarket_Futures_Simulate_SG 新加坡模拟期货市场
	TrdMarket_Futures_Simulate_SG = int32(trdcommon.TrdMarket_TrdMarket_Futures_Simulate_SG)

	// TrdMarket_Futures_Simulate_JP 日本模拟期货市场
	TrdMarket_Futures_Simulate_JP = int32(trdcommon.TrdMarket_TrdMarket_Futures_Simulate_JP)

	// TrdMarket_HK_Fund 香港基金市场
	TrdMarket_HK_Fund = int32(trdcommon.TrdMarket_TrdMarket_HK_Fund)

	// TrdMarket_US_Fund 美国基金市场
	TrdMarket_US_Fund = int32(trdcommon.TrdMarket_TrdMarket_US_Fund)
)

// 证券交易市场枚举，主要用于下单
const (
	// TrdSecMarket_Unknown 未知证券市场
	TrdSecMarket_Unknown = int32(trdcommon.TrdSecMarket_TrdSecMarket_Unknown)

	// TrdSecMarket_HK 香港证券市场
	TrdSecMarket_HK = int32(trdcommon.TrdSecMarket_TrdSecMarket_HK)

	// TrdSecMarket_US 美国证券市场
	TrdSecMarket_US = int32(trdcommon.TrdSecMarket_TrdSecMarket_US)

	// TrdSecMarket_SH 沪股市场
	TrdSecMarket_SH = int32(trdcommon.TrdSecMarket_TrdSecMarket_CN_SH)

	// TrdSecMarket_SZ 深股市场
	TrdSecMarket_SZ = int32(trdcommon.TrdSecMarket_TrdSecMarket_CN_SZ)

	// TrdSecMarket_SG 新加坡期货市场
	TrdSecMarket_SG = int32(trdcommon.TrdSecMarket_TrdSecMarket_SG)

	// TrdSecMarket_JP 日本期货市场
	TrdSecMarket_JP = int32(trdcommon.TrdSecMarket_TrdSecMarket_JP)
)

var trdSecMarketIDs = map[string]int32{
	"HK": TrdSecMarket_HK,
	"US": TrdSecMarket_US,
	"SH": TrdSecMarket_SH,
	"SZ": TrdSecMarket_SZ,
	"SG": TrdSecMarket_SG,
	"JP": TrdSecMarket_JP,
}

// GetTrdMarketID 根据市场名称返回交易市场ID
func GetTrdMarketID(name string) int32 {
	id, ok := trdSecMarketIDs[strings.ToUpper(name)]
	if ok {
		return id
	}

	return 0
}

// 交易类别
const (
	// TrdCategory_Unknown 未知
	TrdCategory_Unknown = int32(trdcommon.TrdCategory_TrdCategory_Unknown)

	// TrdCategory_Security 证券
	TrdCategory_Security = int32(trdcommon.TrdCategory_TrdCategory_Security)

	// TrdCategory_Future 期货
	TrdCategory_Future = int32(trdcommon.TrdCategory_TrdCategory_Future)
)

// 证券公司类型
const (
	// SecurityFirm_Unknown 未知
	SecurityFirm_Unknown = int32(trdcommon.SecurityFirm_SecurityFirm_Unknown)

	// SecurityFirm_FutuSecurities 富途证券（香港）
	SecurityFirm_FutuSecurities = int32(trdcommon.SecurityFirm_SecurityFirm_FutuSecurities)

	// SecurityFirm_FutuInc 富途证券（美国）
	SecurityFirm_FutuInc = int32(trdcommon.SecurityFirm_SecurityFirm_FutuInc)

	// SecurityFirm_FutuSG 富途证券（新加坡）
	SecurityFirm_FutuSG = int32(trdcommon.SecurityFirm_SecurityFirm_FutuSG)

	// SecurityFirm_FutuAU 富途证券（澳洲）
	SecurityFirm_FutuAU = int32(trdcommon.SecurityFirm_SecurityFirm_FutuAU)
)

// 持仓方向
const (
	// 未知方向
	PositionSide_Unknown = int32(trdcommon.PositionSide_PositionSide_Unknown)
	// 多仓，默认情况是多仓
	PositionSide_Long = int32(trdcommon.PositionSide_PositionSide_Long)
	// 空仓
	PositionSide_Short = int32(trdcommon.PositionSide_PositionSide_Short)
)

// 订单类型
const (
	// OrderType_Unknown 未知订单类型
	OrderType_Unknown = int32(trdcommon.OrderType_OrderType_Unknown)

	// OrderType_Normal 普通订单
	OrderType_Normal = int32(trdcommon.OrderType_OrderType_Normal)

	// OrderType_Market 市价订单
	OrderType_Market = int32(trdcommon.OrderType_OrderType_Market)

	// OrderType_AbsoluteLimit 绝对限价订单（港股）
	OrderType_AbsoluteLimit = int32(trdcommon.OrderType_OrderType_AbsoluteLimit)

	// OrderType_Auction 竞价订单（港股）
	OrderType_Auction = int32(trdcommon.OrderType_OrderType_Auction)

	// OrderType_AuctionLimit 竞价限价订单（港股）
	OrderType_AuctionLimit = int32(trdcommon.OrderType_OrderType_AuctionLimit)

	// OrderType_SpecialLimit 特殊限价订单（港股）
	OrderType_SpecialLimit = int32(trdcommon.OrderType_OrderType_SpecialLimit)

	// OrderType_SpecialLimit_All 特别限价且要求全部成交订单（港股）
	OrderType_SpecialLimit_All = int32(trdcommon.OrderType_OrderType_SpecialLimit_All)

	// OrderType_Stop 止损市价单
	OrderType_Stop = int32(trdcommon.OrderType_OrderType_Stop)

	// OrderType_StopLimit 止损限价单
	OrderType_StopLimit = int32(trdcommon.OrderType_OrderType_StopLimit)

	// OrderType_MarketifTouched 触及市价单（止盈）
	OrderType_MarketifTouched = int32(trdcommon.OrderType_OrderType_MarketifTouched)

	// OrderType_LimitifTouched 触及限价单（止盈）
	OrderType_LimitifTouched = int32(trdcommon.OrderType_OrderType_LimitifTouched)

	// OrderType_TrailingStop 跟踪止损市价单
	OrderType_TrailingStop = int32(trdcommon.OrderType_OrderType_TrailingStop)

	// OrderType_TrailingStopLimit 跟踪止损限价单
	OrderType_TrailingStopLimit = int32(trdcommon.OrderType_OrderType_TrailingStopLimit)

	// OrderType_TWAP 时间加权平均价格订单
	OrderType_TWAP = int32(trdcommon.OrderType_OrderType_TWAP)

	// OrderType_TWAPLimit 时间加权平均价格限价订单
	OrderType_TWAPLimit = int32(trdcommon.OrderType_OrderType_TWAPLimit)

	// OrderType_VWAP 成交量加权平均价格订单
	OrderType_VWAP = int32(trdcommon.OrderType_OrderType_VWAP)

	// OrderType_VWAPLimit 成交量加权平均价格限价订单
	OrderType_VWAPLimit = int32(trdcommon.OrderType_OrderType_VWAPLimit)
)

// 订单状态
const (
	// OrderStatus_Unsubmitted 未提交
	OrderStatus_Unsubmitted = int32(trdcommon.OrderStatus_OrderStatus_Unsubmitted)

	// OrderStatus_Unknown 未知状态
	OrderStatus_Unknown = int32(trdcommon.OrderStatus_OrderStatus_Unknown)

	// OrderStatus_WaitingSubmit 等待提交
	OrderStatus_WaitingSubmit = int32(trdcommon.OrderStatus_OrderStatus_WaitingSubmit)

	// OrderStatus_Submitting 提交中
	OrderStatus_Submitting = int32(trdcommon.OrderStatus_OrderStatus_Submitting)

	// OrderStatus_SubmitFailed 提交失败
	OrderStatus_SubmitFailed = int32(trdcommon.OrderStatus_OrderStatus_SubmitFailed)

	// OrderStatus_TimeOut 处理超时，结果未知
	OrderStatus_TimeOut = int32(trdcommon.OrderStatus_OrderStatus_TimeOut)

	// OrderStatus_Submitted 已提交，等待成交
	OrderStatus_Submitted = int32(trdcommon.OrderStatus_OrderStatus_Submitted)

	// OrderStatus_Filled_Part 部分成交
	OrderStatus_Filled_Part = int32(trdcommon.OrderStatus_OrderStatus_Filled_Part)

	// OrderStatus_Filled_All 全部成交
	OrderStatus_Filled_All = int32(trdcommon.OrderStatus_OrderStatus_Filled_All)

	// OrderStatus_Cancelling_Part 正在撤单_部分（部分已成交，正在撤销剩余部分）
	OrderStatus_Cancelling_Part = int32(trdcommon.OrderStatus_OrderStatus_Cancelling_Part)

	// OrderStatus_Cancelling_All 正在撤单_全部
	OrderStatus_Cancelling_All = int32(trdcommon.OrderStatus_OrderStatus_Cancelling_All)

	// OrderStatus_Cancelled_Part 部分成交，剩余部分已撤单
	OrderStatus_Cancelled_Part = int32(trdcommon.OrderStatus_OrderStatus_Cancelled_Part)

	// OrderStatus_Cancelled_All 全部已撤单，无成交
	OrderStatus_Cancelled_All = int32(trdcommon.OrderStatus_OrderStatus_Cancelled_All)

	// OrderStatus_Failed 下单失败，服务拒绝
	OrderStatus_Failed = int32(trdcommon.OrderStatus_OrderStatus_Failed)

	// OrderStatus_Disabled 已失效
	OrderStatus_Disabled = int32(trdcommon.OrderStatus_OrderStatus_Disabled)

	// OrderStatus_Deleted 已删除，无成交的订单才能删除
	OrderStatus_Deleted = int32(trdcommon.OrderStatus_OrderStatus_Deleted)

	// OrderStatus_FillCancelled 成交被撤销，一般遇不到，意思是已经成交的订单被回滚撤销，成交无效变为废单
	OrderStatus_FillCancelled = int32(trdcommon.OrderStatus_OrderStatus_FillCancelled)
)

// 交易方向
const (
	// TrdSide_Unknown 未知
	TrdSide_Unknown = int32(trdcommon.TrdSide_TrdSide_Unknown)

	// TrdSide_Buy 买入
	TrdSide_Buy = int32(trdcommon.TrdSide_TrdSide_Buy)

	// TrdSide_Sell 卖出
	TrdSide_Sell = int32(trdcommon.TrdSide_TrdSide_Sell)
)
