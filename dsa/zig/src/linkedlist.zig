const std = @import("std");
const pageAllocator = std.heap.page_allocator;
const testing = std.testing;

pub fn LinkedList(comptime T: type) type {
    return struct {
        const Self = @This();

        // 1. Define a node type
        const Node = struct {
            value: T,
            next: ?*Node,
        };

        // 2a. Define the linked list properties
        // There should be three: head, length, and allocator
        head: ?*Node,
        length: u32,
        allocator: std.mem.Allocator,

        // 2b. Add a constructor/initializer for our linked list
        pub fn new(allocator: std.mem.Allocator) Self {
            return .{ .length = 0, .head = null, .allocator = allocator };
        }

        // 3. Define an insert method that takes a generic type value
        pub fn insert(self: *Self, value: T) !void {
            // Allocate the memory and create a `Node` for us to use
            var newNode = try self.allocator.create(Node);

            // Next, set the node value and point it's next value to the current head
            const currentHead = self.head;
            newNode.value = value;
            newNode.next = currentHead;

            // Finally, repoint our head to the new node and increment the count
            self.head = newNode;
            self.length += 1;
        }

        // 4. Define a pop method that removes the last inserted node
        pub fn pop(self: *Self) ?T {
            // If we don't have a head, there's no value to pop!
            if (self.head == null) {
                return null;
            }

            // Grab a few temporary values of the current head
            const currentHead = self.head;
            const updatedHead = self.head.?.next;

            // Update head and decrement the length now that we're freeing ourselves of a node
            self.head = updatedHead;
            self.length -= 1;

            return currentHead.?.value;
        }

        // 5. Define a traverse method, printing all the values
        pub fn traverse(self: *Self) void {
            // If we don't have a head, there's nothing traverse!
            if (self.head == null) {
                return;
            }

            // We'll walk our linked list as long as there's a next node available
            var currentNode = self.head;

            while (currentNode != null) : (currentNode = currentNode.?.next) {
                std.log.info("value {}", .{currentNode.?.value});
            }
        }
    };
}
