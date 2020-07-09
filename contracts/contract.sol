pragma solidity ^0.6.0;

contract contrato1 {
    
    bool public IsAuthenticated = false;
    address private owner;
    address private customer;
    
    // modifier to check if caller is owner
    modifier isOwner() {
        require(msg.sender == owner, "Caller is not owner");
        _;
    }

    constructor(address customerAdress) public {
        // 'msg.sender' is sender of current call, contract deployer for a constructor
        owner = msg.sender; 
        customer = customerAdress;
    }
    
    function SetIsAuthentication(bool isAuthenticated) public isOwner {
        IsAuthenticated = isAuthenticated;
    }
    
    function GetCustomer() public view isOwner returns (address) {
        return customer;
    }
}