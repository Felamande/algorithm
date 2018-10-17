class stack:
    def __init__(self):
        self.elems = []
        self.len = 0

    def push(self, v):
        self.elems[self.len] = v
        self.len += 1

    def pop(self):
        self.len -= 1
        return self.elems[self.len]

    def printstack(self):
        print(self.elems)


def main():
    s = stack()
    s.push(3)
    s.push(3)
    s.push(3)
    s.push(3)
    s.push(3)
    s.push(3)
    s.printstack()
