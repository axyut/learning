const std = @import("std");
const linkedList = @import("./linkedlist.zig").LinkedList;

pub fn main() !void {
    // Prints to stderr (it's a shortcut based on `std.io.getStdErr()`)
    std.debug.print("All your {s} are belong to us.\n", .{"codebase"});

    // stdout is for the actual output of your application, for example if you
    // are implementing gzip, then only the compressed bytes should be sent to
    // stdout, not any debugging messages.
    const stdout_file = std.io.getStdOut().writer();
    var bw = std.io.bufferedWriter(stdout_file);
    const stdout = bw.writer();

    try stdout.print("Run `zig build test` to run the tests.\n", .{});

    try bw.flush(); // don't forget to flush!

    // Assign an arena allocator for our linked list to use for creating nodes
    var arena = std.heap.ArenaAllocator.init(std.heap.page_allocator);
    const allocator = arena.allocator();

    // Don't forget to free the memory on exit!
    defer arena.deinit();

    // Declare our linked list and add a few nodes
    var u32LinkedList = linkedList(u32).new(allocator);
    try u32LinkedList.insert(2);
    try u32LinkedList.insert(3);
    try u32LinkedList.insert(1);

    // Finally, traverse the list with the output:
    //    1
    //    3
    //    2
    u32LinkedList.traverse();
}

test "simple test" {
    var list = std.ArrayList(i32).init(std.testing.allocator);
    defer list.deinit(); // try commenting this out and see if zig detects the memory leak!
    try list.append(42);
    try std.testing.expectEqual(@as(i32, 42), list.pop());
}
