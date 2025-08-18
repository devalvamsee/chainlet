pragma solidity ^0.6.6;

contract ChainletGravityCancellation {

    event __ChainletCancelSendToEvmChain(address indexed sender, uint256 id);

    // Cancel a send to chain transaction considering if it hasn't been batched yet.
    function cancelTransaction(uint256 id) public {
        emit __ChainletCancelSendToEvmChain(msg.sender, id);
    }
}
