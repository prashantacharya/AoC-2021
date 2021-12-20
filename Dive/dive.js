const { readFileSync } = require('fs');

const data = readFileSync(process.argv[2], 'utf8');

const commands = data.split('\n');
commands.pop();

const commandObj = commands.map((val) => {
  const [command, by] = val.split(' ');

  return { command, by: +by };
});

let horizontalPosition = 0,
  depth = 0,
  aim = 0;

commandObj.forEach(({ command, by }) => {
  switch (command) {
    case 'forward':
      horizontalPosition += by;
      depth += aim * by;
      break;
    case 'up':
      aim -= by;
      break;
    case 'down':
      aim += by;
  }
});

console.log({ result: horizontalPosition * depth });
