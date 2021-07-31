package ledger

import (
	"github.com/ericlagergren/decimal"
)

func ProtoToDecimal(proto *Number) *decimal.Big {
	var dec decimal.Big
	if proto.Value != nil {
		dec.SetString(proto.GetValue())
	}
	return &dec
}

func DecimalToProto(dec *decimal.Big, proto *Number) {
	str := dec.String()
	proto.Value = &str
}
