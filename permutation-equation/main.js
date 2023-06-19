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
 * Complete the 'permutationEquation' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts INTEGER_ARRAY p as parameter.
 */
function permutationEquation(p) {
    const pMap = {}
    for (let i = 0; i < p.length; i++) {
        pMap[p[i]] = i + 1
    }

    const res = []
    for (let i = 0; i < p.length; i++) {
        const py = pMap[i+1]
        const ppy = pMap[py]
        res.push(ppy)
    }
    
    return res
}

function main() {
    // const ws = fs.createWriteStream(process.env.OUTPUT_PATH);

    const n = parseInt(readLine().trim(), 10);

    const p = readLine().replace(/\s+$/g, '').split(' ').map(pTemp => parseInt(pTemp, 10));

    const result = permutationEquation(p);

    console.log(result);

    // ws.write(result.join('\n') + '\n');

    // ws.end();
}
