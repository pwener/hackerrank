# binary search for index
def search(array, target, lowest=True):
  left = 0
  right = len(array) - 1

  while left <= right:
    mid = (left + right) // 2

    cond = False
    
    if lowest:
      cond = (arr[mid] > target and arr[mid-1] < target)
    else:
      cond = (arr[mid] < target and arr[mid+1] > target)

    if arr[mid] == target or cond:
      return arr[mid]

    if arr[mid] < target:
      left = mid + 1
    else:
      right = mid - 1
  
  return array[0]

arr = [1, 2, 3, 5, 10, 15, 20]

# if I want [2, 18]
print(search(arr, 1))
print(search(arr, 18, False))
