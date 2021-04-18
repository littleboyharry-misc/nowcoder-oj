const readline = require('readline');

const rl = readline.createInterface({
    input: process.stdin,
    output: process.stdout
});
rl.on('line', function (line) {
    m.next(line);
});

const m = (function* () {
    for (; ;) {
        const arr = (yield).split(' ').slice(1).map(i => parseInt(i))
        if (arr.length === 0) break
        console.log(arr.reduce((a, b) => a + b))
    }
})()
m.next()
