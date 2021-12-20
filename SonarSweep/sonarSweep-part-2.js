const fs = require('fs');

const data = fs.readFileSync('./sonarsweep.txt', 'utf8');

const measurements = data.split('\n').map(Number);
measurements.pop();

let depthIncrease = 0;

for (let i = 1; i < measurements.length - 2; i++) {
  const prevSlidingWindowSum =
    measurements[i - 1] + measurements[i] + measurements[i + 1];

  const currentSlidingWindowSum =
    measurements[i] + measurements[i + 1] + measurements[i + 2];

  console.log({ prevSlidingWindowSum, currentSlidingWindowSum });

  if (currentSlidingWindowSum > prevSlidingWindowSum) depthIncrease++;
}

console.log(depthIncrease);
