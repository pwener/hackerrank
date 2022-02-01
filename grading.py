#!/bin/python3

import math
import os
import random
import re
import sys

def findNextDivByFive(number):
  count = 0
  while True:
    current = number + count

    if current % 5 == 0:
      return current
    else:
      count += 1

def diffCheck(number):
  nextDiv = findNextDivByFive(number)

  return (nextDiv - number) < 3

def gradingStudents(grades):
    res = []
    for grade in grades:
      if grade < 38:
        res.append(grade)
        continue

      if diffCheck(grade):
        res.append(findNextDivByFive(grade))
      else:
        res.append(grade)
            
    return res

print(gradingStudents([37, 38]))

# if __name__ == '__main__':
#     fptr = open(os.environ['OUTPUT_PATH'], 'w')

#     grades_count = int(input().strip())

#     grades = []

#     for _ in range(grades_count):
#         grades_item = int(input().strip())
#         grades.append(grades_item)

#     result = gradingStudents(grades)

#     fptr.write('\n'.join(map(str, result)))
#     fptr.write('\n')

#     fptr.close()
