package main

type TXInput struct {
	Txid      []byte
	Vout      int
	Signature []byte
	Pubkey    []byte
}
