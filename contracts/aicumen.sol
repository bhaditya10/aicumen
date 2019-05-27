pragma solidity ^0.5.1;

contract aicumen {

  struct EntityStruct {
    string entityData;
    bool isEntity;
  }

  mapping (address => EntityStruct) public entityStructs;

  function isEntity(address entityAddress) public view returns(bool isIndeed) {
    return entityStructs[entityAddress].isEntity;
  }

  function newEntity(address entityAddress, string memory entityData) public returns(bool success) {
    if(isEntity(entityAddress)) revert(); 
    entityStructs[entityAddress].entityData = entityData;
    entityStructs[entityAddress].isEntity = true;
    return true;
  }
  function getEntity(address entityAddress) public view returns(string memory) {
        return entityStructs[entityAddress].entityData;
   }
}