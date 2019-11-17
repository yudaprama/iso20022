package model

// Goods or services that are part of a commercial trade agreement.
type TransportedGoods1 struct {

	// Reference to the purchase order of the underlying transaction.
	PurchaseOrderReference *DocumentIdentification7 `xml:"PurchsOrdrRef"`

	// Information about the goods and/or services of a trade transaction.
	GoodsDescription *Max70Text `xml:"GoodsDesc,omitempty"`

	// Information important for the users of the message/service, which cannot be captured in any other message component/element. For example: Warehouse number.
	BuyerDefinedInformation []*UserDefinedInformation1 `xml:"BuyrDfndInf,omitempty"`

	// Information important for the users of the message/service, which cannot be captured in any other message component/element. For example: Warehouse number.
	SellerDefinedInformation []*UserDefinedInformation1 `xml:"SellrDfndInf,omitempty"`
}

func (t *TransportedGoods1) AddPurchaseOrderReference() *DocumentIdentification7 {
	t.PurchaseOrderReference = new(DocumentIdentification7)
	return t.PurchaseOrderReference
}

func (t *TransportedGoods1) SetGoodsDescription(value string) {
	t.GoodsDescription = (*Max70Text)(&value)
}

func (t *TransportedGoods1) AddBuyerDefinedInformation() *UserDefinedInformation1 {
	newValue := new(UserDefinedInformation1)
	t.BuyerDefinedInformation = append(t.BuyerDefinedInformation, newValue)
	return newValue
}

func (t *TransportedGoods1) AddSellerDefinedInformation() *UserDefinedInformation1 {
	newValue := new(UserDefinedInformation1)
	t.SellerDefinedInformation = append(t.SellerDefinedInformation, newValue)
	return newValue
}
