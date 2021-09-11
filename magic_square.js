/*
 * Complete the 'formingMagicSquare' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts 2D_INTEGER_ARRAY s as parameter.
 */

function formingMagicSquare(s = []) {
  const triple = [
    [4, 3, 8],
    [8, 1, 6],
    [6, 7, 2],
    [2, 9, 4],
  ];

  const inverseTriple = [...triple].reverse().map(el => [...el].reverse());

  const arr = [
    s[0],
    [s[0][2], s[1][2], s[2][2]],
    [...s[2]].reverse(),
    [s[2][0], s[1][0], s[0][0]]
  ];

  // just to be replaced in begin
  let lesser = 45;

  for (let i = 0; i < 8; i++) {
    const combination = i <= 3 ? triple : inverseTriple;

    if (i > 0) {
      combination.push(combination.shift());
    }

    const cost = findCost(arr, combination);

    if (cost < lesser) {
      lesser = cost;
    }
  }

  const midCostReplace = Math.abs(s[1][1] - 5);

  return lesser + midCostReplace;
};

function findCost(original, changed) {
  let acc = {
    cost: 0,
    elements: []
  };

  for (let i = 0; i < 4; i++) {
    const originalConj = original[i];
    const changedConj = changed[i];
    for (let j = 0; j < 3; j++) {
      const choosedNumber = changedConj[j];

      if (!acc.elements.includes(choosedNumber)) {
        acc.cost = acc.cost + Math.abs(originalConj[j] - choosedNumber);
        acc.elements.push(choosedNumber);
      }

    }
  }

  return acc.cost;
}


const s = [[4, 9, 2], [3, 5, 7], [8, 1, 5]];

console.log(formingMagicSquare(s));
