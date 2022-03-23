'use strict';

import { WriteStream, createWriteStream } from 'fs';
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

class ContactNode {
  public childs: Map<string, ContactNode>;
  public ocurrences: number;

  constructor() {
    this.childs = new Map();
    this.ocurrences = 0;
  }
}

class TreeContact {
  public root: ContactNode = new ContactNode();

  public add(word: string) {
    let current = this.root;

    for (let count = 0; count < word.length; count++) {
      let node = current.childs.get(word[count]);

      if (node == null) {
        node = new ContactNode();
        current.childs.set(word[count], node);
      }

      node.ocurrences += 1;

      current = node;
    }
  }

  public findOcurrences(substr: string) {
    let current = this.root;

    for (let count = 0; count < substr.length; count++) {
      const child = current.childs.get(substr[count]);

      if (!child) {
        return 0;
      }

      current = child;
    }

    return current.ocurrences;
  }
}

function contacts(queries: string[][]): number[] {
  const trie = new TreeContact();
  const ocurrences = [];

  for (const q of queries) {
    const command = q[0];
    const arg = q[1];

    if (command == 'add') {
      trie.add(arg);
    }

    if (command == 'find') {
      ocurrences.push(trie.findOcurrences(arg));
    }
  }

  return ocurrences;
}

function main() {
  const queriesRows: number = parseInt(readLine().trim(), 10);

  let queries: string[][] = Array(queriesRows);

  for (let i: number = 0; i < queriesRows; i++) {
    queries[i] = readLine().replace(/\s+$/g, '').split(' ');
  }

  const result: number[] = contacts(queries);

  console.log(result);
}
