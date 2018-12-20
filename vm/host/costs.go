package host

import "github.com/iost-official/go-iost/core/contract"

// var costs
var (
	Costs = map[string]contract.Cost{
		"JSCost":       contract.NewCost(0, 0, 2000),
		"PutCost":      contract.NewCost(0, 0, 150),
		"GetCost":      contract.NewCost(0, 0, 100),
		"DelCost":      contract.NewCost(0, 0, 100),
		"KeysCost":     contract.NewCost(0, 0, 100),
		"ContextCost":  contract.NewCost(0, 0, 10),
		"EventPrice":   contract.NewCost(0, 0, 1),
		"ReceiptPrice": contract.NewCost(0, 1, 0),
		"CodePrice":    contract.NewCost(0, 0, 1),
		"OpPrice":      contract.NewCost(0, 0, 1),
		"ErrPrice":     contract.NewCost(0, 0, 1),
	}
)

// EventCost return cost based on event size
func EventCost(size int) contract.Cost {
	return Costs["EventPrice"].Multiply(int64(size))
}

// ReceiptCost based on receipt size
func ReceiptCost(size int) contract.Cost {
	return Costs["ReceiptPrice"].Multiply(int64(size))
}

// CodeSavageCost cost in deploy contract based on code size
func CodeSavageCost(size int) contract.Cost {
	return Costs["CodePrice"].Multiply(int64(size))
}

// CommonErrorCost returns cost increased by stack layer
func CommonErrorCost(layer int) contract.Cost {
	return Costs["ErrPrice"].Multiply(int64(layer * 10))
}

// CommonOpCost returns cost increased by stack layer
func CommonOpCost(layer int) contract.Cost {
	return Costs["OpPrice"].Multiply(int64(layer * 10))
}

// DelayTxCost returns cost of a delay transaction.
func DelayTxCost(dataLen int, payer string) contract.Cost {
	cost := Costs["PutCost"]
	cost.Data = int64(dataLen)
	cost.DataList = []contract.DataItem{{Payer: payer, Val: int64(dataLen)}}
	return cost
}

// DelDelayTxCost returns cost of a delay transaction.
func DelDelayTxCost(dataLen int, payer string) contract.Cost {
	cost := Costs["DelCost"]
	cost.Data = -int64(dataLen)
	cost.DataList = []contract.DataItem{{Payer: payer, Val: -int64(dataLen)}}
	return cost
}
