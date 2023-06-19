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
 * Complete the 'appendAndDelete' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts following parameters:
 *  1. STRING s
 *  2. STRING t
 *  3. INTEGER k
 */

function appendAndDelete(s, t, k) {
    let count = s.length + t.length;

    if (k > count) {
        return "Yes"
    }

    for (let i = 0, l = Math.min(s.length, t.length); i < l; i++) {
        if (s[i] !== t[i]) {
            break;
        }
        count -= 2;
    }

    if (count > k) {
        return 'No';
    }

    if ((count - k) % 2 == 0) {
        return 'Yes';
    }

    return 'No';
}

function main() {
    const s = readLine();

    const t = readLine();

    const k = parseInt(readLine().trim(), 10);

    const result = appendAndDelete(s, t, k);

    console.log(result);
}
