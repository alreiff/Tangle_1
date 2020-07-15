package main


import (
    . "github.com/iotaledger/iota.go/api"
    "github.com/iotaledger/iota.go/bundle"
    "github.com/iotaledger/iota.go/converter"
    "github.com/iotaledger/iota.go/trinary"
    "fmt"
)
func main() {
	var node = "https://nodes.devnet.thetangle.org"
	api, err := ComposeAPI(HTTPClientSettings{URI: node})
	must(err)
	const depth = 3;
	const minimumWeightMagnitude = 9;
	const address = trinary.Trytes("ZLGVEQ9JUZZWCZXLWVNTHBDX9G9KZTJP9VEERIIFHY9SIQKYBVAHIMLHXPQVE9IXFDDXNHQINXJDRPFDXNYVAPLZAW")
	const seed = trinary.Trytes("JBN9ZRCOH9YRUGSWIQNZWAIFEZUBDUGTFPVRKXWPAUCEQQFS9NHPQLXCKZKRHVCCUZNF9CZZWKXRZVCWQ")
	var data = "{'message' : 'Hello world'}"
	message, err := converter.ASCIIToTrytes(data)
	must(err)
	transfers := bundle.Transfers{
	    {
		Address: address,
		Value: 0,
		Message: message,
	    },
	}
	trytes, err := api.PrepareTransfers(seed, transfers, PrepareTransfersOptions{})
	must(err)

	myBundle, err := api.SendTrytes(trytes, depth, minimumWeightMagnitude)
	must(err)

	fmt.Println(bundle.TailTransactionHash(myBundle))
}
