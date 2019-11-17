package model

// Specifies the category of the investment fund order.
type FundOrderType2 struct {

	// Specifies the category of the investment fund order.
	OrderType *FundOrderType3Code `xml:"OrdrTp"`

	// Specifies the category of the investment fund order.
	ExtendedOrderType *Extended350Code `xml:"XtndedOrdrTp"`
}

func (f *FundOrderType2) SetOrderType(value string) {
	f.OrderType = (*FundOrderType3Code)(&value)
}

func (f *FundOrderType2) SetExtendedOrderType(value string) {
	f.ExtendedOrderType = (*Extended350Code)(&value)
}
