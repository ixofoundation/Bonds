#!/usr/bin/env bash

wait () {
    echo "Waiting for chain to start..."
    while :
    do
        RET=$(nscli status 2>&1)
        if [[ $RET == ERROR* ]]
        then
            sleep 1
        else
            echo "A few more seconds..."
            sleep 6
            break
        fi
    done
}

tx_from_m () {
    cmd=$1
    shift
    echo "12345678" | nscli tx pricing $cmd --from miguel -y --broadcast-mode block $@
}

tx_from_f () {
    cmd=$1
    shift
    echo "12345678" | nscli tx pricing $cmd --from francesco -y --broadcast-mode block $@
}

RET=$(nscli status 2>&1)
if [[ $RET == ERROR* ]]
then
    wait
fi

echo "Creating cosmic bond..."
tx_from_m create-cosmic-bond edu reservetoken $(nscli keys show reserve -a) 1000000 power_function 12 2 true
echo "Created cosmic-bond..."
nscli query pricing cosmic-bond edu

echo "Miguel buys 10edu..."
tx_from_m buy edu 10edu 1000
echo "Miguel's account..."
nscli query account $(nscli keys show miguel -a)

echo "Francesco buys 10edu..."
tx_from_f buy edu 10edu 1000
echo "Francesco's account..."
nscli query account $(nscli keys show francesco -a)

echo "Miguel sells 10edu..."
tx_from_m sell edu 10edu
echo "Miguel's account..."
nscli query account $(nscli keys show miguel -a)
echo "Miguel made a profit!"

echo "Francesco sells 10edu..."
tx_from_f sell edu 10edu
echo "Francesco's account..."
nscli query account $(nscli keys show francesco -a)
echo "Francesco made a loss!"
