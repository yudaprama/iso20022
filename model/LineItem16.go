package model

// Unit of information  showing the related  provision of products and/or services and monetary summations reported as a discrete line items.
type LineItem16 struct {

	// The unique identification of this invoice line item.
	Identification *Max35Text `xml:"Id,omitempty"`

	// Something that is produced and sold as the result of an industrial process.
	TradeProduct *TradeProduct2 `xml:"TradPdct,omitempty"`

	// Purchase order reference assigned by the buyer related to the provision of products and/or services for this line item.
	BuyerOrderIdentification *DocumentIdentification23 `xml:"BuyrOrdrId,omitempty"`

	// Contract reference related to the provision of products and/or services for this line item.
	ContractIdentification *DocumentIdentification22 `xml:"CtrctId,omitempty"`

	// Specific purchase account for recording debits and credits for accounting purposes.
	PurchaseAccountingAccount []*AccountingAccount1 `xml:"PurchsAcctgAcct,omitempty"`

	// Value of the price, eg, as a currency and value.
	NetPrice []*CurrencyAndAmount `xml:"NetPric,omitempty"`

	// Quantity and conversion factor on which the net price is based for this line item product and/or service.
	NetPriceQuantity *Quantity9 `xml:"NetPricQty,omitempty"`

	// Allowance or charge applied to the net price.
	NetPriceAllowanceCharge []*LineItemAllowanceCharge2 `xml:"NetPricAllwncChrg,omitempty"`

	// Net weight of the product.
	NetWeight *Quantity10 `xml:"NetWght,omitempty"`

	// Gross price of the product and/or service.
	GrossPrice []*CurrencyAndAmount `xml:"GrssPric,omitempty"`

	// Quantity and conversion factor on which the gross price is based for this line item product and/or service.
	GrossPriceQuantity *Quantity9 `xml:"GrssPricQty,omitempty"`

	// Gross weight of the product.
	GrossWeight *Quantity10 `xml:"GrssWght,omitempty"`

	// Logistics service charge for this line item.
	LogisticsCharge []*ChargesDetails4 `xml:"LogstcsChrg,omitempty"`

	// Amount of money due to the government or tax authority, according to various pre-defined parameters such as thresholds or income.
	Tax []*LineItemTax1 `xml:"Tax,omitempty"`

	// Allowance or charge specified for this line item.
	AllowanceCharge []*LineItemAllowanceCharge2 `xml:"AllwncChrg,omitempty"`

	// Modification on the value of goods and / or services. For example: rebate, discount, surcharge
	FinancialAdjustment []*Adjustment6 `xml:"FinAdjstmnt,omitempty"`

	// Quantity billed for this line item.
	BilledQuantity *Quantity10 `xml:"BlldQty,omitempty"`

	// Number of product packages delivered.
	PackageQuantity *DecimalNumber `xml:"PackgQty,omitempty"`

	// Number of units per package in this line item for a supply chain trade delivery.
	PerPackageUnitQuantity *Quantity10 `xml:"PerPackgUnitQty,omitempty"`

	// Physical packaging of the product.
	Packaging []*Packaging1 `xml:"Packgng,omitempty"`

	// Quantity that is free of charge for this line item.
	ChargeFreeQuantity *Quantity10 `xml:"ChrgFreeQty,omitempty"`

	// Quantity value on which the quantity measurement started for a line item. For instance the start amount of a meter reading for an electricity supplier.
	MeasureQuantityStart *Quantity10 `xml:"MeasrQtyStart,omitempty"`

	// Quantity value on which the quantity measurement ended for a line item. For instance the end amount of a meter reading for an electricity supplier.
	MeasureQuantityEnd *Quantity10 `xml:"MeasrQtyEnd,omitempty"`

	// Date/time on which the clock time measure started for a line item.
	MeasureDateTimeStart *ISODateTime `xml:"MeasrDtTmStart,omitempty"`

	// Date/time on which the clock time measure ended for a line item.
	MeasureDateTimeEnd *ISODateTime `xml:"MeasrDtTmEnd,omitempty"`

	// Party to whom the goods must be delivered in the end.
	ShipTo *TradeParty3 `xml:"ShipTo,omitempty"`

	// Specifies the applicable Incoterm and associated location.
	Incoterms *Incoterms3 `xml:"Incotrms,omitempty"`

	// Actual delivery date/time of the products and/or services for this line item.
	DeliveryDateTime *ISODateTime `xml:"DlvryDtTm,omitempty"`

	// Delivery note related to the delivery of the products and/or services for this line item.
	DeliveryNoteIdentification *DocumentIdentification22 `xml:"DlvryNoteId,omitempty"`

	// Monetary totals for this line item.
	MonetarySummation *LineItemMonetarySummation1 `xml:"MntrySummtn,omitempty"`

	// Note included in this line item.
	IncludedNote []*AdditionalInformation1 `xml:"InclNote,omitempty"`
}

func (l *LineItem16) SetIdentification(value string) {
	l.Identification = (*Max35Text)(&value)
}

func (l *LineItem16) AddTradeProduct() *TradeProduct2 {
	l.TradeProduct = new(TradeProduct2)
	return l.TradeProduct
}

