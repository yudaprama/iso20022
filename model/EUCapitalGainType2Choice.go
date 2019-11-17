package model

// Choice of an EU capital gain type.
type EUCapitalGainType2Choice struct {

	// Specifies whether capital gain is in the scope of the European directive on taxation of savings income in the form of interest payments (Council Directive 2003/48/EC 3 June), or an income realised upon sale, a refund or redemption of shares and units, etc.
	EUCapitalGain *EUCapitalGain2Code `xml:"EUCptlGn"`

	// Specifies whether capital gain is in the scope of the European directive on taxation of savings income in the form of interest payments (Council Directive 2003/48/EC 3 June), or an income realised upon sale, a refund or redemption of shares and units, etc.
	Proprietary *GenericIdentification38 `xml:"Prtry"`
}

func (e *EUCapitalGainType2Choice) SetEUCapitalGain(value string) {
	e.EUCapitalGain = (*EUCapitalGain2Code)(&value)
}

func (e *EUCapitalGainType2Choice) AddProprietary() *GenericIdentification38 {
	e.Proprietary = new(GenericIdentification38)
	return e.Proprietary
}
