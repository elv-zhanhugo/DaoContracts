// SPDX-License-Identifier: MIT
pragma solidity ^0.8.2;

import "../@openzeppelin/contracts/token/ERC20/extensions/ERC20Votes.sol";

contract DEX {
    event Bought(uint256 amount);

    ERC20Votes public token;

    constructor(ERC20Votes _token) {
        token = _token;
    }

    function buy() payable public {
        uint256 amountTobuy = msg.value;
        uint256 dexBalance = token.balanceOf(address(this));
        require(amountTobuy > 0, "You need to send some ether");
        require(amountTobuy <= dexBalance, "Not enough tokens in the reserve");
        token.transfer(msg.sender, amountTobuy);
        emit Bought(amountTobuy);
    }
}
