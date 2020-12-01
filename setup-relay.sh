#!/bin/sh

if [[ -z "$1" ]]
then
   echo "ganacheAccountPrivateKey is empty";
   exit 1
fi

re='^0x.+$'
if ! [[ $1 =~ $re ]]
then
   echo "ganacheAccountPrivateKey is not starting with '0x'";
   exit 1
fi

if [[ -z "$2" ]]
then
   echo "genesisBlock is empty";
   exit 1
fi

re='^[0-9]+$'
if ! [[ $2 =~ $re ]] ; then
   echo "genesisBlock not a number"
   exit 1
fi

if [ -z "$3" ]
then
   echo "stakeInETH is empty";
fi

if ! [[ $3 =~ $re ]] ; then
   echo "stakeInETH not a number"
   exit 1
fi

ganacheAccountPrivateKey=$1
genesisBlock=$2
stakeInETH=$3
stakeInWei="${stakeInETH}000000000000000000"
epoch=$(($genesisBlock / 30000))

printf "${ganacheAccountPrivateKey}\nY\n" | go run main.go init

go run main.go deploy ethash

go run main.go submit epoch $epoch

go run main.go deploy testimonium --genesis $genesisBlock

go run main.go stake deposit $stakeInWei

go run main.go submit block -l
