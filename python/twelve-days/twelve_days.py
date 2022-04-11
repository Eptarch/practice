class Verse:

    def __init__(self, ord, gift):
        self.ord = ord
        self.gift = gift


carol = [
    Verse("first", "a Partridge in a Pear Tree"),
    Verse("second", "two Turtle Doves"),
    Verse("third", "three French Hens"),
    Verse("fourth", "four Calling Birds"),
    Verse("fifth", "five Gold Rings"),
    Verse("sixth", "six Geese-a-Laying"),
    Verse("seventh", "seven Swans-a-Swimming"),
    Verse("eighth", "eight Maids-a-Milking"),
    Verse("ninth", "nine Ladies Dancing"),
    Verse("tenth", "ten Lords-a-Leaping"),
    Verse("eleventh", "eleven Pipers Piping"),
    Verse("twelfth", "twelve Drummers Drumming"),
]


def recite(start_verse, end_verse: int) -> str:
    return [verse(i) for i in range(start_verse, end_verse + 1)]


def verse(n: int) -> str:
    result = [
        f"On the {carol[n - 1].ord} day of Christmas my true love gave to me: "
    ]
    c = n
    while c:
        if n == 1:
            result.append(f"{carol[c - 1].gift}.")
        elif c == 1:
            result.append(f"and {carol[c - 1].gift}.")
        else:
            result.append(f"{carol[c - 1].gift}, ")
        c -= 1
    return "".join(result)
