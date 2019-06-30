#!/usr/bin/env bash
nsd init local --chain-id pricingchain

printf "12345678\n12345678\n" | nscli keys add miguel
printf "12345678\n12345678\n" | nscli keys add francesco
printf "12345678\n12345678\n" | nscli keys add shaun
printf "12345678\n12345678\n" | nscli keys add reserve

nsd add-genesis-account $(nscli keys show miguel -a) 1000000reservetoken,100000000stake
nsd add-genesis-account $(nscli keys show francesco -a) 1000000reservetoken,100000000stake
nsd add-genesis-account $(nscli keys show shaun -a) 1000000reservetoken,100000000stake

nscli config chain-id pricingchain
nscli config output json
nscli config indent true
nscli config trust-node true

echo "12345678" | nsd gentx --name miguel

nsd collect-gentxs
nsd validate-genesis

nsd start & nscli rest-server --chain-id pricingchain --trust-node && fg