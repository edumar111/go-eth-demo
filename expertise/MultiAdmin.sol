pragma solidity ^0.4.24;

import './MultiOwned.sol';

contract MultiAdmin is MultiOwned {
    
  event AdminAdded(address indexed newAdmin);
  event AdminRemoved(address indexed formerAdmin);
  
  mapping (address => bool) public admins;
  
  modifier onlyAdmin() {
    require(isAdmin(msg.sender));
    _;
  }

  constructor() MultiOwned(msg.sender) public {
    admins[msg.sender] = true;
    emit AdminAdded(msg.sender);
  }

  function isAdmin(address addr) public view returns(bool) {
    return admins[addr];
  }

  function addAdmin(address newAdmin) public onlyOwner {
    admins[newAdmin] = true;
    emit AdminAdded(newAdmin);
  }

  function removeAdmin(address admin) public onlyOwner {
    admins[admin] = false;
    emit AdminRemoved(admin);
  }
  
}