func (l *LineItem16) AddBuyerOrderIdentification() *DocumentIdentification23 {
	l.BuyerOrderIdentification = new(DocumentIdentification23)
	return l.BuyerOrderIdentification
}

func (l *LineItem16) AddContractIdentification() *DocumentIdentification22 {
	l.ContractIdentification = new(DocumentIdentification22)
	return l.ContractIdentification
}

func (l *LineItem16) AddPurchaseAccountingAccount() *AccountingAccount1 {
	newValue := new(AccountingAccount1)
	l.PurchaseAccountingAccount = append(l.PurchaseAccountingAccount, newValue)
	return newValue
}

func (l *LineItem16) AddNetPrice(value, currency string) {
	l.NetPrice = append(l.NetPrice, NewCurrencyAndAmount(value, currency))
}

func (l *LineItem16) AddNetPriceQuantity() *Quantity9 {
	l.NetPriceQuantity = new(Quantity9)
	return l.NetPriceQuantity
}

func (l *LineItem16) AddNetPriceAllowanceCharge() *LineItemAllowanceCharge2 {
	newValue := new(LineItemAllowanceCharge2)
	l.NetPriceAllowanceCharge = append(l.NetPriceAllowanceCharge, newValue)
	return newValue
}

func (l *LineItem16) AddNetWeight() *Quantity10 {
	l.NetWeight = new(Quantity10)
	return l.NetWeight
}

func (l *LineItem16) AddGrossPrice(value, currency string) {
	l.GrossPrice = append(l.GrossPrice, NewCurrencyAndAmount(value, currency))
}

func (l *LineItem16) AddGrossPriceQuantity() *Quantity9 {
	l.GrossPriceQuantity = new(Quantity9)
	return l.GrossPriceQuantity
}

func (l *LineItem16) AddGrossWeight() *Quantity10 {
	l.GrossWeight = new(Quantity10)
	return l.GrossWeight
}

func (l *LineItem16) AddLogisticsCharge() *ChargesDetails4 {
	newValue := new(ChargesDetails4)
	l.LogisticsCharge = append(l.LogisticsCharge, newValue)
	return newValue
}

func (l *LineItem16) AddTax() *LineItemTax1 {
	newValue := new(LineItemTax1)
	l.Tax = append(l.Tax, newValue)
	return newValue
}

func (l *LineItem16) AddAllowanceCharge() *LineItemAllowanceCharge2 {
	newValue := new(LineItemAllowanceCharge2)
	l.AllowanceCharge = append(l.AllowanceCharge, newValue)
	return newValue
}

func (l *LineItem16) AddFinancialAdjustment() *Adjustment6 {
	newValue := new(Adjustment6)
	l.FinancialAdjustment = append(l.FinancialAdjustment, newValue)
	return newValue
}

func (l *LineItem16) AddBilledQuantity() *Quantity10 {
	l.BilledQuantity = new(Quantity10)
	return l.BilledQuantity
}

func (l *LineItem16) SetPackageQuantity(value string) {
	l.PackageQuantity = (*DecimalNumber)(&value)
}

func (l *LineItem16) AddPerPackageUnitQuantity() *Quantity10 {
	l.PerPackageUnitQuantity = new(Quantity10)
	return l.PerPackageUnitQuantity
}

func (l *LineItem16) AddPackaging() *Packaging1 {
	newValue := new(Packaging1)
	l.Packaging = append(l.Packaging, newValue)
	return newValue
}

func (l *LineItem16) AddChargeFreeQuantity() *Quantity10 {
	l.ChargeFreeQuantity = new(Quantity10)
	return l.ChargeFreeQuantity
}

func (l *LineItem16) AddMeasureQuantityStart() *Quantity10 {
	l.MeasureQuantityStart = new(Quantity10)
	return l.MeasureQuantityStart
}

func (l *LineItem16) AddMeasureQuantityEnd() *Quantity10 {
	l.MeasureQuantityEnd = new(Quantity10)
	return l.MeasureQuantityEnd
}

func (l *LineItem16) SetMeasureDateTimeStart(value string) {
	l.MeasureDateTimeStart = (*ISODateTime)(&value)
}

func (l *LineItem16) SetMeasureDateTimeEnd(value string) {
	l.MeasureDateTimeEnd = (*ISODateTime)(&value)
}

func (l *LineItem16) AddShipTo() *TradeParty3 {
	l.ShipTo = new(TradeParty3)
	return l.ShipTo
}

func (l *LineItem16) AddIncoterms() *Incoterms3 {
	l.Incoterms = new(Incoterms3)
	return l.Incoterms
}

func (l *LineItem16) SetDeliveryDateTime(value string) {
	l.DeliveryDateTime = (*ISODateTime)(&value)
}

func (l *LineItem16) AddDeliveryNoteIdentification() *DocumentIdentification22 {
	l.DeliveryNoteIdentification = new(DocumentIdentification22)
	return l.DeliveryNoteIdentification
}

func (l *LineItem16) AddMonetarySummation() *LineItemMonetarySummation1 {
	l.MonetarySummation = new(LineItemMonetarySummation1)
	return l.MonetarySummation
}

func (l *LineItem16) AddIncludedNote() *AdditionalInformation1 {
	newValue := new(AdditionalInformation1)
	l.IncludedNote = append(l.IncludedNote, newValue)
	return newValue
}
