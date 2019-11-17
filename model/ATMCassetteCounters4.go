package model

// Counters of media inside an ATM cassette.
type ATMCassetteCounters4 struct {

	// Type of counters.
	Type *ATMCounterType1Code `xml:"Tp"`

	// Number of added media during servicing operations.
	AddedNumber *Number `xml:"AddedNb,omitempty"`

	// Number of removed media during servicing operations.
	RemovedNumber *Number `xml:"RmvdNb,omitempty"`

	// Total number of media out of the cassette.
	DispensedNumber *Number `xml:"DspnsdNb,omitempty"`

	// Total number of media deposited in the cassette.
	DepositedNumber *Number `xml:"DpstdNb,omitempty"`

	// Total number of recycled media from the cassette.
	RecycledNumber *Number `xml:"RcycldNb,omitempty"`

	// Total number of retracted media originating from the cassette.
	RetractedNumber *Number `xml:"RtrctdNb,omitempty"`

	// Total number of media from this cassette which are on the reject bin.
	RejectedNumber *Number `xml:"RjctdNb,omitempty"`

	// Total number of media presented to the customer.
	PresentedNumber *Number `xml:"PresntdNb,omitempty"`
}

func (a *ATMCassetteCounters4) SetType(value string) {
	a.Type = (*ATMCounterType1Code)(&value)
}

func (a *ATMCassetteCounters4) SetAddedNumber(value string) {
	a.AddedNumber = (*Number)(&value)
}

func (a *ATMCassetteCounters4) SetRemovedNumber(value string) {
	a.RemovedNumber = (*Number)(&value)
}

func (a *ATMCassetteCounters4) SetDispensedNumber(value string) {
	a.DispensedNumber = (*Number)(&value)
}

func (a *ATMCassetteCounters4) SetDepositedNumber(value string) {
	a.DepositedNumber = (*Number)(&value)
}

func (a *ATMCassetteCounters4) SetRecycledNumber(value string) {
	a.RecycledNumber = (*Number)(&value)
}

func (a *ATMCassetteCounters4) SetRetractedNumber(value string) {
	a.RetractedNumber = (*Number)(&value)
}

func (a *ATMCassetteCounters4) SetRejectedNumber(value string) {
	a.RejectedNumber = (*Number)(&value)
}

func (a *ATMCassetteCounters4) SetPresentedNumber(value string) {
	a.PresentedNumber = (*Number)(&value)
}
