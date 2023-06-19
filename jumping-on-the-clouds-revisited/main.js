'use strict';

const fs = require('fs');
const { format } = require('path');

process.stdin.resume();
process.stdin.setEncoding('utf-8');

let inputString = '';
let currentLine = 0;

process.stdin.on('data', inputStdin => {
    inputString += inputStdin;
});

process.stdin.on('end', _ => {
    inputString = inputString.replace(/\s*$/, '')
        .split('\n')
        .map(str => str.replace(/\s*$/, ''));

    main();
});

function readLine() {
    return inputString[currentLine++];
}

// Complete the jumpingOnClouds function below.
function jumpingOnClouds(c, k) {
    let energy = 100;

    let i = k

    while (true) {
        const index = i % c.length

        if (c[index]) {
            console.log("thunderbolt", energy, index)
            energy = energy - 3
            console.log("loses energy", energy)

        } else {
            console.log("normal", energy, index)
            energy = energy - 1
            console.log("loses energy", energy)
        }

        if (index == 0) {
            break
        }

        i += k
        console.log("position:", index)
        console.log("-------------------")
    }

    
    return energy
}

function main() {
    // const ws = fs.createWriteStream(process.env.OUTPUT_PATH);

    const nk = readLine().split(' ');

    const n = parseInt(nk[0], 10);

    const k = parseInt(nk[1], 10);

    const c = readLine().split(' ').map(cTemp => parseInt(cTemp, 10));

    let result = jumpingOnClouds(c, k);

    console.log(result)
    // ws.write(result + "\n");

    // ws.end();
}
