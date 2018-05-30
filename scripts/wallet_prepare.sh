echo "1. Set the NODE_ID"
export NODE_ID=3001

echo "2. Create wallets"
brancoin createWallet
brancoin createWallet

echo "3. Copy genesis blockchain"
cp blockchain_genesis.db blockchain_3001.db
