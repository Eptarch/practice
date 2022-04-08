class Matrix:
    def __init__(self, matrix_string: str) -> None:
        self.content = [
            [int(y) for y in x.split(" ")] for x in matrix_string.split("\n")
        ]

    def row(self, index: int) -> [int]:
        return self.content[index - 1]

    def column(self, index: int) -> [int]:
        return [x[index - 1] for x in self.content]

