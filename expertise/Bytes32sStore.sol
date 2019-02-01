pragma solidity ^0.4.24;

import './MultiAdmin.sol';

contract Bytes32sStore is MultiAdmin {
    
  mapping (bytes32 => bytes32) private bytes32s;
  mapping (bytes32 => bytes32[]) private bytes32Array;

  // core functions
  function _createBytes32(bytes32 key, bytes32 value) internal returns (bool) {
    if (bytes32s[key] != bytes32(0x0)) return false;

    return _updateBytes32(key, value);
  }
  
  function _createBytes32Array(bytes32 key, bytes32[] values) internal returns (bool) {
    bool success = true;
    bool save = true;
    bytes32[] memory tempArray = bytes32Array[key];
    
    if (tempArray.length == 0) {
      for (uint256 j = 0; j < values.length; j++) {
        bytes32Array[key].push(values[j]);
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
           bytes32Array[key].push(values[i]);
       }
      }
    }

    return success;
  }
  


  function _createBytes32s(bytes32[] keys, bytes32[] values) internal returns (bool) {
    if (keys.length != values.length) return false;
    
    bool success = true;
    
    for (uint256 i = 0; i < keys.length; i++) {
       bool operationSuccess = _createBytes32(keys[i], values[i]);
       if (success && !operationSuccess) success = false;
    }
    
    return success;
  }

  function _updateBytes32(bytes32 key, bytes32 value) internal returns (bool) {
    bytes32s[key] = value;
    
    return true;
  }

  function _updateBytes32s(bytes32[] keys, bytes32[] values) internal returns (bool) {
    if (keys.length != values.length) return false;
    
    bool success = true;
    
    for (uint256 i = 0; i < keys.length; i++) {
      bool operationSuccess = _updateBytes32(keys[i], values[i]);
      if (success && !operationSuccess) success = false;
    }
    
    return success;
  }

  function _removeBytes32(bytes32 key) internal returns (bool) {
    delete bytes32s[key];
    
    return true;
  }

  function _removeBytes32s(bytes32[] keys) internal returns (bool) {
    bool success = true;
    
    for (uint256 i = 0; i < keys.length; i++) {
      bool operationSuccess = _removeBytes32(keys[i]);
      if (success && !operationSuccess) success = false;
    }
    
    return success;
  }
  
  function _removeBytes32Array(bytes32 key, bytes32[] values) internal returns (bool) {
    bool success = false;
    bytes32[] memory tempArray = bytes32Array[key];
    
    for (uint256 i = 0; i < values.length; i++) {
      for (uint256 k = 0; k < tempArray.length; k++) { 
        if(tempArray[k] == values[i]) {
         delete bytes32Array[key][k];
         bytes32Array[key][k] = tempArray[tempArray.length -1];
         bytes32Array[key].length--;
         success = true;
        }
      }
    }
    
    return success;
  }

  function _readBytes32(bytes32 key) internal view returns (bytes32) {
    return bytes32s[key];
  }

  function _readBytes32s(bytes32[] keys) internal view returns (bytes32[]) {
    bytes32[] memory result = new bytes32[](keys.length);
    
    for (uint256 i = 0; i < keys.length; i++) {
      result[i] = _readBytes32(keys[i]);
    }
    
    return result;
  }
  
  function _readBytes32Array(bytes32 key) internal view returns (bytes32[]) {
    return bytes32Array[key];
  }

  // external
  function createBytes32(bytes32 key, bytes32 value) external onlyAdmin returns (bool) {
    return _createBytes32(key, value);
  }
  
  function createBytes32s(bytes32[] keys, bytes32[] values) external onlyAdmin returns (bool) {
    return _createBytes32s(keys, values);
  }
  
  function createBytes32Array(bytes32 key, bytes32[] values) external onlyAdmin returns (bool) {
      return _createBytes32Array(key, values);
  }
  

  function updateBytes32(bytes32 key, bytes32 value) external onlyAdmin returns (bool) {
    return _updateBytes32(key, value);   
  }
  
  function updateBytes32s(bytes32[] keys, bytes32[] values) external onlyAdmin returns (bool) {
    return _updateBytes32s(keys, values);
  }
  
  function removeBytes32(bytes32 key) external onlyAdmin returns (bool) {
    return _removeBytes32(key);
  }
  
  function removeBytes32s(bytes32[] keys) external onlyAdmin returns (bool) {
    return _removeBytes32s(keys);
  }
  
  function removeBytes32Array(bytes32 key, bytes32[] values) external onlyAdmin returns (bool) {
    return _removeBytes32Array(key, values); 
  }

  function readBytes32(bytes32 key) external view returns (bytes32) {
    return _readBytes32(key);
  }
  
  function readBytes32s(bytes32[] keys) external view returns (bytes32[]) {
    return _readBytes32s(keys);
  }
  
  function readBytes32Array(bytes32 key) external view returns (bytes32[]) {
      return _readBytes32Array(key);
  }
}