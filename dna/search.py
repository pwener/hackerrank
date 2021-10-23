# binary search for index
def search(array, target, lowest=True):
  left = 0
  right = len(array) - 1

  while left <= right:
    mid = (left + right) // 2

    cond = False
    
    if lowest:
      cond = (array[mid] > target and array[mid-1] < target)
    else:
      cond = (array[mid] < target and array[mid+1] > target)

    if array[mid] == target or cond:
      return array[mid]

    if array[mid] < target:
      left = mid + 1
    else:
      right = mid - 1
  
  return array[0]

arr = [1, 3, 5, 10, 15, 20]

# if I want [2, 18]
print(search(arr, 2))
print(search(arr, 18, False))
