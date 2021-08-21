/*
 * Complete the 'kangaroo' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts following parameters:
 *  1. INTEGER x1
 *  2. INTEGER v1
 *  3. INTEGER x2
 *  4. INTEGER v2
 */

function kangaroo(x1, v1, x2, v2) {
    const r = (x1 - x2) / (v2 - v1)

    console.log(r);
    
    return r > 0 && r % 1 == 0 ? 'YES' : 'NO'
}

console.log(kangaroo(21, 6, 47, 3));

// console.log(kangaroo(0, 2, 5, 3))

// console.log(kangaroo(0, 3, 4, 2))