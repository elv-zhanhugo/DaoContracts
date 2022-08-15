#!/bin/bash
set -Eeuo pipefail

run_solc_abi() {
    solc --abi "${1}" --optimize -o "${2}" --evm-version istanbul --overwrite
}


run_solc_bin() {
    solc --bin "${1}" --optimize -o "${2}" --evm-version istanbul --overwrite
}

run_abigen(){
  if abigen --bin "${1}" --abi "${2}" --pkg="${3}" --out "${4}"; then
    echo "SUCCESS : The go binding for ${2} is present at ${4}"
  else
    echo "FAILED : error occured while creating go binding!"
    exit 1
  fi
}


hash solc || {
    echo "Error : solc is not found"
    exit 1
}

# we should also check the version here
hash abigen || {
    echo "Error : abigen is not found (must use go-ethereum/cmd/abigen)"
    exit 1
}


separator="====================================================================================================="

parent_dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
build_dir="$( mkdir -p "$parent_dir/build" && cd "$parent_dir/build" && pwd )"
box_dir="$( mkdir -p "$build_dir/box" && cd "$build_dir/box" && pwd )"
dex_dir="$( mkdir -p "$build_dir/dex" && cd "$build_dir/dex" && pwd )"
governance_timelock_dir="$( mkdir -p "$build_dir/governance_timelock" && cd "$build_dir/governance_timelock" && pwd )"
governor_contract_dir="$( mkdir -p "$build_dir/governor_contract" && cd "$build_dir/governor_contract" && pwd )"
governance_token_dir="$( mkdir -p "$build_dir/governance_token" && cd "$build_dir/governance_token" && pwd )"

run_solc_abi "$parent_dir/contracts/Box.sol" "$box_dir/abi"
run_solc_bin "$parent_dir/contracts/Box.sol" "$box_dir/bin"
run_abigen "$box_dir/bin/Box.bin" "$box_dir/abi/Box.abi" "box" "$box_dir/box.go"
echo -e "\n${separator}\n"

run_solc_abi "$parent_dir/contracts/Dex.sol" "$dex_dir/abi"
run_solc_bin "$parent_dir/contracts/Dex.sol" "$dex_dir/bin"
run_abigen "$dex_dir/bin/Dex.bin" "$dex_dir/abi/Dex.abi" "dex" "$dex_dir/Dex.go"
echo -e "\n${separator}\n"

run_solc_abi "$parent_dir/contracts/GovernanceToken.sol" "$governance_token_dir/abi"
run_solc_bin "$parent_dir/contracts/GovernanceToken.sol" "$governance_token_dir/bin"
run_abigen "$governance_token_dir/bin/GovernanceToken.bin" "$governance_token_dir/abi/GovernanceToken.abi" "governance_token" "$governance_token_dir/governance_token.go"
echo -e "\n${separator}\n"

run_solc_abi "$parent_dir/contracts/GovernanceTimeLock.sol" "$governance_timelock_dir/abi"
run_solc_bin "$parent_dir/contracts/GovernanceTimeLock.sol" "$governance_timelock_dir/bin"
run_abigen "$governance_timelock_dir/bin/GovernanceTimeLock.bin" "$governance_timelock_dir/abi/GovernanceTimeLock.abi" "governance_timelock" "$governance_timelock_dir/governance_timelock.go"
echo -e "\n${separator}\n"

run_solc_abi "$parent_dir/contracts/GovernorContract.sol" "$governor_contract_dir/abi"
run_solc_bin "$parent_dir/contracts/GovernorContract.sol" "$governor_contract_dir/bin"
run_abigen "$governor_contract_dir/bin/GovernorContract.bin" "$governor_contract_dir/abi/GovernorContract.abi" "governor_contract" "$governor_contract_dir/governor_contract.go"
echo -e "\n${separator}\n"
