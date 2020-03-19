# Tron Transfer Gateway Development Documentation

## General Knowlege
* There are two address representation in tron base58 (like bitcoin) and hex string (like in Ethereum)
* Whatever happened inside contract is hex string base
* Tron has personal sign
* The way Tron transfer gateway works is the same as Ethereum trasnfer gateway
* ERC20 is called TRC20
* TRX has 6 decimal points
* `tronWeb` is `web3` equivalent

## To run end to end testing script
* `yarn run test tron-test.js`


## Installing Tron-quickstart (Ganache equivalent)
* docker must be installed
* `docker pull trontools/quickstart`
* `docker run -it --rm -p 9090:9090 -p 50051:50051  --name tron -e “defaultBalance=100000” -e “showQueryString=true” -e “showBody=true” -e “formatJson=true” -e "mnemonic=position become taxi suggest ivory rack strike water aunt tobacco benefit crane"  trontools/quickstart`

## Installing Tronbox (Truffle equivalent)
`npm install -g tronbox`

## To complie, migrate, test
* `tronbox compile`
* `tronbox migrate`
* `tronbox test`

## Useful links
* get event from a contract
    * https://api.shasta.trongrid.io/event/contract/TRD8tyaXT4nvX4nH3Z4qnkgA3rWaShmB7c
* get transaction info
    * https://shasta.tronscan.org/#/transaction/9122b94f8b58c6faa409b967ced030e4686ff2f5844b106d4b7863a6e986a9a8
* address conversion tool
    * https://tronscan.org/#/tools/tron-convert-tool

