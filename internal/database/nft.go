package database

type NFT struct {
	name       string
	properties properties
}

type properties struct {
	power int
}
