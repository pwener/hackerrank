class Vertex:
  def __init__(self, parentIndex=None, char=None):
    self.char = char
    self.parentIndex = parentIndex
    self.children = {}
    self.leaf = False
  
  def __str__(self):
    return "{}\t |{} \t|{}  | {} ".format(
      self.parentIndex,
      self.char,
      self.leaf,
      self.children
    )

class Trie:
  def __init__(self):
    root = Vertex()

    self.trie = [root]
    self.size = 1
    self.root = 0

  def add(self, pattern):
    currentVertex = self.root

    for char in pattern:
      if not char in self.trie[currentVertex].children:
        self.trie.append(Vertex(
          currentVertex,
          char
        ))

        self.trie[currentVertex].children[char] = self.size
        self.size += 1

      currentVertex = self.trie[currentVertex].children[char]
    
    self.trie[currentVertex].leaf = True

  # find pattern using bfs.
  def bfs(self, index):
    visited = [
      self.trie[index].char
    ]
    queue = [
      self.trie[index]
    ]

    while queue:
      s = queue.pop(0)
      
      # print("mark {} at {}".format(s.char, s.parentIndex), end=" ")

      for char, pos in s.children.items():
        if char not in visited:
          vertex = self.trie[pos]
          visited.append(char)
          queue.append(vertex)
          print("visit {} at {}".format(char, pos), end="\n")

          
  def __str__(self):
    s = ""
    for v in self.trie:
      s += v.__str__()
      s += "\n"
    return s


t = Trie()

t.add('abc')
t.add('af')
t.add('abd')

# print("parent\t |char  | leaf  | children ".format())
# print(t)

t.bfs(1)





