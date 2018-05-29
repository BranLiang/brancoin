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
