'use strict';

const fs = require('fs');

process.stdin.resume();
process.stdin.setEncoding('utf-8');

let inputString = '';
let currentLine = 0;

process.stdin.on('data', function (inputStdin) {
    inputString += inputStdin;
});

process.stdin.on('end', function () {
    inputString = inputString.split('\n');

    main();
});

function readLine() {
    return inputString[currentLine++];
}

/*
 * Complete the 'findDigits' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER n as parameter.
 */

function findDigits(n) {
    const str = n.toString();

    return str.split("").reduce((acc, s) => {
        console.log("s: ", s);
        console.log("acc: ", acc);
        console.log("parseInt(s): ", n % parseInt(s));
        console.log("--------------------")

        if (s == 0) {
            return acc
        }

        if (n % parseInt(s) == 0) {
            return acc + 1
        }

        return acc
    }, 0)
}

function main() {
    // const ws = fs.createWriteStream(process.env.OUTPUT_PATH);

    const t = parseInt(readLine().trim(), 10);

    for (let tItr = 0; tItr < t; tItr++) {
        const n = parseInt(readLine().trim(), 10);

        const result = findDigits(n);

        console.log(result);
        // ws.write(result + '\n');
    }

    // ws.end();
}
