/*
 * Complete the 'getTotalX' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY a
 *  2. INTEGER_ARRAY b
 */
function getTotalX(a, b) {
    const lowestCommonMultiple = lcm(a);
    const maxCommon = b.reduce((acc, next) => gdc(acc, next));

    let count = 1;
    let score = 0;

    while ((lowestCommonMultiple * count) <= maxCommon) {
        if ((maxCommon % (lowestCommonMultiple * count)) === 0) {
            score++;
        }
        count++;
    }

    return score;
}

function gdc(a, b) {
    if (b == 0) {
        return a
    }

    return gdc(b, a % b)
}

function lcm(values) {
    if (values.length === 1) {
        return values[0];
    }

    return values.reduce((acc, next) => {
        return acc * next / gdc(acc, next);
    })
}

const a = [2, 3, 6];
const b = [42, 84];

console.log(getTotalX(a, b))