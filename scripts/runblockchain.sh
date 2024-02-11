geth --datadir /home/kori/blockchain/data --networkid 1008  \
--http --http.addr 127.0.0.1 --http.vhosts "*" \
--http.api "db,net,eth,web3,personal" --http.corsdomain "*" \
--snapshot=false --mine --miner.threads 1 --allow-insecure-unlock \
console 2> /home/kori/blockchain/data/geth.log