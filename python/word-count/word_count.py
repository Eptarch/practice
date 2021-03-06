import re
from collections import Counter


def count_words(sentence: str) -> dict[str: int]:
    return Counter(re.findall(r"[^\W_]+(?:'[a-z]+)?", sentence.lower()))
