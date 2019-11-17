package model

// Specifies the billing adjustments for a specific service.
type BillingServiceAdjustment1 struct {

	// Identifies the type of adjustment.
	Type *ServiceAdjustmentType1Code `xml:"Tp"`

	// Free-form description and clarification of the adjustment.
	Description *Max140Text `xml:"Desc"`

	// Amount of the adjustment, expressed in the settlement currency.
	//
	// Usage: If the amount would reduce charges due then the amount should be negatively signed.
	Amount *AmountAndDirection34 `xml:"Amt"`

	// Specifies whether the balance amount requires an adjustment.
	BalanceRequiredAmount *AmountAndDirection34 `xml:"BalReqrdAmt,omitempty"`

	// Date on which the situation causing the service adjustment occurred. If the date is not known then used the last day of the month in which the situation occurred or the date of the billing statement which reported the original service to which this adjustment applies.
	ErrorDate *ISODate `xml:"ErrDt,omitempty"`

	// Financial institution's own, internal service identification code, used to uniquely identify the service within the financial institution.
	AdjustmentIdentification *Max35Text `xml:"AdjstmntId,omitempty"`

	// Defines the financial institution sub-service identification if the financial institution service identification code is used for more than one service.
	SubService *BillingSubServiceIdentification1 `xml:"SubSvc,omitempty"`

	// Change in the service price, expressed in the pricing currency. A negative value indicates a price reduction.
	PriceChange *AmountAndDirection34 `xml:"PricChng,omitempty"`

	// Price that was applied to the service, prior to the change, expressed in the pricing currency.
	OriginalPrice *AmountAndDirection34 `xml:"OrgnlPric,omitempty"`

	// New, adjusted service price, expressed in the pricing currency.
	NewPrice *AmountAndDirection34 `xml:"NewPric,omitempty"`

	// Change in the service volume. A negative value indicates a volume reduction.
	VolumeChange *DecimalNumber `xml:"VolChng,omitempty"`

	// Original service volume.
	OriginalVolume *DecimalNumber `xml:"OrgnlVol,omitempty"`

	// New, adjusted service volume.
	NewVolume *DecimalNumber `xml:"NewVol,omitempty"`

	// Service charge that was applied to the service, prior to the change, expressed in the pricing currency.
	OriginalChargeAmount *AmountAndDirection34 `xml:"OrgnlChrgAmt,omitempty"`

	// New, adjusted service charge, expressed in the pricing currency.
	NewChargeAmount *AmountAndDirection34 `xml:"NewChrgAmt,omitempty"`
}

func (b *BillingServiceAdjustment1) SetType(value string) {
	b.Type = (*ServiceAdjustmentType1Code)(&value)
}

func (b *BillingServiceAdjustment1) SetDescription(value string) {
	b.Description = (*Max140Text)(&value)
}

func (b *BillingServiceAdjustment1) AddAmount() *AmountAndDirection34 {
	b.Amount = new(AmountAndDirection34)
	return b.Amount
}

func (b *BillingServiceAdjustment1) AddBalanceRequiredAmount() *AmountAndDirection34 {
	b.BalanceRequiredAmount = new(AmountAndDirection34)
	return b.BalanceRequiredAmount
}

func (b *BillingServiceAdjustment1) SetErrorDate(value string) {
	b.ErrorDate = (*ISODate)(&value)
}

func (b *BillingServiceAdjustment1) SetAdjustmentIdentification(value string) {
	b.AdjustmentIdentification = (*Max35Text)(&value)
}

func (b *BillingServiceAdjustment1) AddSubService() *BillingSubServiceIdentification1 {
	b.SubService = new(BillingSubServiceIdentification1)
	return b.SubService
}

func (b *BillingServiceAdjustment1) AddPriceChange() *AmountAndDirection34 {
	b.PriceChange = new(AmountAndDirection34)
	return b.PriceChange
}

func (b *BillingServiceAdjustment1) AddOriginalPrice() *AmountAndDirection34 {
	b.OriginalPrice = new(AmountAndDirection34)
	return b.OriginalPrice
}

func (b *BillingServiceAdjustment1) AddNewPrice() *AmountAndDirection34 {
	b.NewPrice = new(AmountAndDirection34)
	return b.NewPrice
}

func (b *BillingServiceAdjustment1) SetVolumeChange(value string) {
	b.VolumeChange = (*DecimalNumber)(&value)
}

func (b *BillingServiceAdjustment1) SetOriginalVolume(value string) {
	b.OriginalVolume = (*DecimalNumber)(&value)
}

func (b *BillingServiceAdjustment1) SetNewVolume(value string) {
	b.NewVolume = (*DecimalNumber)(&value)
}

func (b *BillingServiceAdjustment1) AddOriginalChargeAmount() *AmountAndDirection34 {
	b.OriginalChargeAmount = new(AmountAndDirection34)
	return b.OriginalChargeAmount
}

func (b *BillingServiceAdjustment1) AddNewChargeAmount() *AmountAndDirection34 {
	b.NewChargeAmount = new(AmountAndDirection34)
	return b.NewChargeAmount
}
