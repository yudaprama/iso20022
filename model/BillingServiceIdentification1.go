package model

// Identification of the service to be billed.
type BillingServiceIdentification1 struct {

	// Financial institution's own, internal service identification code, different from the common code.
	// Usage: The financial institution own code is used to uniquely identify the service within the financial institution.
	Identification *Max35Text `xml:"Id"`

	// Defines the financial institution sub-service identification if the financial institution service identification code is used for more than one service.
	SubService *BillingSubServiceIdentification1 `xml:"SubSvc,omitempty"`

	// Specifies further details to describe the financial institution service description, which is not the standard description related to the common code.
	Description *Max70Text `xml:"Desc"`

	// Standard reference code used to uniquely identify this service across financial institutions. This is not the financial institution’s internal bank service identification.
	CommonCode *BillingServiceCommonIdentification1 `xml:"CmonCd,omitempty"`

	// Type used to classify or organise different services by common characteristics.
	ServiceType *Max12Text `xml:"SvcTp,omitempty"`
}

func (b *BillingServiceIdentification1) SetIdentification(value string) {
	b.Identification = (*Max35Text)(&value)
}

func (b *BillingServiceIdentification1) AddSubService() *BillingSubServiceIdentification1 {
	b.SubService = new(BillingSubServiceIdentification1)
	return b.SubService
}

func (b *BillingServiceIdentification1) SetDescription(value string) {
	b.Description = (*Max70Text)(&value)
}

func (b *BillingServiceIdentification1) AddCommonCode() *BillingServiceCommonIdentification1 {
	b.CommonCode = new(BillingServiceCommonIdentification1)
	return b.CommonCode
}

func (b *BillingServiceIdentification1) SetServiceType(value string) {
	b.ServiceType = (*Max12Text)(&value)
}
