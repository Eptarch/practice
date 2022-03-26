def eat_ghost(power_pellet_active, touching_ghost: bool) -> bool:
    """
    Return wether Pac-man is being able to eat a ghost.

    :param power_pellet_active: bool - does the player have an active power pellet?
    :param touching_ghost:  bool - is the player touching a ghost?
    :return: bool - is pacman able to eat a ghost?
    """
    return power_pellet_active and touching_ghost


def score(touching_power_pellet, touching_dot: bool) -> bool:
    """
    Return wether Pac-man scores.

    :param touching_power_pellet: bool - does the player have an active power pellet?
    :param touching_dot:  bool - is the player touching a dot?
    :return: bool - Pac-man scores?
    """
    return touching_power_pellet or touching_dot


def lose(power_pellet_active, touching_ghost):
    """
    Return wether Pac-man loses.

    :param power_pellet_active: bool - does the player have an active power pellet?
    :param touching_ghost: bool - is the player touching a ghost?
    :return: bool - Pac-man loses?
    """
    return touching_ghost and not power_pellet_active


def win(has_eaten_all_dots, power_pellet_active, touching_ghost):
    """
    Return wether Pac-man wins.

    :param has_eaten_all_dots: bool - has the player "eaten" all the dots?
    :param power_pellet_active: bool - does the player have an active power pellet?
    :param touching_ghost:  bool - is the player touching a ghost?
    :return: bool - Pac-man wins?
    """
    return has_eaten_all_dots and not lose(power_pellet_active, touching_ghost)
