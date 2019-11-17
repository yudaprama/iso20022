package model

// Choice of entitlement calculation method.
type Entitlement1Choice struct {

	// Number of votes assigned to one security.
	EntitlementRatio *DecimalNumber `xml:"EntitlmntRatio"`

	// Specifies the calculation method of the number of votes assigned to one security. This element should be used when the entitlement calculation rule is complex.
	EntitlementDescription *Max35Text `xml:"EntitlmntDesc"`
}

func (e *Entitlement1Choice) SetEntitlementRatio(value string) {
	e.EntitlementRatio = (*DecimalNumber)(&value)
}

func (e *Entitlement1Choice) SetEntitlementDescription(value string) {
	e.EntitlementDescription = (*Max35Text)(&value)
}
