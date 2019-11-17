package model

// Specifies the category of the investment fund order.
type FundOrderType3 struct {

	// Specifies the category of the investment fund order.
	OrderType *FundOrderType4Code `xml:"OrdrTp"`

	// Specifies the category of the investment fund order.
	ExtendedOrderType *Extended350Code `xml:"XtndedOrdrTp"`
}

func (f *FundOrderType3) SetOrderType(value string) {
	f.OrderType = (*FundOrderType4Code)(&value)
}

func (f *FundOrderType3) SetExtendedOrderType(value string) {
	f.ExtendedOrderType = (*Extended350Code)(&value)
}
