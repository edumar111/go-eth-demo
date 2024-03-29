pragma solidity ^0.4.24;

import './MultiAdmin.sol';

contract AddressesStore is MultiAdmin {
  mapping (bytes32 => address) private addresses;
  mapping (bytes32 => address[]) private addressesArray;

  // core functions
  function _createAddress(bytes32 key, address value) internal returns (bool) {
    if (addresses[key] != address(0x0)) return false;

    return _updateAddress(key, value);
  }

  function _createAddressArray(bytes32 key, address[] values) internal returns (bool) {
    bool success = true;
    bool save = true;
    address[] memory tempArray = addressesArray[key];
    
    if (tempArray.length == 0) {
      for (uint256 j = 0; j < values.length; j++) {
        addressesArray[key].push(values[j]);
      }
    }else {
      for (uint256 i = 0; i < values.length; i++) {
       save = true;
       for (uint256 k= 0; k < tempArray.length; k++) {
         if(values[i] == tempArray[k]){
            save = false;
         }
       }
       if (save) {
           addressesArray[key].push(values[i]);
       }
      }
    }

    return success;
  }
  
  function _createAddresses(bytes32[] keys, address[] values) internal returns (bool) {
    if (keys.length != values.length) return false;
    
    bool success = true;
    
    for (uint256 i = 0; i < keys.length; i++) {
       bool operationSuccess = _createAddress(keys[i], values[i]);
       if (success && !operationSuccess) success = false;
    }
    
    return success;
  }

  function _updateAddress(bytes32 key, address value) internal returns (bool) {
    addresses[key] = value;
    
    return true;
  }

  function _updateAddresses(bytes32[] keys, address[] values) internal returns (bool) {
    if (keys.length != values.length) return false;
    
    bool success = true;
    
    for (uint256 i = 0; i < keys.length; i++) {
      bool operationSuccess = _updateAddress(keys[i], values[i]);
      if (success && !operationSuccess) success = false;
    }
    
    return success;
  }

  function _removeAddress(bytes32 key) internal returns (bool) {
    delete addresses[key];
    
    return true;
  }

  function _removeAddresses(bytes32[] keys) internal returns (bool) {
    bool success = true;
    
    for (uint256 i = 0; i < keys.length; i++) {
      bool operationSuccess = _removeAddress(keys[i]);
      if (success && !operationSuccess) success = false;
    }
    
    return success;
  }
  
    function _removeAddressArray(bytes32 key, address[] values) internal returns (bool) {
    bool success = false;
    address[] memory tempArray = addressesArray[key];
    
    for (uint256 i = 0; i < values.length; i++) {
      for (uint256 k = 0; k < tempArray.length; k++) { 
        if(tempArray[k] == values[i]) {
         delete addressesArray[key][k];
         addressesArray[key][k] = tempArray[tempArray.length -1];
         addressesArray[key].length--;
         success = true;
        }
      }
    }
    
    return success;
    }

  function _readAddress(bytes32 key) internal view returns (address) {
    return addresses[key];
  }

  function _readAddresses(bytes32[] keys) internal view returns (address[]) {
    address[] memory result = new address[](keys.length);
    
    for (uint256 i = 0; i < keys.length; i++) {
      result[i] = _readAddress(keys[i]);
    }
    
    return result;
  }
    
    function _readAddressArray(bytes32 key) internal view returns (address[]) {
    return addressesArray[key];
  }

  // external
  function createAddress(bytes32 key, address value) external onlyAdmin returns (bool) {
    return _createAddress(key, value);
  }
  
  function createAddresses(bytes32[] keys, address[] values) external onlyAdmin returns (bool) {
    return _createAddresses(keys, values);
  }
    function createAddressArray(bytes32 key, address[] values) external onlyAdmin returns (bool) {
      return _createAddressArray(key, values);
  }
  function updateAddress(bytes32 key, address value) external onlyAdmin returns (bool) {
    return _updateAddress(key, value);   
  }
  
  function updateAddresses(bytes32[] keys, address[] values) external onlyAdmin returns (bool) {
    return _updateAddresses(keys, values);
  }
  
  function removeAddress(bytes32 key) external onlyAdmin returns (bool) {
    return _removeAddress(key);
  }
  
  function removeAddresses(bytes32[] keys) external onlyAdmin returns (bool) {
    return _removeAddresses(keys);
  }
  
  function removeAddressArray(bytes32 key, address[] values) external onlyAdmin returns (bool) {
    return _removeAddressArray(key, values); 
  }

  function readAddress(bytes32 key) external view returns (address) {
    return _readAddress(key);
  }
  
  function readAddresses(bytes32[] keys) external view returns (address[]) {
    return _readAddresses(keys);
  }
  
  function readAddressArray(bytes32 key) external view returns (address[]) {
      return _readAddressArray(key);
  }
}