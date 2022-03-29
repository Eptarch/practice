from typing import Union

def get_coordinate(record: tuple[str, str]) -> str:
    """

    :param record: tuple - a (treasure, coordinate) pair.
    :return: str - the extracted map coordinate.
    """

    return record[1]


def convert_coordinate(coordinate: str) -> tuple[str, str]:
    """

    :param coordinate: str - a string map coordinate
    :return:  tuple - the string coordinate seperated into its individual components.
    """
    
    return tuple(coordinate)


def compare_records(azara_record: tuple[str, str], rui_record: tuple[str, str, str]) -> bool:
    """

    :param azara_record: tuple - a (treasure, coordinate) pair.
    :param rui_record: tuple - a (location, coordinate, quadrant) trio.
    :return: bool - True if coordinates match, False otherwise.
    """
    return convert_coordinate(azara_record[1]) == rui_record[1]


def create_record(azara_record: tuple[str, str], rui_record: tuple[str, str, str]) -> Union[tuple[str, str, str, str], str]:
    """

    :param azara_record: tuple - a (treasure, coordinate) pair.
    :param rui_record: tuple - a (location, coordinate, quadrant) trio.
    :return:  tuple - combined record, or "not a match" if the records are incompatible.
    """
    return azara_record + rui_record if compare_records(azara_record, rui_record) else "not a match"
    


def clean_up(combined_record_group: tuple[tuple[str, str], tuple[str, str, str]]) -> str:
    """

    :param combined_record_group: tuple of tuples - everything from both participants.
    :return: string of tuples separated by newlines - everything "cleaned". Excess coordinates and information removed.
    """
    return "\n".join([str(record[:1] + record[2:]) for record in combined_record_group]) + "\n"
    
