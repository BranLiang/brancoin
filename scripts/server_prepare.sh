echo "Reset"
rm *.db && rm *.dat

echo "1. Set the seed NODE_ID=3000"
export NODE_ID=3000

echo "2. Create wallet"
addr=`brancoin createWallet`

echo "3. Create blockchain"
brancoin createBlockchain $addr

echo "4. Backup the blockchain to blockchain_genesis.db"
cp blockchain_3000.db blockchain_genesis.db
brancoin listAddresses

echo "5. Set the wallet NODE_ID=3001"
export NODE_ID=3001

echo "6. Create wallets"
brancoin createWallet
brancoin createWallet

echo "7. Copy genesis blockchain"
cp blockchain_genesis.db blockchain_3001.db

echo "8. Set the miner node NODE_ID=3002"
export NODE_ID=3002

echo "9. Create miner wallet"
brancoin createWallet

echo "10. Copy genesis blockchain"
cp blockchain_genesis.db blockchain_3002.db
