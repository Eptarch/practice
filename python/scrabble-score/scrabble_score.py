SCORES = {
    1: ["A", "E", "I", "O", "U", "L", "N", "R", "S", "T"],
    2: ["D", "G"],
    3: ["B", "C", "M", "P"],
    4: ["F", "H", "V", "W", "Y"],
    5: ["K"],
    8: ["J", "X"],
    10: ["Q", "Z"]
}


def score(word: str) -> int:
    return sum(
        [
            score for score, letters in SCORES.items()
            for letter in word.upper() if letter in letters
        ]
    )
