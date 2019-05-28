pragma solidity ^0.5.1;

contract aicumen {

  mapping (string => string) keyval;

  function setEntity(string memory addr, string memory name) public returns (bool){
    keyval[name]  = addr;
    return true;
  }
  
  function getEntity(string memory name) public view returns(string memory) {
        return keyval[name];
   }
}