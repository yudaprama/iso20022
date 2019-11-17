package model

// Specifies the detailed parameters a service to be billed.
type BillingServiceParameters3 struct {

	// Specifies the details to fully identify the bank service.
	BankService *BillingServiceIdentification3 `xml:"BkSvc"`

	// Count or number of items (volume) involved in the charge.
	Volume *DecimalNumber `xml:"Vol,omitempty"`
}

func (b *BillingServiceParameters3) AddBankService() *BillingServiceIdentification3 {
	b.BankService = new(BillingServiceIdentification3)
	return b.BankService
}

func (b *BillingServiceParameters3) SetVolume(value string) {
	b.Volume = (*DecimalNumber)(&value)
}
