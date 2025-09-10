pragma solidity ^0.6.6;

contract CltBridge {

    event __ChainletSendCltToIbc(address sender, string recipient, uint256 amount);

    // Pay the contract a certain CRO amount and trigger a CRO transfer
    // from the contract to recipient through IBC
    function send_clt_to_crypto_org(string memory recipient) public payable {
        emit __ChainletSendCltToIbc(msg.sender, recipient, msg.value);
    }
}
