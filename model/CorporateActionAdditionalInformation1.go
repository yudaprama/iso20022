package model

// Provides additional information about the delivery details, beneficial owner details, etc.
type CorporateActionAdditionalInformation1 struct {

	// Provides information about the beneficial owner of the securities.
	BeneficialOwnerDetails []*BeneficialOwner1 `xml:"BnfclOwnrDtls,omitempty"`

	// Provides information required for the registration.
	RegistrationDetails *Max350Text `xml:"RegnDtls,omitempty"`

	// Identification of the receiver of outturned resources (cash/securities) in case the resources need to be delivered outside the CSD.
	ReceiverIdentification *PartyIdentification2Choice `xml:"RcvrId,omitempty"`

	// Whether or not certification is required from the account owner.
	// Y: certification required
	// N: no certification required
	CertificationIndicator *YesNoIndicator `xml:"CertfctnInd,omitempty"`

	// Type of certification which is required.
	CertificationType *BeneficiaryCertificationType1FormatChoice `xml:"CertfctnTp,omitempty"`

	// Provides information about the delivery details of proceeds.
	DeliveryDetails []*ProceedsDelivery1 `xml:"DlvryDtls,omitempty"`

	// Provides additional details pertaining to the corporate action instruction.
	AdditionalInstruction *Max350Text `xml:"AddtlInstr,omitempty"`
}

func (c *CorporateActionAdditionalInformation1) AddBeneficialOwnerDetails() *BeneficialOwner1 {
	newValue := new(BeneficialOwner1)
	c.BeneficialOwnerDetails = append(c.BeneficialOwnerDetails, newValue)
	return newValue
}

func (c *CorporateActionAdditionalInformation1) SetRegistrationDetails(value string) {
	c.RegistrationDetails = (*Max350Text)(&value)
}

func (c *CorporateActionAdditionalInformation1) AddReceiverIdentification() *PartyIdentification2Choice {
	c.ReceiverIdentification = new(PartyIdentification2Choice)
	return c.ReceiverIdentification
}

func (c *CorporateActionAdditionalInformation1) SetCertificationIndicator(value string) {
	c.CertificationIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateActionAdditionalInformation1) AddCertificationType() *BeneficiaryCertificationType1FormatChoice {
	c.CertificationType = new(BeneficiaryCertificationType1FormatChoice)
	return c.CertificationType
}

func (c *CorporateActionAdditionalInformation1) AddDeliveryDetails() *ProceedsDelivery1 {
	newValue := new(ProceedsDelivery1)
	c.DeliveryDetails = append(c.DeliveryDetails, newValue)
	return newValue
}

func (c *CorporateActionAdditionalInformation1) SetAdditionalInstruction(value string) {
	c.AdditionalInstruction = (*Max350Text)(&value)
}
