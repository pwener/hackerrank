/*
 * Complete the 'formingMagicSquare' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts 2D_INTEGER_ARRAY s as parameter.
 */

function formingMagicSquare(s = []) {
    const sums = [rowsSum(s), colsSum(s), diagSum(s)].flat();

    const common = sums.reduce((acc, sum) => {
        acc[sum] = acc[sum] ? acc[sum] + 1 : 1;
        return acc;
    }, {});

    const mostCommon = Object.keys(common).reduce((acc, i) => {
        if (common[i] > acc.count) {
            return {
                sum: i,
                count: common[i]
            };
        }
        return acc;
    }, {sum: 0, count: 0});

    // console.log('common:', mostCommon.sum);

    const hashSquare = new Map();

    s.forEach((line, i) => {
        line.forEach((el, j) => {
            hashSquare.set(`s${i}${j}`, el);
        })
    });

    const findRelation = (mc, hashMap) => {
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

        const elements = [...lines, ...columns, ...diagonals];

        console.log('common:', mc);

        const exceptions = elements.filter(conj => {
            const sum = conj.reduce((acc, index) => acc + hashMap.get(index), 0);

            console.log(`sum is ${sum}| ${sum} !== ${mc}`, sum != mc);
            
            return mc != sum;
        });

        return exceptions;
    }

    // console.log('common:', mostCommon.sum);
    
    return findRelation(mostCommon.sum, hashSquare);
}

function rowsSum(s) {
    return s.map(line => line.reduce((acc, el) => acc + el));
}

function colsSum(s) {
    return s.reduce((acc, line) => {
        const [i, j, k] = acc;
        const [c1, c2, c3] = line;
        return [
            i + c1,
            j + c2,
            k + c3,
        ];
    }, [0, 0, 0]);
}

function diagSum(s) {
    return [
        s[0][0] + s[1][1] + s[2][2],
        s[0][2] + s[1][1] + s[2][0],
    ]
}

function validate(s) {
    const magicConstant = s[0].reduce((acc, el) => acc + el);

    const rowsValidator = () => {
        return rowsSum(s).filter(e => e === magicConstant).length === 3;
    }

    const colValidator = () => {
        return colsSum(s).filter(e => e === magicConstant).length === 3;
    }

    const diagValidator = () => {
        return diagSum(s).filter(e => e === magicConstant).length === 2;
    }

    return rowsValidator() && colValidator() && diagValidator();
}


const s = [[4, 8, 2], [4, 5, 7], [6, 1, 6]]
// const s2 = [[4, 9, 2], [3, 5, 7], [8, 1, 6]]

console.log(formingMagicSquare(s));
