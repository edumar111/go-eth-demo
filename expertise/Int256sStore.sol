pragma solidity^0.4.24;

import './MultiAdmin.sol';

contract Int256sStore is MultiAdmin {
  mapping (bytes32 => int256) private int256s;
  mapping (bytes32 => int256[]) private int256sArray;

  // core functions
  function _createInt256(bytes32 key, int256 value) internal returns (bool) {
    if (int256s[key] > 0) return false;

    return _updateInt256(key, value);
  }

  function _createInt256s(bytes32[] keys, int256[] values) internal returns (bool) {
    if (keys.length != values.length) return false;
    
    bool success = true;
    
    for (uint256 i = 0; i < keys.length; i++) {
       bool operationSuccess = _createInt256(keys[i], values[i]);
       if (success && !operationSuccess) success = false;
    }
    
    return success;
  }

  function _updateInt256(bytes32 key, int256 value) internal returns (bool) {
    int256s[key] = value;
    
    return true;
  }

  function _updateInt256s(bytes32[] keys, int256[] values) internal returns (bool) {
    if (keys.length != values.length) return false;
    
    bool success = true;
    
    for (uint256 i = 0; i < keys.length; i++) {
      bool operationSuccess = _updateInt256(keys[i], values[i]);
      if (success && !operationSuccess) success = false;
    }
    
    return success;
  }

  function _removeInt256(bytes32 key) internal returns (bool) {
    delete int256s[key];
    
    return true;
  }

  function _removeInt256s(bytes32[] keys) internal returns (bool) {
    bool success = true;
    
    for (uint256 i = 0; i < keys.length; i++) {
      bool operationSuccess = _removeInt256(keys[i]);
      if (success && !operationSuccess) success = false;
    }
    
    return success;
  }

  function _readInt256(bytes32 key) internal view returns (int256) {
    return int256s[key];
  }

  function _readInt256s(bytes32[] keys) internal view returns (int256[]) {
    int256[] memory result = new int256[](keys.length);
    
    for (uint256 i = 0; i < keys.length; i++) {
      result[i] = _readInt256(keys[i]);
    }
    
    return result;
  }
  
  function _createInt256sArray(bytes32 key, int256[] values) internal returns (bool) {
    bool success = true;
    bool save = true;
    int256[] memory tempArray = int256sArray[key];
    
    if (tempArray.length == 0) {
      for (uint256 j = 0; j < values.length; j++) {
        int256sArray[key].push(values[j]);
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
           int256sArray[key].push(values[i]);
       }
      }
    }

    return success;
  }

  function _removeInt256sArray(bytes32 key, int256[] values) internal returns (bool) {
    bool success = false;
    int256[] memory tempArray = int256sArray[key];
    
    for (uint256 i = 0; i < values.length; i++) {
      for (uint256 k = 0; k < tempArray.length; k++) { 
        if(tempArray[k] == values[i]) {
         delete int256sArray[key][k];
         int256sArray[key][k] = tempArray[tempArray.length -1];
         int256sArray[key].length--;
         success = true;
        }
      }
    }
    
    return success;
  }
  
  function _readInt256sArray(bytes32 key) internal view returns (int256[]) {
    return int256sArray[key];
  }

  // external
  function createInt256(bytes32 key, int256 value) external onlyAdmin returns (bool) {
    return _createInt256(key, value);
  }
  
  function createInt256s(bytes32[] keys, int256[] values) external onlyAdmin returns (bool) {
    return _createInt256s(keys, values);
  }
  
  function updateInt256(bytes32 key, int256 value) external onlyAdmin returns (bool) {
    return _updateInt256(key, value);   
  }
  
  function updateInt256s(bytes32[] keys, int256[] values) external onlyAdmin returns (bool) {
    return _updateInt256s(keys, values);
  }
  
  function removeInt256(bytes32 key) external onlyAdmin returns (bool) {
    return _removeInt256(key);
  }
  
  function removeInt256s(bytes32[] keys) external onlyAdmin returns (bool) {
    return _removeInt256s(keys);
  }

  function readInt256(bytes32 key) external view returns (int256) {
    return _readInt256(key);
  }
  
  function readInt256s(bytes32[] keys) external view returns (int256[]) {
    return _readInt256s(keys);
  }
  
  function createInt256sArray(bytes32 key, int256[] values) external onlyAdmin returns (bool) {
      return _createInt256sArray(key, values);
  }
  
  function removeInt256sArray(bytes32 key, int256[] values) external onlyAdmin returns (bool) {
    return _removeInt256sArray(key, values); 
  }
  
  function readInt256sArray(bytes32 key) external view returns (int256[]) {
      return _readInt256sArray(key);
  }
}