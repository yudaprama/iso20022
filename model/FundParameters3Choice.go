package model

// Choice of fund parameters.
type FundParameters3Choice struct {

	// Specifies that there is no criteria for the report. The request is a request for all reports, rather than reports attributed to a specific fund manager, date or financial instrument.
	NoCriteria *NoCriteria1Code `xml:"NoCrit"`

	// Report parameters.
	Parameters *FundParameters4 `xml:"Params"`
}

func (f *FundParameters3Choice) SetNoCriteria(value string) {
	f.NoCriteria = (*NoCriteria1Code)(&value)
}

func (f *FundParameters3Choice) AddParameters() *FundParameters4 {
	f.Parameters = new(FundParameters4)
	return f.Parameters
}
