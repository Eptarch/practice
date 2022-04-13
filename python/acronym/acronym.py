import re


def abbreviate(words: str) -> str:
    return "".join(
        [word[0] for word in re.findall(r"[A-Za-z]+(?:'[a-z]+)?", words)]
    ).upper()
