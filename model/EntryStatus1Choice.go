package model

type EntryStatus1Choice struct {
	Cd    *ExternalEntryStatus1Code `xml:"Cd,omitempty"`
	Prtry *Max35Text                `xml:"Prtry,omitempty"`
}
