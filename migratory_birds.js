const fs = require('fs');

function migratoryBirds(arr) {
    const countMap = new Map();
    const max = {value: 0, count: 0};

    arr.forEach((el) => {
        const count = countMap.get(el);
        if (count) {
            if (count > max.count || (count == max.count && el < max.value)) {
                max.value = el;
                max.count = count;
            }
            countMap.set(
                el,
                count + 1
            );
        } else {
            countMap.set(
                el, 1
            )
        }
    });
    return max.value;
}

const input01 = [1, 4, 4, 4, 5, 3];

const input02 = [1, 2, 3, 4, 5, 4, 3, 2, 1, 3, 4];

const input04 = fs.readFileSync('input04.txt', 'utf8').split(' ').map(e => parseInt(e, 10));

// console.log(input04);

console.log(migratoryBirds(input04));
