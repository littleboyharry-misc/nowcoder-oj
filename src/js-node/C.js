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
        const [a, b] =
            (yield)
                .split(' ', 2)
                .map(i => parseInt(i))
        if (!(a && b)) break
        else console.log(a + b);
    }
})()
m.next()
