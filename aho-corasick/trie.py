from tabulate import tabulate

class Vertex:
  def __init__(self, parentIndex=None, char=None):
    self.char = char
    self.parentIndex = parentIndex
    self.child = {}
    self.suffix = None
    self.endlink = None
    self.pattern = None
    self.leaf = False
  
  def __str__(self):
    return [
      self.parentIndex,
      self.char,
      self.leaf,
      self.child,
      self.suffix,
      self.endlink
    ]

class Trie:
  def __init__(self):
    root = Vertex()

    self.trie = [root]
    self.size = 1
    self.root = 0

  def add(self, pattern):
    currentVertex = self.root

    for char in pattern:
      if not char in self.trie[currentVertex].child:
        self.trie.append(Vertex(
          currentVertex,
          char
        ))

        self.trie[currentVertex].child[char] = self.size
        self.size += 1

      currentVertex = self.trie[currentVertex].child[char]
    
    self.trie[currentVertex].leaf = True
    self.trie[currentVertex].pattern = pattern

  # find suffix and endlink
  def build_suffix_and_endlink(self):
    queue = [
      self.root
    ]

    while queue:
      s = queue.pop(0)

      # print("iterate over {}".format(self.trie[s].char))
      self.calculate_suffix(s)
      
      for key, pos in self.trie[s].child.items():
        vertex = self.trie[s].child[key]
        if vertex:
          queue.append(vertex)

  def set_root_and_children_sufix(self, vertex):
    vertex_node = self.trie[vertex]
    vertex_node.suffix = self.root
    vertex_node.endlink = self.root

    if vertex_node != self.root and vertex_node.leaf:
      vertex_node.endlink = vertex
      

  # vertex is a index
  def calculate_suffix(self, vertex):
    vertex_node = self.trie[vertex]

    # if root or his childrens just set sufix as root
    if vertex == self.root or vertex_node.parentIndex == self.root:
      self.set_root_and_children_sufix(vertex)
      return

    # backtracking
    parentSuffix = self.trie[vertex_node.parentIndex].suffix
    while True:
      sufixNode = self.trie[parentSuffix]
      # if found vertex as child of sufixNode, set them as sufix
      if self.trie[vertex].char in sufixNode.child:
        vertex_node.suffix = sufixNode.child[self.trie[vertex].char]
        break

      # if parent sufix is root, there are no matches
      if parentSuffix == self.root:
        vertex_node.suffix = self.root
        break
      
      # keep coming back, navigating by sufix
      parentSuffix = sufixNode.suffix
    
    if vertex_node.leaf:
      vertex_node.endlink = vertex
    else:
      vertex_node.endlink = self.trie[vertex_node.suffix].endlink

  def process(self, text):
    state = self.root

    result = []

    for char in text:
      while True:
        if char in self.trie[state].child:
          state = self.trie[state].child[char]
          break
      
        if state == self.root:
          break
        
        state = self.trie[state].suffix

      check_state = state

      while True:
        check_state = self.trie[check_state].endlink

        if check_state == self.root:
          break
        
        result.append(self.trie[check_state].pattern)

        check_state = self.trie[check_state].suffix
    
    return result

  def __str__(self):
    table = []
    headers = [
      "parentIndex",
      "char",
      "leaf",
      "child",
      "suffix",
      "endlink"
    ]
    for v in self.trie:
      table.append(v.__str__())

    return tabulate(table, headers, tablefmt="github")


t = Trie()

t.add('abd')
t.add('abc')
t.add('b')
t.add('bcd')
t.add('bd')

t.build_suffix_and_endlink()

print(t)

print(t.process('abcd'))





