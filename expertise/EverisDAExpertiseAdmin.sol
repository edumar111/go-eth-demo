pragma solidity ^0.4.24;
import "./MultiAdmin.sol";
import "./EternalStorage.sol";

contract EverisDAExpertiseAdmin is MultiAdmin {
        EternalStorage public eternalStorage = EternalStorage(0);
        event SubjectAdded(string _lineofService, string _subject);
        event SubjectRemoved(string _lineofService, string _subject);
        event GoalCreated(string id, string category, string name, string level, string subject, uint points);
        event GoalDeleted(string id, string category);
        event AchievementAddedByAdmin(string id, string category, address user);
        event UpdatedGoal(string id, string key, string value);

        constructor(address _eternalStorage) public {
            eternalStorage = EternalStorage(_eternalStorage);
        }

        function createGoal(string id, string category, string goalName, string level, string subject, uint points) public onlyAdmin {
            bytes32 idB = stringToBytes32(id);
            eternalStorage.createBytes32(keccak256(abi.encodePacked(idB, "goal.name")), stringToBytes32(goalName));
            eternalStorage.createBytes32(keccak256(abi.encodePacked(idB, "goal.level")), stringToBytes32(level));
            eternalStorage.createBytes32(keccak256(abi.encodePacked(idB, "goal.subject")), stringToBytes32(subject));
            eternalStorage.createUint256(keccak256(abi.encodePacked(idB, "goal.points")), points);

            bytes32[] memory tempId = new bytes32[](1);
            tempId[0] = idB;
            eternalStorage.createBytes32Array(keccak256(abi.encodePacked(category)), tempId);
            emit GoalCreated(id, category, goalName, level, subject, points);
        }

        function deleteGoal(string id, string category) public onlyAdmin {
            bytes32 idB = stringToBytes32(id);
            bytes32[] memory tempId = new bytes32[](1);
            tempId[0] = idB;
            eternalStorage.removeBytes32Array(keccak256(abi.encodePacked(category)), tempId);
            emit GoalDeleted(id, category);
        }

        function updateGoal(string id, string key, string value) public onlyAdmin {
            bytes32 idB = stringToBytes32(id);
            eternalStorage.updateBytes32(keccak256(abi.encodePacked(idB, key)), stringToBytes32(value));
            emit UpdatedGoal(id,key,value);
        }

        function addSubject(string _subject) public onlyAdmin {
            bytes32 subject = stringToBytes32(_subject);
            bytes32[] memory tempVal = new bytes32[](1);
            tempVal[0] = subject;
            eternalStorage.createBytes32Array(keccak256(abi.encodePacked("da", "subject")), tempVal);
            emit SubjectAdded("da", _subject);
        }

        function removeSubject(string _subject) public onlyAdmin {
            bytes32 subject = stringToBytes32(_subject);
            bytes32[] memory tempVal = new bytes32[](1);
            tempVal[0] = subject;
            eternalStorage.removeBytes32Array(keccak256(abi.encodePacked("da", "subject")), tempVal);
            emit SubjectRemoved("da", _subject);
        }

        function addAchievement(string id, string category, address user) public onlyAdmin {
            bytes32 idB = stringToBytes32(id);
            bytes32[] memory tempId = new bytes32[](1);
            tempId[0] = idB;
            eternalStorage.createBytes32Array(keccak256(abi.encodePacked(user, category)), tempId);
            emit AchievementAddedByAdmin(id, category, user);
        }

        function addOffice(string office) public onlyAdmin {
            bytes32 officeB = stringToBytes32(office);
            bytes32[] memory tempVal = new bytes32[](1);
            tempVal[0] = officeB;
            eternalStorage.createBytes32Array(keccak256(abi.encodePacked("everis.offices")), tempVal);
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
