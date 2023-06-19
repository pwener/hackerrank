'use strict';

const fs = require('fs');

process.stdin.resume();
process.stdin.setEncoding('utf-8');

let inputString = '';
let currentLine = 0;

process.stdin.on('data', function(inputStdin) {
    inputString += inputStdin;
});

process.stdin.on('end', function() {
    inputString = inputString.split('\n');

    main();
});

function readLine() {
    return inputString[currentLine++];
}

/*
 * Complete the 'squares' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER a
 *  2. INTEGER b
 */
function squares(a, b) {
    const squareList = []
    
    for (let i = Math.trunc(Math.sqrt(a)); i <= Math.sqrt(b); i++) {
        const n = Math.pow(i, 2);
        if (n > b) {
            break;
        }

        if (n >= a) {
            squareList.push(n);
        }
    }

    console.log(squareList)

    return squareList.length
}

function main() {
    const q = parseInt(readLine().trim(), 10);

    for (let qItr = 0; qItr < q; qItr++) {
        const firstMultipleInput = readLine().replace(/\s+$/g, '').split(' ');

        const a = parseInt(firstMultipleInput[0], 10);

        const b = parseInt(firstMultipleInput[1], 10);

        const result = squares(a, b);

        console.log(result);
    }
}
