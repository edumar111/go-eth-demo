solc --evm-version ^0.4.24 --abi EverisDAExpertise.sol -o build
abigen --bin=./build/Expertise.bin --abi=./build/Expertise.abi --pkg=expertise --out=Expertise.go