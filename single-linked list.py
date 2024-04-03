class LinkedList:
    class Node:
        def __init__(self, data):
            self.data = data
            self.next = None

    head = None
    tail = None
    def append(self, element):
        new_node = self.Node(element)
        if self.head is None:
            self.head = new_node
            return
        if self.tail is None:
            self.head.next = new_node
            self.tail = new_node
            return
        self.tail.next = new_node
        self.tail = new_node

    def out(self):
        strs = ""
        node = self.head
        while node:
            strs += str(node.data) + " "
            node = node.next
        return strs.rstrip()


    def delete(self, value):
        if self.head.data == value:
            self.head = self.head.next
        node = self.head
        while node:
            if node.next and node.next.data == value:
                del node.next.data
                node.next = node.next.next
                return
            node = node.next


    def insert(self, element, value):
        node = self.head
        node_value = self.Node(value)
        while node:
            if node.data == element:
                temp = node.next
                node.next = node_value
                node_value.next = temp
                return
            node = node.next