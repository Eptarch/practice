def convert(number: int) -> str:
    t = {3: "Pling", 5: "Plang", 7: "Plong"}
    return "".join(
        drop for divisor, drop in t.items() if number % divisor == 0
    ) or str(number)
