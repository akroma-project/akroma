#!/bin/sh
rm -fr ~/.akromatests/testnet/geth/chaindata

./build/bin/geth --datadir ~/.akromatests/testnet/ --testnet --rpcport 8544 --rpc --port 30302 --nodiscover --nat extip:127.0.0.1 --mine --miner.threads 1 --miner.etherbase 0x96ecaaf5b1669d374406f7e22c217fc839ca6bd1 --miner.extradata akroma.org-miner