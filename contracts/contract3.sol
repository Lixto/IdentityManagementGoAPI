pragma solidity ^0.6.0;

contract contrato3 {
    
    bool public IsAuthenticated = false;
    uint256 public ExpirationTime;
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
    
    function SetAuthentication(bool isAuthenticated, uint256 expirationTime) public isOwner {
        IsAuthenticated = isAuthenticated;
        ExpirationTime = expirationTime;
    }
    
    function GetCustomer() public view isOwner returns (address) {
        return customer;
    }
    
    function GetOwner() public view returns (address){
        return owner;
    }
}