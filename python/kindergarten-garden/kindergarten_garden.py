plants = {
    "V": "Violets",
    "R": "Radishes",
    "G": "Grass",
    "C": "Clover",
}


class Garden:
    students = (
        "Alice",
        "Bob",
        "Charlie",
        "David",
        "Eve",
        "Fred",
        "Ginny",
        "Harriet",
        "Ileana",
        "Joseph",
        "Kincaid",
        "Larry",
    )

    def __init__(self, diagram: str, students: list[str] = None) -> None:
        if students is not None:
            self.students = sorted(students)
        self.diagram = diagram.splitlines()

    def plants(self, student: str) -> list[str]:
        ind = self.students.index(student) * 2
        return [plants[pl[i]] for pl in self.diagram for i in (ind, ind + 1)]
