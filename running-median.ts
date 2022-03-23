'use strict';

process.stdin.resume();
process.stdin.setEncoding('utf-8');

let inputString: string = '';
let inputLines: string[] = [];
let currentLine: number = 0;

process.stdin.on('data', function (inputStdin: string): void {
  inputString += inputStdin;
});

process.stdin.on('end', function (): void {
  inputLines = inputString.split('\n');
  inputString = '';

  main();
});

function readLine(): string {
  return inputLines[currentLine++];
}

class MedianHeap {
  public heap: Array<number>;

  constructor() {
    this.heap = [];
  }

  insert(value: number) {
    this.heap.push(value);

    let index = this.heap.length - 1; // 2

    const current = this.heap[index]; // 5

    while (index > 0) {
      let parentIndex = Math.floor((index - 1) / 2); // 0
      let parent = this.heap[parentIndex]; // 12

      if (parent > current) {
        this.heap[parentIndex] = current;
        this.heap[index] = parent;
        index = parentIndex;
      } else {
        break;
      }
    }
  }
}

/*
 * Complete the 'runningMedian' function below.
 *
 * The function is expected to return a DOUBLE_ARRAY.
 * The function accepts INTEGER_ARRAY a as parameter.
 */

function runningMedian(a: number[]): number[] {
  const results = [];
  const tree = new MedianHeap();

  for (let i = 0; i < a.length; i++) {
    tree.insert(a[i]);

    console.log(tree);
    console.log('\n');

    const mid = Math.floor(tree.heap.length / 2);

    if (tree.heap.length % 2 == 0) {
      results.push(tree.heap[mid]);
    } else {
      results.push((tree.heap[mid] + tree.heap[mid + 1]) / 2);
    }
  }

  return results;
}

function main() {
  const aCount: number = parseInt(readLine().trim(), 10);

  let a: number[] = [];

  for (let i: number = 0; i < aCount; i++) {
    const aItem: number = parseInt(readLine().trim(), 10);

    a.push(aItem);
  }

  const result: number[] = runningMedian(a);

  console.log(result);
}
