/*
 * Complete the 'pickingNumbers' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY a as parameter.
 */
function pickingNumbers(a = []) {
    const groups = a.sort().reduce((acc, curr, index) => {
        const next = a[(index + 1) % a.length];

        if (Math.abs(curr - next) <= 1) {
            acc.push([curr, next]);
        }

        return acc;
    }, []);

    const count = {};

    a.forEach((el) => {        
        count[el] = count[el] ? count[el] + 1 : 1;
    });

    return groups.reduce((acc, curr) => {
        const [a, b] = curr;
        const sum = a != b ? count[a] + count[b] : count[a];
        // console.log(`count of a=${a} and b=${b} is `, sum);
        if ( sum > acc) {
            return sum;
        }
        return acc;
    }, 0);
}

const a = [4, 6, 5, 3, 3, 1];
const b = [1, 2, 2, 3, 1, 2];

console.log(pickingNumbers(b));