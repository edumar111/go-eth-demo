pragma solidity ^0.4.24;

import './MultiAdmin.sol';

contract Uint256sStore is MultiAdmin {
  mapping (bytes32 => uint256) private uint256s;
  mapping (bytes32 => uint256[]) private uint256sArray;

  // core functions
  function _createUint256(bytes32 key, uint256 value) internal returns (bool) {
    if (uint256s[key] > 0) return false;

    return _updateUint256(key, value);
  }

  function _createUint256s(bytes32[] keys, uint256[] values) internal returns (bool) {
    if (keys.length != values.length) return false;
    
    bool success = true;
    
    for (uint256 i = 0; i < keys.length; i++) {
       bool operationSuccess = _createUint256(keys[i], values[i]);
       if (success && !operationSuccess) success = false;
    }
    
    return success;
  }

  function _updateUint256(bytes32 key, uint256 value) internal returns (bool) {
    uint256s[key] = value;
    
    return true;
  }

  function _updateUint256s(bytes32[] keys, uint256[] values) internal returns (bool) {
    if (keys.length != values.length) return false;
    
    bool success = true;
    
    for (uint256 i = 0; i < keys.length; i++) {
      bool operationSuccess = _updateUint256(keys[i], values[i]);
      if (success && !operationSuccess) success = false;
    }
    
    return success;
  }

  function _removeUint256(bytes32 key) internal returns (bool) {
    delete uint256s[key];
    
    return true;
  }

  function _removeUint256s(bytes32[] keys) internal returns (bool) {
    bool success = true;
    
    for (uint256 i = 0; i < keys.length; i++) {
      bool operationSuccess = _removeUint256(keys[i]);
      if (success && !operationSuccess) success = false;
    }
    
    return success;
  }

  function _readUint256(bytes32 key) internal view returns (uint256) {
    return uint256s[key];
  }

  function _readUint256s(bytes32[] keys) internal view returns (uint256[]) {
    uint256[] memory result = new uint256[](keys.length);
    
    for (uint256 i = 0; i < keys.length; i++) {
      result[i] = _readUint256(keys[i]);
    }
    
    return result;
  }


  function _createUint256sArray(bytes32 key, uint256[] values) internal returns (bool) {
    bool success = true;
    bool save = true;
    uint256[] memory tempArray = uint256sArray[key];
    
    if (tempArray.length == 0) {
      for (uint256 j = 0; j < values.length; j++) {
        uint256sArray[key].push(values[j]);
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
           uint256sArray[key].push(values[i]);
       }
      }
    }

    return success;
  }

  function _removeUint256sArray(bytes32 key, uint256[] values) internal returns (bool) {
    bool success = false;
    uint256[] memory tempArray = uint256sArray[key];
    
    for (uint256 i = 0; i < values.length; i++) {
      for (uint256 k = 0; k < tempArray.length; k++) { 
        if(tempArray[k] == values[i]) {
         delete uint256sArray[key][k];
         uint256sArray[key][k] = tempArray[tempArray.length -1];
         uint256sArray[key].length--;
         success = true;
        }
      }
    }
    
    return success;
  }
  
  function _readUint256sArray(bytes32 key) internal view returns (uint256[]) {
    return uint256sArray[key];
  }

  // external
  function createUint256(bytes32 key, uint256 value) external onlyAdmin returns (bool) {
    return _createUint256(key, value);
  }
  
  function createUint256s(bytes32[] keys, uint256[] values) external onlyAdmin returns (bool) {
    return _createUint256s(keys, values);
  }
  
  function updateUint256(bytes32 key, uint256 value) external onlyAdmin returns (bool) {
    return _updateUint256(key, value);   
  }
  
  function updateUint256s(bytes32[] keys, uint256[] values) external onlyAdmin returns (bool) {
    return _updateUint256s(keys, values);
  }
  
  function removeUint256(bytes32 key) external onlyAdmin returns (bool) {
    return _removeUint256(key);
  }
  
  function removeUint256s(bytes32[] keys) external onlyAdmin returns (bool) {
    return _removeUint256s(keys);
  }

  function readUint256(bytes32 key) external view returns (uint256) {
    return _readUint256(key);
  }
  
  function readUint256s(bytes32[] keys) external view returns (uint256[]) {
    return _readUint256s(keys);
  }
  
  function createUint256sArray(bytes32 key, uint256[] values) external onlyAdmin returns (bool) {
      return _createUint256sArray(key, values);
  }
  
  function removeUint256sArray(bytes32 key, uint256[] values) external onlyAdmin returns (bool) {
    return _removeUint256sArray(key, values); 
  }
  
  function readUint256sArray(bytes32 key) external view returns (uint256[]) {
      return _readUint256sArray(key);
  }
}