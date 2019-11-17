package model

// Data common to all transactions of a data set.
type CommonData1 struct {

	// Data related to the environment of the transaction, common to a set of transaction.
	Environment *CardPaymentEnvironment5 `xml:"Envt,omitempty"`

	// Data related to the context of the transaction, common to a set of transaction.
	Context *CardPaymentContext3 `xml:"Cntxt,omitempty"`

	// Type of transaction being undertaken for the main service, common to a set of transaction..
	TransactionType *CardPaymentServiceType1Code `xml:"TxTp,omitempty"`

	// Service in addition to the main service.
	AdditionalService []*CardPaymentServiceType2Code `xml:"AddtlSvc,omitempty"`

	// Additional attribute of the service type.
	ServiceAttribute *CardPaymentServiceType3Code `xml:"SvcAttr,omitempty"`

	// Category code conform to ISO 18245, related to the type of services or goods the merchant provides for the transaction.
	MerchantCategoryCode *Min3Max4Text `xml:"MrchntCtgyCd,omitempty"`

	// Unique identification of the reconciliation period between the acceptor and the acquirer. This identification might be linked to the identification of the settlement for further verification by the merchant.
	ReconciliationIdentification *Max35Text `xml:"RcncltnId,omitempty"`

	// Currency associated with the set of transaction.
	Currency *CurrencyCode `xml:"Ccy,omitempty"`
}

func (c *CommonData1) AddEnvironment() *CardPaymentEnvironment5 {
	c.Environment = new(CardPaymentEnvironment5)
	return c.Environment
}

func (c *CommonData1) AddContext() *CardPaymentContext3 {
	c.Context = new(CardPaymentContext3)
	return c.Context
}

func (c *CommonData1) SetTransactionType(value string) {
	c.TransactionType = (*CardPaymentServiceType1Code)(&value)
}

func (c *CommonData1) AddAdditionalService(value string) {
	c.AdditionalService = append(c.AdditionalService, (*CardPaymentServiceType2Code)(&value))
}

func (c *CommonData1) SetServiceAttribute(value string) {
	c.ServiceAttribute = (*CardPaymentServiceType3Code)(&value)
}

func (c *CommonData1) SetMerchantCategoryCode(value string) {
	c.MerchantCategoryCode = (*Min3Max4Text)(&value)
}

func (c *CommonData1) SetReconciliationIdentification(value string) {
	c.ReconciliationIdentification = (*Max35Text)(&value)
}

func (c *CommonData1) SetCurrency(value string) {
	c.Currency = (*CurrencyCode)(&value)
}
