package model

// Specifies the reason why the transaction was generated.
type GeneratedReason6 struct {

	// Specifies the reason why the transaction was generated.
	Code *GeneratedReasons6Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *RestrictedFINXMax210Text `xml:"AddtlRsnInf,omitempty"`
}

func (g *GeneratedReason6) AddCode() *GeneratedReasons6Choice {
	g.Code = new(GeneratedReasons6Choice)
	return g.Code
}

func (g *GeneratedReason6) SetAdditionalReasonInformation(value string) {
	g.AdditionalReasonInformation = (*RestrictedFINXMax210Text)(&value)
}
