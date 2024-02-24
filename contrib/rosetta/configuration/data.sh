#!/bin/sh

set -e

wait_sedad() {
  timeout 30 sh -c 'until nc -z $0 $1; do sleep 1; done' localhost 9090
}
# this script is used to recreate the data dir
echo clearing /root/.sedaapp
rm -rf /root/.sedaapp
echo initting new chain
# init config files
sedad init sedad --chain-id testing

# create accounts
sedad keys add fd --keyring-backend=test

addr=$(sedad keys show fd -a --keyring-backend=test)
val_addr=$(sedad keys show fd  --keyring-backend=test --bech val -a)

# give the accounts some money
sedad add-genesis-account "$addr" 1000000000000stake --keyring-backend=test

# save configs for the daemon
sedad gentx fd 10000000stake --chain-id testing --keyring-backend=test

# input genTx to the genesis file
sedad collect-gentxs
# verify genesis file is fine
sedad validate-genesis
echo changing network settings
sed -i 's/127.0.0.1/0.0.0.0/g' /root/.sedaapp/config/config.toml

# start sedad
echo starting sedad...
sedad start --pruning=nothing &
pid=$!
echo sedad started with PID $pid

echo awaiting for sedad to be ready
wait_sedad
echo sedad is ready
sleep 10


# send transaction to deterministic address
echo sending transaction with addr $addr
sedad tx bank send "$addr" cosmos19g9cm8ymzchq2qkcdv3zgqtwayj9asv3hjv5u5 100stake --yes --keyring-backend=test --broadcast-mode=block --chain-id=testing

sleep 10

echo stopping sedad...
kill -9 $pid

echo zipping data dir and saving to /tmp/data.tar.gz

tar -czvf /tmp/data.tar.gz /root/.sedaapp

echo new address for bootstrap.json "$addr" "$val_addr"
