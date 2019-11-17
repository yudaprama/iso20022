package model

// Specifies the detailed parameters a service to be billed.
type BillingServiceParameters1 struct {

	// Specifies the details to fully identify the bank service.
	BankService *BillingServiceIdentification1 `xml:"BkSvc"`

	// Count or number of items (volume) involved in the charge.
	Volume *DecimalNumber `xml:"Vol,omitempty"`
}

func (b *BillingServiceParameters1) AddBankService() *BillingServiceIdentification1 {
	b.BankService = new(BillingServiceIdentification1)
	return b.BankService
}

func (b *BillingServiceParameters1) SetVolume(value string) {
	b.Volume = (*DecimalNumber)(&value)
}
