

def moveCard(grid, startRow, startCol, commands):
    rows, cols = len(grid), len(grid[0])
    currentRow, currentCol = startRow, startCol
    
    for command in commands:
        if command == "up" and currentRow > 0:
            currentRow -= 1
        elif command == "down" and currentRow < rows - 1:
            currentRow += 1
        elif command == "left" and currentCol > 0:
            currentCol -= 1
        elif command == "right" and currentCol < cols - 1:
            currentCol += 1
    
    return currentRow, currentCol

grid = [
    [1, 2, 3],
    [4, 5, 6],
    [7, 8, 9]
]
startRow = 1
startCol = 1
commands = ["up", "left", "down", "right", "up"]
print(moveCard(grid, startRow, startCol, commands))  # Output: (2, 1)
print(grid)