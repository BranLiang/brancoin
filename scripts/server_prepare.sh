echo "Reset"
rm *.db && rm *.dat

echo "1. Set the NODE_ID"
export NODE_ID=3000

echo "2. Create wallet"
addr=`brancoin createWallet`

echo "3. Create blockchain"
brancoin createBlockchain $addr

echo "4. Backup the blockchain to blockchain_genesis.db"
cp blockchain_3000.db blockchain_genesis.db
brancoin listAddresses
