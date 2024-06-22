function moveCard(grid, startRow, startCol, commands) {
    var rows = grid.length;
    var cols = grid[0].length;
    var currentRow = startRow;
    var currentCol = startCol;
    for (var _i = 0, commands_1 = commands; _i < commands_1.length; _i++) {
        var command = commands_1[_i];
        if (command === "up" && currentRow > 0) {
            currentRow -= 1;
        } else if (command === "down" && currentRow < rows - 1) {
            currentRow += 1;
        } else if (command === "left" && currentCol > 0) {
            currentCol -= 1;
        } else if (command === "right" && currentCol < cols - 1) {
            currentCol += 1;
        }
    }
    return [currentRow, currentCol];
}

var grid = [
    [1, 2, 3],
    [4, 5, 6],
    [7, 8, 9],
];
var startRow = 1;
var startCol = 1;
var commands = ["up", "left", "down", "right", "down"];
console.log(moveCard(grid, startRow, startCol, commands)); // Output: [2, 1]
