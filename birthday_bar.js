/*
 * Complete the 'birthday' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY s
 *  2. INTEGER d
 *  3. INTEGER m
 */

function birthday(s, d, m) {
    let result = 0;
    const lim = (s.length - m);

    console.log('iterate over 0|', lim);

    if (lim <= 0) {
        if (sum(s) == d) {
            return 1;
        }
    }

    for (let i = 0; i <= lim; i++) {
        const section = s.slice(i, i + m);

        if (sum(section) == d) {
            console.log('valid sec: ', section);
            result++;
        }
    }
    return result;
}

function sum(arr) {
    return arr.reduce((a, next) => a + next);
}

const arr = [2, 5, 1, 3, 4, 4, 3, 5, 1, 1, 2, 1, 4, 1, 3, 3, 4, 2, 1];

console.log(birthday(arr, 18, 7));