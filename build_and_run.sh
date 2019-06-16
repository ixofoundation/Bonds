make install
rm -rf ~/.ns*

nsd init myvalidator --chain-id namechain

printf "12345678\n12345678\n" | nscli keys add miguel
printf "12345678\n12345678\n" | nscli keys add francesco
printf "12345678\n12345678\n" | nscli keys add shaun
printf "12345678\n12345678\n" | nscli keys add reserve

nsd add-genesis-account $(nscli keys show miguel -a) 1000000reservetoken,100000000stake
nsd add-genesis-account $(nscli keys show francesco -a) 1000000reservetoken,100000000stake
nsd add-genesis-account $(nscli keys show shaun -a) 1000000reservetoken,100000000stake

nscli config chain-id namechain
nscli config output json
nscli config indent true
nscli config trust-node true

printf "12345678\n12345678\n" | nsd gentx --name miguel

nsd collect-gentxs
nsd validate-genesis