package adapt

import "github.com/hyperjiang/futu/pb/trdcommon"

// 交易市场
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

const (
	// 未知方向
	PositionSide_Unknown = int32(trdcommon.PositionSide_PositionSide_Unknown)
	// 多仓，默认情况是多仓
	PositionSide_Long = int32(trdcommon.PositionSide_PositionSide_Long)
	// 空仓
	PositionSide_Short = int32(trdcommon.PositionSide_PositionSide_Short)
)
