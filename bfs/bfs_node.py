class Node:
  def __init__(self, value):
    self.value = value
    self.children = []

  def linkTo(self, values):
    for value in values:
      self.children.append(value)

A = Node('A')
B = Node('B')
C = Node('C')
D = Node('D')
E = Node('E')
F = Node('F')

A.linkTo([B, C])
B.linkTo([D, E])
C.linkTo([F])

graph = [A, B, C, D, E, F]

visited = []
queue = []

def bfs(visited, graph, node):
  visited.append(node.value)
  queue.append(node)

  while queue:
    s = queue.pop(0)
    print(s.value, end=" ")

    for neighbour in s.children:
      if neighbour.value not in visited:
        visited.append(neighbour.value)
        queue.append(neighbour)
  
  print("")

bfs(visited, graph, graph[1])

