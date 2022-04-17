class Luhn:
    def __init__(self, card_num: str) -> None:
        self.data = card_num

    def valid(self) -> bool:
        double = True
        prep = []
        for r in self.data[::-1]:
            if r.isdigit():
                double = not double
                d = int(r)
                if double:
                    d *= 2
                    if d > 9:
                        d -= 9
                prep.append(d)
            elif not r.isspace():
                return False
        if len(prep) < 2:
            return False
        return sum(v for v in prep) % 10 == 0
