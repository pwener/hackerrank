/*
 * Complete the 'countApplesAndOranges' function below.
 *
 * The function accepts following parameters:
 *  1. INTEGER s
 *  2. INTEGER t
 *  3. INTEGER a
 *  4. INTEGER b
 *  5. INTEGER_ARRAY apples
 *  6. INTEGER_ARRAY oranges
 */

function countApplesAndOranges(s, t, a, b, apples, oranges) {
    const isBetween =  (n) => {
        if (n >= s && n <= t) {
            return 1;
        }
        return 0;
    }

    const applesCount = apples.reduce((acc, apple) => {
        return acc + isBetween(a + apple);
    }, 0);

    console.log(applesCount, '\n');

    const orangesCount = oranges.reduce((acc, orange) => {
        return acc + isBetween(b + orange);
    }, 0);

    console.log(orangesCount);
}

console.log(countApplesAndOranges())