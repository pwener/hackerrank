/*
 * Complete the 'formingMagicSquare' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts 2D_INTEGER_ARRAY s as parameter.
 */

function formingMagicSquare(s = [], costs = []) {
    const sum = sumOf(s.flatMap(a => a));

    let target = sum;

    while(target % 3 !== 0) {
        target += 1;
    }

    target = target / 3;

    const hashSquare = new Map();

    s.forEach((line, i) => {
        line.forEach((el, j) => {
            hashSquare.set(`s${i}${j}`, el);
        })
    });

    const [lines, cols, diags] = findRelation(hashSquare);

    if (isMagicSquare(hashSquare, target)) {
        console.log('MAGIC!!!');
        return sumOf(costs);
    }

    for (el of hashSquare.keys()) {
        const i = parseInt(el[1]);

        const linesV = valuesOf(lines[i], hashSquare);

        if (sumOf(linesV) === target) {
            continue;
        }

        const j = parseInt(el[2]);

        const colsV = valuesOf(cols[j], hashSquare);

        if (sumOf(colsV) === target) {
            continue;
        }

        if ((i + j) % 2 === 0) {
            const d1 = valuesOf(diags[0], hashSquare);
            const d2 = valuesOf(diags[1], hashSquare);
            const isD1 = i == j && sumOf(d1) == target;
            const skipByD2 = sumOf(d2) == target;

            if (isD1 || skipByD2) {
                console.log('skip by d', isD1, skipByD2);
                continue;
            }
        }

        console.log('changing: ', el);

        const l = lines.filter(line => line.includes(el))[0];
        const sumOfLine = sumOf(valuesOf(l, hashSquare));

        console.log('sum: ', l);
        const diff = target - sumOfLine;

        costs.push(Math.abs(diff));

        const value = hashSquare.get(el);

        console.log('value:', value);
        console.log('diff:', diff);
        hashSquare.set(el, value + diff);

        break;
    }

    const result = Array.from(hashSquare.values()).reduce((r, _, index, arr) => {
        if (index % 3 === 0) {
            r.push(arr.slice(index, index + 3));
        }
        return r;
    }, []);

    return formingMagicSquare(result, costs);
}



function findRelation(hashMap) {
    const keys = Array.from(hashMap.keys());

    const lines = keys.reduce((r, _, index, arr) => {
        if (index % 3 === 0) {
            r.push(arr.slice(index, index + 3));
        }
        return r;
    }, []);

    const columns = keys.reduce((acc, current) => {
        const index = parseInt(current.slice(-1));
        acc[index].push(current);
        return acc;
    }, [[],[],[]]);

    const diagonals = keys.reduce((acc, current) => {
        const row = parseInt(current[1]);
        const col = parseInt(current[2]);

        if (row == col) {
            acc[0].push(current);
        }

        if (row + col == 2) {
            acc[1].push(current);
        }

        return acc;
    }, [[], []]);

    return [lines, columns, diagonals];
};

function sumOf(arr) {
   return arr.reduce((acc, a) => acc + a, 0);
}

function valuesOf(arr, hashMap) {
    return arr.map(k => hashMap.get(k));
}

function isMagicSquare(hashSquare, target) {
    const [lines, cols, diags] = findRelation(hashSquare);

    const spread = [...lines, ...cols, ...diags].map(arr => valuesOf(arr, hashSquare))
    
    console.log(spread);

    return spread.every((a) => sumOf(a) === target);
}

const s = [[4, 8, 2], [4, 5, 7], [6, 1, 6]];
const s2 = [ [ 4, 9, 2 ], [ 4, 5, 7 ], [ 6, 1, 6 ] ];
const s3 = [ [ 4, 9, 2 ], [ 3, 5, 7 ], [ 6, 1, 6 ] ];
const s4 = [ [ 4, 9, 2 ], [ 3, 5, 7 ], [ 8, 1, 6 ] ];

console.log(formingMagicSquare(s));