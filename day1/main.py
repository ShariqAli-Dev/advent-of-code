"""
Find the elf with the most calories and return it given the text document input
"""
import os

dir_path = os.path.dirname(os.path.realpath(__file__))

with open(f"{dir_path}/input.txt") as plainInput:
    lines = [line.strip() for line in plainInput.readlines()]

l = 0
maxCalories = 0

while l < len(lines):
    elfCalories = 0

    while l + 1 != len(lines) and lines[l] != "":
        elfCalories += int(lines[l])
        l += 1

    maxCalories = max(maxCalories, elfCalories)
    l += 1

print(f"max calories, {maxCalories}")
