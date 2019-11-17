package model

// Identification of the original transaction.
type CardPaymentTransaction65 struct {

	// Global reference of the sale transaction for the sale system.
	SaleReferenceIdentification *Max35Text `xml:"SaleRefId,omitempty"`

	// Unique identification of the transaction assigned by the POI (Point Of Interaction).
	TransactionIdentification *TransactionIdentifier1 `xml:"TxId"`

	// Identification of the POI (Point Of Interaction) performing the transaction.
	POIIdentification *GenericIdentification32 `xml:"POIId"`

	// Link to a previous currency conversion.
	CurrencyConversion *CurrencyConversion7 `xml:"CcyConvs"`
}

func (c *CardPaymentTransaction65) SetSaleReferenceIdentification(value string) {
	c.SaleReferenceIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction65) AddTransactionIdentification() *TransactionIdentifier1 {
	c.TransactionIdentification = new(TransactionIdentifier1)
	return c.TransactionIdentification
}

func (c *CardPaymentTransaction65) AddPOIIdentification() *GenericIdentification32 {
	c.POIIdentification = new(GenericIdentification32)
	return c.POIIdentification
}

func (c *CardPaymentTransaction65) AddCurrencyConversion() *CurrencyConversion7 {
	c.CurrencyConversion = new(CurrencyConversion7)
	return c.CurrencyConversion
}
