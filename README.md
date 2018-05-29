# Blockchain in Go

A blockchain implementation in Go, as described in repository: [blockchain-go](https://github.com/Jeiwan/blockchain_go). Additionally, more featured comments and tests are added for better understanding.

## Getting start

Set the env NODE_ID for the terminal tab
```shell
$ export NODE_ID=3000
```

Create a new wallet
```shell
$ brancoin createWallet
$ Your new address: 1MNPDSUr3YPNt35hCmpXbwxcukPd9Chjjy
```

List all addresses
```shell
$ brancoin listAddresses
1MNPDSUr3YPNt35hCmpXbwxcukPd9Chjjy
```

Print all wallets
```shell
$ brancoin printWallets
--- Wallet 1MNPDSUr3YPNt35hCmpXbwxcukPd9Chjjy:
       PrivateKey:  accb4c9d7d4e883051e4c82498a60307786fe92872382d4e651c40734f37dad1
       PublicKey: dcecab0be55135df7e3c26dce6ea08a5dd6bcf0faedabf84fa985034e0c71ac24d0b398f615f1c543578782a92c18465cac15550cca96155e4be93ce11921f13
--- Wallet 16SUZuXC4CvTjXMMnd3JYTXFJQJp7ckMUQ:
       PrivateKey:  96faf90c54e93338db89132d8d29b8df9e757a8abf74e3bda28a61b3af6952f5
       PublicKey: 55e1bec7b984a2bdf8a8e955edc2a6bc282d1763a99e4e1fb2d889a0fdccbdd93cedb4418bea486a9ed3923994d6423589c1dcb98d4bfb0bc664296f7f8b4df3
```

Create a new blockchain
```shell
$ brancoin createBlockchain
b26975db665935e8d9bed8be7c23dac4c0b21e5278aa7115920bd10574e9c3b1

Done!
```

Print the Blockchain
```shell
$ brancoin printChain
============ Block 00003dbbc3b835daea705487fbd5e24fa878b51c4134925f0611f967859eb87a ============
Height: 0
Prev. block:
PoW: true

--- Transaction 6d344199abfc9ed9a80604a94c5277f9b1a4f47ab542cd9ee7a2e4daf0801556:
     Input 0:
       TXID:
       Out:       -1
       Signature:
       PubKey:    5468652054696d65732030332f4a616e2f32303039204368616e63656c6c6f72206f6e206272696e6b206f66207365636f6e64206261696c6f757420666f722062616e6b73
     Output 0:
       Value:  10
       Script: c0f3a5185313c61c482d3b41335a9c31bbe65de4
```

Make transactions
```shell
$ brancoin send -f 1JbEcaQoX52MsacPJmFZcs5xkNfrFQrqax -t 18PQQ2vSoFuzKbqJzLF3U5u29pAVh6VJfs -a 1
61995dcd8194235a17882f64dd34e68cf4f68f75f1f430c7c91cfc1931a9671e

Success!
```
