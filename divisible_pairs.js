/*
 * Complete the 'divisibleSumPairs' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. INTEGER k
 *  3. INTEGER_ARRAY ar
 */

function divisibleSumPairs(n, k, ar) {
    const arr = [...ar];
    let match = 0;

    while(arr.length > 0) {
        const curr = arr[0];
        for (let i = 1; i < arr.length; i++) {
            if ((curr + arr[i]) % k === 0) {
                match++;
            }
        }
        arr.shift();
    }

    return match;
}


console.log(divisibleSumPairs(6, 3, [1, 3, 2, 6, 1, 2]))