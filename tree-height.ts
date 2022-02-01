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

class NodeTree {
  public value: number;
  public level: number;
  public left?: NodeTree;
  public right?: NodeTree;

  constructor(value: number, level: number, left?: NodeTree, right?: NodeTree) {
    this.value = value;
    this.left = left;
    this.right = right;
    this.level = level;
  }

  static create(value: number, level: number) {
    return new NodeTree(value, level);
  }
}

class BinaryTree {
  private root: NodeTree | null = null;
  public length = 0;

  public insert(value: number) {
    if (this.root == null) {
      this.root = NodeTree.create(value, 0);
    } else {
      let current = this.root;

      while (true) {
        if (value > current.value) {
          if (!current.right) {
            current.right = NodeTree.create(value, current.level + 1);
            current = current.right;
            break;
          } else {
            current = current.right;
          }
        }

        if (value < current.value) {
          if (!current.left) {
            current.left = NodeTree.create(value, current.level + 1);
            current = current.left;
            break;
          } else {
            current = current.left;
          }
        }
      }

      if (current.level > this.length) {
        this.length = current.level;
      }
    }
  }
}

function findHeight(values: number[]) {
  const tree = new BinaryTree();

  values.forEach((v) => {
    tree.insert(v);
  });

  return tree.length;
}

// expect 3
function main() {
  readLine();

  const input: number[] = readLine()
    .replace(/\s+$/g, '')
    .split(' ')
    .map((arrTemp) => parseInt(arrTemp, 10));

  const len = findHeight(input);

  console.log(len);
}
