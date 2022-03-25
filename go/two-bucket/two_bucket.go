package twobucket

import "errors"

type Bucket struct {
    tag string
    capacity int
    level int
}

func (bucket *Bucket) isEmpty() bool {
    return bucket.level == 0
}

func (bucket *Bucket) isFull() bool {
    return bucket.level == bucket.capacity
}

func (bucket *Bucket) pourInto(otherBucket *Bucket) {
    var amount int
    available := otherBucket.capacity - otherBucket.level
    if bucket.level < available {
        amount = bucket.level
    } else {
        amount = available
    }
    bucket.level -= amount
    otherBucket.level += amount
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func Solve(sizeBucketOne, sizeBucketTwo, goalAmount int, startBucket string) (goalBucket string, numSteps, otherBucketLevel int, err error) {
	if sizeBucketOne < 1 || sizeBucketTwo < 1 || goalAmount < 1 || (startBucket != "one" && startBucket != "two") {
        return "", 0, 0, errors.New("invalid arguments")
    } else if goalAmount % gcd(sizeBucketOne, sizeBucketTwo) != 0 {
        return "", 0, 0, errors.New("no solutions")
    }
    var first, second *Bucket
    switch startBucket {
        default: first, second = &Bucket{"two", sizeBucketTwo, 0}, &Bucket{"one", sizeBucketOne, 0}
        case "one": first, second = &Bucket{"one", sizeBucketOne, 0}, &Bucket{"two", sizeBucketTwo, 0}
    }
    for first.level != goalAmount && second.level != goalAmount {
        switch {
            default: first.pourInto(second)
            case first.isEmpty(): first.level = first.capacity
            case second.isFull(): second.level = 0
            case first.capacity == goalAmount: first.level = first.capacity
            case second.capacity == goalAmount: second.level = second.capacity
        }
        numSteps++
    }
    if first.level == goalAmount {
        goalBucket, otherBucketLevel = first.tag, second.level
    } else {
        goalBucket, otherBucketLevel = second.tag, first.level
    }
    return
}

