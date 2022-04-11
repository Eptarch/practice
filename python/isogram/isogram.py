def is_isogram(s: str) -> bool:
    s = s.lower().replace(" ", "").replace("-", "")
    return len(s) == len(set(s))
