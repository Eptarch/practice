from typing import Optional


def two_fer(name: Optional[str] = None) -> str:
    return f"One for {name if name else 'you'}, one for me."
