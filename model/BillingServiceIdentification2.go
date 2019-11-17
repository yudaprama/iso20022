package model

// Identification of the service to be billed.
type BillingServiceIdentification2 struct {

	// Financial institution's own, internal service identification code, different from the common code.
	// Usage: The financial institution own code is used to uniquely identify the service within the financial institution.
	Identification *Max35Text `xml:"Id"`

	// Defines the financial institution sub-service identification if the financial institution service identification code is used for more than one service.
	SubService *BillingSubServiceIdentification1 `xml:"SubSvc,omitempty"`

	// Specifies further details to describe the financial institution service description, which is not the standard description related to the common code.
	Description *Max70Text `xml:"Desc"`
}

func (b *BillingServiceIdentification2) SetIdentification(value string) {
	b.Identification = (*Max35Text)(&value)
}

func (b *BillingServiceIdentification2) AddSubService() *BillingSubServiceIdentification1 {
	b.SubService = new(BillingSubServiceIdentification1)
	return b.SubService
}

func (b *BillingServiceIdentification2) SetDescription(value string) {
	b.Description = (*Max70Text)(&value)
}
