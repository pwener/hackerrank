# Links

https://www.toptal.com/algorithms/aho-corasick-algorithm
https://en.wikipedia.org/wiki/Aho%E2%80%93Corasick_algorithm
https://en.wikipedia.org/wiki/Breadth-first_search
https://www.hackerearth.com/practice/algorithms/graphs/breadth-first-search/tutorial/

# Entendendo Vertex

Dado que vertex tem as seguintes propriedades:

- size is the lenght of aho trie
- leaf flag to sinalize that some word ends here
- parent link to the parent vertex
- sufix link to sufix of current vertex
- endWordLink link to the leaf vertex of the maximum-length word

# Adicionando os padrões

Dado aho como:

```
- trie is  the list of vertex
- wordsLen is the list of words lenght
- size is current size of aho structure
- root is the first element
```

Dado um padrão(pattern) de id único(w), temos:

```
- init currentVertex as root
- for each p in pattern
  - if p are not settled as children of currentVertex:
    - append to aho trie a new vertex of p
    - set the children of currentVertex as aho trie size
    - increment aho trie size
  - set new currentVertex as value of children p of aho trie of currentVertex

- last p will sets leaf to true
- append pattern lenght to wordsLen
```

# Preparando a estrutura aho-corasick usando BFS

Baseado em https://en.wikipedia.org/wiki/Breadth-first_search

```
- Start a queue of int with root
- while queue size > 0
  - dequeue and atributes element as currentVertex
  - Calculate Sufix Link of aho[currentVertex]
  - Enqueue all childrens of aho[currentVertex] to queue

```

# Calculando o Sufix Link

```
- For each vertex value of Vertex

  - if vertex is root:
    - set sufixLink as root
    - set endWordLink as root

  - if vertex parent is root:
    - set sufixLink as root
    - if leaf:
      - set endWordLink as vertex
    - else:
      - set endWordLink root
      - return

  - while:
    - set current as sufixLink of parent vertex
    - set chVertex as parent char
    - if children of current are settled with chVertex:
      - set sufixLink of vertex as sufixLink of ahio[current]
      - break
    - if current is root:
      - set sufixLink as root
      - break
    - else:
      current is sufixLink of ahio[current]

  - if vertex is leaf:
    - set endWordLink as vertex
  - else
    - set endWordLink as sufixLink of ahio[vertex.sufixLink]

```
