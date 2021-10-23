from bisect import bisect_left, bisect_right, bisect

# binary search for index
def search_left(array, target):
  i = bisect_left(array, target)
  return array[i]

# binary search for index
def search_right(array, target):
  i = bisect_right(array, target)
  return array[i-1]

arr = [1, 3, 5, 10, 15, 20, 22]

# if I want [2, 18]
print(search_left(arr, 3))
print(search_right(arr, 19))
