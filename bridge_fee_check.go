package poly_bridge_sdk

import (
	"fmt"
)

type BridgeFeeCheck struct {
	sdk_pro []*BridgeSdkPro
}

func NewBridgeFeeCheck(sdks [][]string, slot uint64) *BridgeFeeCheck {
	feeCheck := &BridgeFeeCheck{sdk_pro:make([]*BridgeSdkPro, 0)}
	for _, sdk := range sdks {
		sdkPro := NewBridgeSdkPro(sdk, slot)
		feeCheck.sdk_pro = append(feeCheck.sdk_pro, sdkPro)
	}
	return feeCheck
}

func (check *BridgeFeeCheck) CheckFee(checks []*CheckFeeReq) ([]*CheckFeeRsp, error) {
	id2FeeRsp := make(map[string]*CheckFeeRsp)
	rsps := make([]*CheckFeeRsp, 0)
	for _, sdk := range check.sdk_pro {
		rsp, err := sdk.CheckFee(checks)
		if err != nil {
			continue
		}
		for _, newItem := range rsp {
			oldItem, ok := id2FeeRsp[fmt.Sprintf("%d%s", newItem.ChainId, newItem.Hash)]
			if !ok {
				rsps = append(rsps, newItem)
				id2FeeRsp[fmt.Sprintf("%d%s", newItem.ChainId, newItem.Hash)] = newItem
			} else {
				if oldItem.PayState == STATE_NOTCHECK {
					oldItem.PayState = newItem.PayState
					oldItem.MinProxyFee = newItem.MinProxyFee
					oldItem.Amount = newItem.Amount
				}
			}
		}
	}
	return rsps, nil
}
