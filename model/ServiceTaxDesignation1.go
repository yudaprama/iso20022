package model

// Designates the tax calculation to be applied on a service.
type ServiceTaxDesignation1 struct {

	// Identifies the taxable status of the service.
	Code *ServiceTaxDesignation1Code `xml:"Cd"`

	// Defines the tax region associated with the service. This element must be present if taxes are involved with any portion of the statement.
	Region *Max35Text `xml:"Rgn,omitempty"`

	// Provides free form explanations of the various tax codes used within the statement.
	TaxReason []*TaxReason1 `xml:"TaxRsn,omitempty"`
}

func (s *ServiceTaxDesignation1) SetCode(value string) {
	s.Code = (*ServiceTaxDesignation1Code)(&value)
}

func (s *ServiceTaxDesignation1) SetRegion(value string) {
	s.Region = (*Max35Text)(&value)
}

func (s *ServiceTaxDesignation1) AddTaxReason() *TaxReason1 {
	newValue := new(TaxReason1)
	s.TaxReason = append(s.TaxReason, newValue)
	return newValue
}
