const fs = require('fs');

const data = fs.readFileSync('./sonarsweep.txt', 'utf8');

const numberArray = data.split('\n').map(Number);
numberArray.pop();

let depthIncrease = 0;

for (let i = 1; i < numberArray.length; i++) {
  if (numberArray[i] > numberArray[i - 1]) depthIncrease++;
}

console.log(depthIncrease);
