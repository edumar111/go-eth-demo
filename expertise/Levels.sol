pragma solidity ^0.4.24;

contract Levels {

    event StepDefined(uint minScore, uint maxScore, string indexed level, string indexed rank);

    struct Step {
        string level;
        string rank;
        uint maxScore;
    }
    
    mapping (uint => Step) steps;

    constructor() public {
        uint maxint = 0; maxint--;
        defineStep(0, 49, 'dan1', 'everis DA Shodan');
        defineStep(50, 99, 'dan2', 'everis DA Nidan');
        defineStep(100, 299, 'dan3', 'everis DA Sandan');
        defineStep(300, 599, 'dan4', 'everis DA Yondan');
        defineStep(600, 999, 'dan5', 'everis DA Godan');
        defineStep(1000, 1499, 'dan6', 'everis DA Rokudan');
        defineStep(1500, 2099, 'dan7', 'everis DA Shichidan');
        defineStep(2100, 2799, 'dan8', 'everis DA Hachidan');
        defineStep(2800, 3599, 'dan9', 'everis DA Kudan');
        defineStep(3600, maxint, 'dan10', 'everis DA Jūdan');
    }

    function defineStep(uint minScore, uint maxScore, string level, string rank) public {
        steps[minScore] = Step(level, rank, maxScore);
        emit StepDefined(minScore, maxScore, level, rank);
    }

    function getLevelRank(uint score) public view returns (string level, string rank) {
        Step memory step = steps[0];
        while (score > step.maxScore) {
            step = steps[step.maxScore + 1];
        }
        return (step.level, step.rank);
    }
}