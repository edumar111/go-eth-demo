pragma solidity ^0.4.24;

import './AddressesStore.sol';
import './BoolsStore.sol';
import './Bytes32sStore.sol';
import './Int256sStore.sol';
import './Uint256sStore.sol';

contract EternalStorage is AddressesStore, BoolsStore, Bytes32sStore, Int256sStore, Uint256sStore {}