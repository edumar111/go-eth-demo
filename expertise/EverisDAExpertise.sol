pragma solidity ^0.4.24;

import './EternalStorage.sol';
import "./MultiOwned.sol";

contract LevelsInterface {
    function getLevelRank(uint score) public view returns (string level, string rank);
}

contract EverisDAExpertise is MultiOwned{

    LevelsInterface public levelManager;
    EternalStorage public eternalStorage = EternalStorage(0);
    string public constant ABI_URL = "QmYsCfXnKGNBkPz6dfYuQuHkBD5rkUToAAVD7eAJZswhU1";
    string[5] public categories;

    //Events
    event LevelManagerSet(address oldLevelManager, address newLevelManager);
    event AchievementAdded(string id, string category, address user);
    event AchievementRemoved(string id, string category, address user);
    event AchievementInfoAdded(address _addr, string _idAchievement, string _key, string _value);
    event AchievementInfoAdded(address _addr, string _idAchievement, string _key, bytes32 _value);
    event InfoAdded(address _addr, string _key, string _value);
    event UserAdded(address user);

    constructor(address _levelManager, address _eternalStorage) MultiOwned(msg.sender) public {
        eternalStorage = EternalStorage(_eternalStorage);
        categories = ['courses', 'handsOnExp', 'evangelism', 'education', 'research'];
        setLevelManager(_levelManager);
    }

    function setLevelManager(address newLevelManager) public onlyOwner {
        levelManager = LevelsInterface(newLevelManager);
        emit LevelManagerSet(levelManager, newLevelManager);
    }

    function expertise(address user) public view returns(uint score, string level, string rank) {
        score = getAllScore(user);
        (level, rank) = levelManager.getLevelRank(score);
        return;
    }
    
    function addAchievement(string id, string category) public {
        bytes32 idB = stringToBytes32(id);
        bytes32[] memory tempId = new bytes32[](1);
        tempId[0] = idB;
        eternalStorage.createBytes32Array(keccak256(abi.encodePacked(msg.sender, category)), tempId);
        emit AchievementAdded(id, category, msg.sender);
    }

    function removeAchievement(string id, string category) public {
        bytes32 idB = stringToBytes32(id);
        bytes32[] memory tempId = new bytes32[](1);
        tempId[0] = idB;
        eternalStorage.removeBytes32Array(keccak256(abi.encodePacked(msg.sender, category)), tempId);
        emit AchievementRemoved(id, category, msg.sender);
    }
    
    function setAchievementInfo(string _id, string _key, string _value) public {
        bytes32 value = stringToBytes32(_value);
        bytes32 id = stringToBytes32(_id);
        eternalStorage.updateBytes32(keccak256(abi.encodePacked(id, _key)), value);
        emit AchievementInfoAdded(msg.sender, _id, _key, _value);
    }
    
    function setAchievementInfoBytes32(string _id, string _key, bytes32 _value) public {
        bytes32 id = stringToBytes32(_id);
        eternalStorage.updateBytes32(keccak256(abi.encodePacked(id, _key)), _value);
        emit AchievementInfoAdded(msg.sender, _id, _key, _value);
    }
    
    function getAchievementInfo(string _id, string _key) public view returns (bytes32){
        bytes32 id = stringToBytes32(_id);
        bytes32 info = eternalStorage.readBytes32(keccak256(abi.encodePacked(id, _key)));
        return info;
    }
    
    function _addUser(address user) private {
        address[] memory tempVal = new address[](1);
        tempVal[0] = user;
        eternalStorage.createAddressArray(keccak256(abi.encodePacked('listUsers')), tempVal);
        emit UserAdded(user);
    }
    
    function getUsers() public view returns(address[]) {
        address[] memory users = eternalStorage.readAddressArray(keccak256(abi.encodePacked('listUsers')));
        return (users);
    }

    function setUserInfo(string _key, string _value) public {
        bytes32 value = stringToBytes32(_value);
        bytes32[] memory tempVal = new bytes32[](1);
        tempVal[0] = value;
        eternalStorage.createBytes32Array(keccak256(abi.encodePacked(msg.sender, _key)), tempVal);
        _addUser(msg.sender);
        emit InfoAdded(msg.sender, _key, _value);
    }
    
    function getUserInfo(address _user, string _key) public view returns (bytes32[]){
        bytes32[] memory myInfo = eternalStorage.readBytes32Array(keccak256(abi.encodePacked(_user,_key)));
        return (myInfo);
    }
    
    function getSubject() public view returns (bytes32[]){
        bytes32[] memory subjects = eternalStorage.readBytes32Array(keccak256(abi.encodePacked("da","subject")));
        return (subjects);
    }
    
    function getOffices() public view returns (bytes32[]){
        bytes32[] memory offices = eternalStorage.readBytes32Array(keccak256(abi.encodePacked("everis.offices")));
        return (offices);
    }
    
    function getAchievementsIds(address user) public view returns (bytes32[], bytes32[]){
        bytes32[] idAchievements;
        bytes32[] idGoals;
        for (uint i = 0; i < categories.length; i++) {
            bytes32[] memory myAchievementsId = eternalStorage.readBytes32Array(keccak256(abi.encodePacked(user, categories[i])));
            for (uint j = 0; j < myAchievementsId.length; j++) {
                idAchievements.push(myAchievementsId[j]);
                idGoals.push(eternalStorage.readBytes32(keccak256(abi.encodePacked(myAchievementsId[j], "goalId"))));
            }
        }
        return (idAchievements, idGoals);
    }
    
    function getAchievements(address user, string category) public view returns (bytes32[], bytes32[], bytes32[], uint[], bytes32[]){
        bytes32[] memory myAchievementsId = eternalStorage.readBytes32Array(keccak256(abi.encodePacked(user, category)));
        return getArraysAchievements(myAchievementsId, category);
    }
    
    function getArraysAchievements(bytes32[] myAchievementsId, string category) internal view returns (bytes32[], bytes32[], bytes32[], uint[], bytes32[]){
        bytes32[] memory ids = new bytes32[](myAchievementsId.length);
        bytes32[] memory names = new bytes32[](myAchievementsId.length);
        bytes32[] memory levels = new bytes32[](myAchievementsId.length);
        bytes32[] memory subjects = new bytes32[](myAchievementsId.length);
        uint[] memory points = new uint[](myAchievementsId.length);
        for (uint i = 0; i < myAchievementsId.length; i++) {
            bytes32 goalId = eternalStorage.readBytes32(keccak256(abi.encodePacked(myAchievementsId[i], "goalId")));
            ids[i] = myAchievementsId[i];
            names[i] = eternalStorage.readBytes32(keccak256(abi.encodePacked(goalId, "goal.name")));
            levels[i] = eternalStorage.readBytes32(keccak256(abi.encodePacked(goalId, "goal.level")));
            subjects[i] = eternalStorage.readBytes32(keccak256(abi.encodePacked(goalId, "goal.subject")));
            points[i] = eternalStorage.readUint256(keccak256(abi.encodePacked(goalId, "goal.points")));
        }
        return (names, levels, subjects, points, ids);
    }
    
    function getGoals(string category) public view returns (bytes32[], bytes32[], bytes32[], bytes32[], uint[]){
        bytes32[] memory availableIds = eternalStorage.readBytes32Array(keccak256(abi.encodePacked(category)));
        uint length = availableIds.length;
        bytes32[] memory ids = new bytes32[](length);
        bytes32[] memory names = new bytes32[](length);
        bytes32[] memory levels = new bytes32[](length);
        bytes32[] memory subjects = new bytes32[](length);
        uint[] memory points = new uint[](length);
        for (uint i = 0; i < length; i++) {
            ids[i] = availableIds[i];
            names[i] = eternalStorage.readBytes32(keccak256(abi.encodePacked(availableIds[i], "goal.name")));
            levels[i] = eternalStorage.readBytes32(keccak256(abi.encodePacked(availableIds[i], "goal.level")));
            subjects[i] = eternalStorage.readBytes32(keccak256(abi.encodePacked(availableIds[i], "goal.subject")));
            points[i] = eternalStorage.readUint256(keccak256(abi.encodePacked(availableIds[i], "goal.points")));
        }
        return (ids, names, levels, subjects, points);
    }

    function getScore(address user, string category) public view returns (uint score) {
        bytes32[] memory myAchievementsId = eternalStorage.readBytes32Array(keccak256(abi.encodePacked(user,category)));
        uint length = myAchievementsId.length;
        score = 0;
        for (uint i = 0; i < length; i++) {
            bytes32 goalId = eternalStorage.readBytes32(keccak256(abi.encodePacked(myAchievementsId[i], "goalId")));
            score += eternalStorage.readUint256(keccak256(abi.encodePacked(goalId, "goal.points")));
        }
        return (score);
    }

    function getAllScore(address user) public view returns (uint256) {
        uint256 score = 0;
        uint length = categories.length;

        for (uint i = 0; i < length; i++) {
           bytes32[] memory myAchievementsId = eternalStorage.readBytes32Array(keccak256(abi.encodePacked(user,categories[i])));
            for (uint j = 0; j < myAchievementsId.length; j++) {
                bytes32 goalId = eternalStorage.readBytes32(keccak256(abi.encodePacked(myAchievementsId[j], "goalId")));
                score += eternalStorage.readUint256(keccak256(abi.encodePacked(goalId, "goal.points")));
            }
        }
        return (score);
    }
    
    function stringToBytes32(string memory source) public pure returns (bytes32 result) {
        bytes memory tempEmptyStringTest = bytes(source);
        if (tempEmptyStringTest.length == 0) {
            return 0x0;
        }
        assembly {
            result := mload(add(source, 32))
        }
    }

}