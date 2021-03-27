package test

import (
	"fmt"
	"poly_bridge_sdk"
	"testing"
)

func TestCheckFee(t *testing.T) {
	sdk := poly_bridge_sdk.NewBridgeFeeCheck([][]string{[]string{"https://bridge.poly.network/testnet/v1/"},[]string{"http://40.115.136.96:22000/v1/"}}, 5)
	checks := make([]*poly_bridge_sdk.CheckFeeReq, 0)
	checks = append(checks, &poly_bridge_sdk.CheckFeeReq {
		ChainId: 2,
		Hash: "0000000000000000000000000000000000000000000000000000000000001c24",
	})
	checks = append(checks, &poly_bridge_sdk.CheckFeeReq {
		ChainId: 2,
		Hash: "9acc3566087e41bf1943f6bd8fdecc1207820f0c7c00ee71cb7e11ae9388626e",
	})
	checks = append(checks, &poly_bridge_sdk.CheckFeeReq {
		ChainId: 714,
		Hash: "9acc3566087e41bf1943f6bd8fdecc1207820f0c7c00ee71cb7e11ae9388626d",
	})
	feeRsps, _ := sdk.CheckFee(checks)
	for _, feeRsp := range feeRsps {
		fmt.Printf("chain id: %d, hash: %s, has pay: %d, amount: %s\n", feeRsp.ChainId, feeRsp.Hash, feeRsp.PayState, feeRsp.Amount)
	}
}
