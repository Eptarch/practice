import heapq


def latest(scores: list[int]) -> int:
    return scores[-1]


def personal_best(scores: list[int]) -> int:
    return max(scores)


def personal_top_three(scores: list[int]) -> [int]:
    return heapq.nlargest(3, scores)

