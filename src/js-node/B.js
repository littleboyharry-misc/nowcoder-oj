const readline = require('readline');

const rl = readline.createInterface({
    input: process.stdin,
    output: process.stdout
});
rl.on('line', function (line) {
    handle.next(line);
});
const handle = (function* () {
    let number = parseInt(yield)
    while (number-- > 0) {
        console.log(
            (yield)
                .split(' ')
                .map(i => parseInt(i))
                .reduce((a, b) => a + b)
        )
    }
})()
handle.next();
