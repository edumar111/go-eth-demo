pragma solidity ^0.4.24;

import './MultiAdmin.sol';

contract BoolsStore is MultiAdmin {
  mapping (bytes32 => bool) private bools;
    mapping (bytes32 => bool[]) private boolsArray;


  // core functions
  function _createBool(bytes32 key, bool value) internal returns (bool) {
    if (bools[key]) return false;

    return _updateBool(key, value);
  }

  function _createBools(bytes32[] keys, bool[] values) internal returns (bool) {
    if (keys.length != values.length) return false;
    
    bool success = true;
    
    for (uint256 i = 0; i < keys.length; i++) {
       bool operationSuccess = _createBool(keys[i], values[i]);
       if (success && !operationSuccess) success = false;
    }
    
    return success;
  }
  
  function _createBoolsArray(bytes32 key, bool[] values) internal returns (bool) {
    bool success = true;
    bool save = true;
    bool[] memory tempArray = boolsArray[key];
    
    if (tempArray.length == 0) {
      for (uint256 j = 0; j < values.length; j++) {
        boolsArray[key].push(values[j]);
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
           boolsArray[key].push(values[i]);
       }
      }
    }

    return success;
  }

  function _removeBoolsArray(bytes32 key, bool[] values) internal returns (bool) {
    bool success = false;
    bool[] memory tempArray = boolsArray[key];
    
    for (uint256 i = 0; i < values.length; i++) {
      for (uint256 k = 0; k < tempArray.length; k++) { 
        if(tempArray[k] == values[i]) {
         delete boolsArray[key][k];
         boolsArray[key][k] = tempArray[tempArray.length -1];
         boolsArray[key].length--;
         success = true;
        }
      }
    }
    
    return success;
  }
  
  function _readBoolsArray(bytes32 key) internal view returns (bool[]) {
    return boolsArray[key];
  }
  function _updateBool(bytes32 key, bool value) internal returns (bool) {
    bools[key] = value;
    
    return true;
  }

  function _updateBools(bytes32[] keys, bool[] values) internal returns (bool) {
    if (keys.length != values.length) return false;
    
    bool success = true;
    
    for (uint256 i = 0; i < keys.length; i++) {
      bool operationSuccess = _updateBool(keys[i], values[i]);
      if (success && !operationSuccess) success = false;
    }
    
    return success;
  }

  function _removeBool(bytes32 key) internal returns (bool) {
    delete bools[key];
    
    return true;
  }

  function _removeBools(bytes32[] keys) internal returns (bool) {
    bool success = true;
    
    for (uint256 i = 0; i < keys.length; i++) {
      bool operationSuccess = _removeBool(keys[i]);
      if (success && !operationSuccess) success = false;
    }
    
    return success;
  }

  function _readBool(bytes32 key) internal view returns (bool) {
    return bools[key];
  }

  function _readBools(bytes32[] keys) internal view returns (bool[]) {
    bool[] memory result = new bool[](keys.length);
    
    for (uint256 i = 0; i < keys.length; i++) {
      result[i] = _readBool(keys[i]);
    }
    
    return result;
  }

  // external
  function createBool(bytes32 key, bool value) external onlyAdmin returns (bool) {
    return _createBool(key, value);
  }
  
  function createBools(bytes32[] keys, bool[] values) external onlyAdmin returns (bool) {
    return _createBools(keys, values);
  }
  
  function updateBool(bytes32 key, bool value) external onlyAdmin returns (bool) {
    return _updateBool(key, value);   
  }
  
  function updateBools(bytes32[] keys, bool[] values) external onlyAdmin returns (bool) {
    return _updateBools(keys, values);
  }
  
  function removeBool(bytes32 key) external onlyAdmin returns (bool) {
    return _removeBool(key);
  }
  
  function removeBools(bytes32[] keys) external onlyAdmin returns (bool) {
    return _removeBools(keys);
  }

  function readBool(bytes32 key) external view returns (bool) {
    return _readBool(key);
  }
  
  function readBools(bytes32[] keys) external view returns (bool[]) {
    return _readBools(keys);
  }
  
  function createBytes32Array(bytes32 key, bool[] values) external onlyAdmin returns (bool) {
      return _createBoolsArray(key, values);
  }
  
  function removeBoolsArray(bytes32 key, bool[] values) external onlyAdmin returns (bool) {
    return _removeBoolsArray(key, values); 
  }
  
  function readBoolsArray(bytes32 key) external view returns (bool[]) {
      return _readBoolsArray(key);
  }
}