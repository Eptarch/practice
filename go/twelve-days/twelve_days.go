package twelve

import "strings"

type Day struct {
    ordinal, gift string
}

var carol = map[int]*Day{
    1:  {ordinal: "first",    gift: "a Partridge in a Pear Tree"},
    2:  {ordinal: "second",   gift: "two Turtle Doves"},
    3:  {ordinal: "third",    gift: "three French Hens"},
    4:  {ordinal: "fourth",   gift: "four Calling Birds"},
    5:  {ordinal: "fifth",    gift: "five Gold Rings"},
    6:  {ordinal: "sixth",    gift: "six Geese-a-Laying"},
    7:  {ordinal: "seventh",  gift: "seven Swans-a-Swimming"},
    8:  {ordinal: "eighth",   gift: "eight Maids-a-Milking"},
    9:  {ordinal: "ninth",    gift: "nine Ladies Dancing"},
    10: {ordinal: "tenth",    gift: "ten Lords-a-Leaping"},
    11: {ordinal: "eleventh", gift: "eleven Pipers Piping"},
    12: {ordinal: "twelfth",  gift: "twelve Drummers Drumming"},
}

func Verse(i int) string {
    var sb strings.Builder
    sb.WriteString("On the ")
    sb.WriteString(carol[i].ordinal)
    sb.WriteString(" day of Christmas my true love gave to me: ")
    for c := i; c > 0; c-- {
        switch {
            case i == 1:
                sb.WriteString(carol[c].gift)
                sb.WriteString(".")
            case c == 1:
                sb.WriteString("and ")
                sb.WriteString(carol[c].gift)
                sb.WriteString(".")
            default:
                sb.WriteString(carol[c].gift)
                sb.WriteString(", ")
        }
    }
    return sb.String()
}

func Song() string {
    var sb strings.Builder
    for i := 1; i <= 12; i++ {
        sb.WriteString(Verse(i))
        if i != 12 {
            sb.WriteString("\n")
        }
    }
    return sb.String()
}
