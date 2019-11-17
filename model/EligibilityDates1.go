package model

// Dates determining the entitlement.
type EligibilityDates1 struct {

	// Date at which the positions are struck to note which parties will receive the entitlement, e.g. record date, book close date...
	EntitlementFixingDate *ISODate `xml:"EntitlmntFxgDt"`
}

func (e *EligibilityDates1) SetEntitlementFixingDate(value string) {
	e.EntitlementFixingDate = (*ISODate)(&value)
}
