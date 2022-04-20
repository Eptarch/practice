class Clock:
    def __init__(self, hour, minute: int) -> None:
        self.m = (hour * 60 + minute) % 1440

    def __repr__(self) -> str:
        return f"Clock({self.m // 60}, {self.m % 60})"

    def __str__(self) -> str:
        return f"{self.m // 60:02d}:{self.m % 60:02d}"

    def __eq__(self, other: "Clock") -> bool:
        return self.m == other.m

    def __add__(self, m: int) -> "Clock":
        return Clock(0, self.m + m)

    def __sub__(self, m: int) -> "Clock":
        return Clock(0, self.m - m)
