"""
Find the elf with the most calories and return it given the text document input
"""
import os

dir_path = os.path.dirname(os.path.realpath(__file__))
lines = []
maxCalories = 0
currentLine = 0

with open(f"{dir_path}/input.txt") as plainInput:
    for line in plainInput.readlines():
        a = line.strip()
        lines.append(a)

while currentLine < len(lines):
    elfCalories = 0

    while lines[currentLine] != "" and currentLine + 1 < len(lines):
        elfCalories += int(lines[currentLine])
        currentLine += 1

    maxCalories = max(elfCalories, maxCalories)
    currentLine += 1

print(maxCalories)
