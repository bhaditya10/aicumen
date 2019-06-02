pragma solidity ^0.5.9;

contract CarContract
{   
    
    // structure list
    struct  Car{  
        uint256 cid;
        string Model;
        bytes32 Color;
        bytes32 Type;
        string Engine;
        string Brand;
        uint32 TopSpeed;
        uint32 Weight;
        bool Available;
        uint256 LaunchDate;
        uint256 Price;
    }
    mapping(uint256 => Car) cars;
    mapping(uint256 => address) ownerOfCar;
    mapping (uint256 => uint256[]) PriceToCar;
    mapping (string => uint256[]) ModelsToCar;
    
    // constructor for contract
    constructor() public {
    }
   
    
    function getCar(uint256 cid) public view returns (string memory, bytes32, bytes32, string memory, string memory, uint32, uint32, bool, uint256, uint256) {
        Car memory c = cars[cid];
        return (
            c.Model,
            c.Color,
            c.Type,
            c.Engine,
            c.Brand,
            c.TopSpeed,
            c.Weight,
            c.Available,
            c.LaunchDate,
            c.Price
        );
    }
    
    function stringToBytes32(string memory source) public view returns (bytes32 result) {
    bytes memory tempEmptyStringTest = bytes(source);
    if (tempEmptyStringTest.length == 0) {
        return 0x0;
    }

    assembly {
        result := mload(add(source, 32))
    }
    
}
    
    function Availability(uint256 index) public view returns(bool){ // change on what you need to return
        return cars[index].Available;
    }
    
    function setCar(uint256 cid, string memory _Model, string memory _Color, string memory _Type, string memory _Engine, string memory _Brand,
        uint32 _TopSpeed, uint32 _Weight, bool _Available, uint256 _LaunchDate, uint256 _Price) public{
        require(cars[cid].Weight == 0 , "car is already Available");
        bytes32 clr = stringToBytes32(_Color);
        bytes32 ty = stringToBytes32(_Type);
        Car storage car = cars[cid];
        car.Model=_Model;
        car.Color=clr;
        car.Type= ty;
        car.Engine=_Engine;
        car.Brand=_Brand;
        car.TopSpeed=_TopSpeed;
        car.Weight=_Weight;
        car.Available=_Available;
        car.LaunchDate=_LaunchDate;
        car.Price= _Price;

        // ownerOfCar[cid] = msg.sender;
        PriceToCar[_Price].push(cid);
        
  }
  
  function getCarLengthByPrice(uint256 _Price) public returns (uint256){
      return PriceToCar[_Price].length;
  }
  
  function getCarIdForPrice(uint256 _Price, uint256 _index) public returns (uint256){
    return PriceToCar[_Price][_index];  
  }
    
    
   // buy function
   function buy(uint256 cid) payable public{
       Car storage c = cars[cid];
       require(msg.value ==  c.Price, "price mismatch");
       require(c.Available == true, "car not Available");
       ownerOfCar[cid] = msg.sender;
       c.Available = false;
   }
}
