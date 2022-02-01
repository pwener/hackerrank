'use strict';

import { count } from 'console';

class Nozim {
  public value: number;
  public left?: Nozim;
  public right?: Nozim;

  constructor(value: number) {
    this.value = value;
  }

  static create(value: number) {
    return new Nozim(value);
  }
}

class Arvere {
  private root: Nozim | null = null;

  public insert(value: number) {
    if (this.root == null) {
      this.root = Nozim.create(value);
    } else {
      let current = this.root;

      while (true) {
        if (value > current.value) {
          if (current.right) {
            current = current.right;
          } else {
            current.right = Nozim.create(value);
            break;
          }
        }

        if (value < current.value) {
          if (current.left) {
            current = current.left;
          } else {
            current.left = Nozim.create(value);
            break;
          }
        }
      }
    }
  }

  public getTransversal() {
    const nodes = [];

    const queue = [this.root];

    while (queue.length !== 0) {
      const curr = queue.shift();

      nodes.push(curr?.value);

      if (curr?.left != null) {
        queue.push(curr?.left);
      }

      if (curr?.right != null) {
        queue.push(curr?.right);
      }
    }

    return nodes;
  }
}

function levelOrder(values: number[]) {
  const tree = new Arvere();

  values.forEach((v) => tree.insert(v));

  return tree.getTransversal().join(' ');
}

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

function main() {
  readLine();

  const input: number[] = readLine()
    .replace(/\s+$/g, '')
    .split(' ')
    .map((arrTemp) => parseInt(arrTemp, 10));

  console.log(levelOrder(input));
}
