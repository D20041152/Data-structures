class Tree:
    class Node:
        left = right = None

        def __init__(self, data) -> None:
            self.data = data

        def __str__(self) -> str:
            return str(self.data)

    HEAD = None

    def balance(self, mas, node):
        n = len(mas) // 2
        root = node
        if not self.HEAD:
            root = self.Node(node)
            self.HEAD = root
        lefts = mas[:n]
        rights = mas[n + 1:]
        root.left = self.Node(lefts[len(lefts) // 2])
        root.right = self.Node(rights[len(rights) // 2])
        if n <= 1:
            return
        else:
            self.balance(lefts, root.left)
            self.balance(rights, root.right)
        return
