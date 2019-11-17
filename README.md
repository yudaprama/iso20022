# ISO20022

This repository contains a full set of Golang parse to the ISO-20022 data catalogs [ISO-20022 specifications](https://www.iso20022.org/full_catalogue.page)

## Usage

See `examples/` directory for an example of usage

```go
import (
	"encoding/xml"
	"github.com/yudaprama/iso20022/pacs"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	messages, err := ioutil.ReadFile("./example-message.xml")
	if err != nil {
		log.Fatalf("Unable to read file:  %v", err)
		os.Exit(1)
	}

	var messageParsed pacs.Document00800106
	err = xml.Unmarshal(messages, &messageParsed)
	if err != nil {
		log.Fatalf("Unable to parse file:  %v", err)
		os.Exit(1)
	}
	
	log.Printf("Interbank Settlement Date:  %v", messageParsed.Message.GroupHeader.InterbankSettlementDate)
}

```

## Message Types

Message types covers ISO-20022 messages:

* [ACMT - Account Management](https://github.com/yudaprama/iso20022/tree/master/acmt)
* [ADMI - Administration](https://github.com/yudaprama/iso20022/tree/master/admi)
* [AUTH - Authorities](https://github.com/yudaprama/iso20022/tree/master/auth)
* [CAAA - Acceptor to Acquirer Card Transactions](https://github.com/yudaprama/iso20022/tree/master/caaa)
* [CAAM - ATM Management](https://github.com/yudaprama/iso20022/tree/master/caam)
* [CAIN - Acquirer to Issuer Card Transactions](https://github.com/yudaprama/iso20022/tree/master/cain)
* [CAMT - Cash Management](https://github.com/yudaprama/iso20022/tree/master/camt)
* [CATM - Terminal Management](https://github.com/yudaprama/iso20022/tree/master/catm)
* [CATP - ATM Card Transactions](https://github.com/yudaprama/iso20022/tree/master/catp)
* [COLR - Collateral Management](https://github.com/yudaprama/iso20022/tree/master/colr)
* [FXTR - Foreign Exchange Trade](https://github.com/yudaprama/iso20022/tree/master/fxtr)
* [HEAD - Business Application Header](https://github.com/yudaprama/iso20022/tree/master/head)
* [PACS - Payments Clearing and Settlement](https://github.com/yudaprama/iso20022/tree/master/pacs)
* [PAIN - Payments Initiation](https://github.com/yudaprama/iso20022/tree/master/pain)
* [REDA - Reference Data](https://github.com/yudaprama/iso20022/tree/master/reda)
* [REMT - Payments Remittance Advice](https://github.com/yudaprama/iso20022/tree/master/remt)
* [SECL - Securities Clearing](https://github.com/yudaprama/iso20022/tree/master/secl)
* [SEEV - Securities Events](https://github.com/yudaprama/iso20022/tree/master/seev)
* [SEMT - Securities Management](https://github.com/yudaprama/iso20022/tree/master/semt)
* [SESE - Securities Settlement](https://github.com/yudaprama/iso20022/tree/master/sese)
* [SETR - Securities Trade](https://github.com/yudaprama/iso20022/tree/master/setr)
* [SUPL - Supplementary Data](https://github.com/yudaprama/iso20022/tree/master/supl)
* [TREA - Treasury](https://github.com/yudaprama/iso20022/tree/master/trea)
* [TSIN - Trade Services Initiation](https://github.com/yudaprama/iso20022/tree/master/tsin)
* [TSMT - Trade Services Management](https://github.com/yudaprama/iso20022/tree/master/tsmt)
* [TSRV - Trade Services](https://github.com/yudaprama/iso20022/tree/master/tsrv)
