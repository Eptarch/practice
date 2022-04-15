class School:
    def __init__(self) -> None:
        self.ns = dict()
        self.log = list()

    def add_student(self, name: str, grade: int) -> None:
        if name in self.ns:
            return self.log.append(False)
        self.ns[name] = grade
        self.log.append(True)

    def roster(self) -> list[str]:
        return [
            a for s in [
                self.grade(n) for n in set(self.ns.values())
            ] for a in s
        ]

    def grade(self, grade_number: int) -> list[str]:
        return sorted(n for n, g in self.ns.items() if g == grade_number)

    def added(self) -> list[bool]:
        return self.log
