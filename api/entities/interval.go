package entities

import "time"

type (
	Interval struct {
		Begin time.Time `json:"begin"`
		End   time.Time `json:"end"`
	}
)

func (first Interval) Overlapping(second Interval) bool {
	return second.Begin.Equal(first.Begin) || second.End.Equal(first.End) || (second.Begin.After(first.Begin) && second.Begin.Before(first.End)) || (first.Begin.After(second.Begin) && first.Begin.Before(second.End)) || second.End.After(first.Begin) && (second.End.Before(first.End) || second.End.Equal(first.End)) || first.End.After(second.Begin) && (first.End.Before(second.End) || first.End.Equal(second.End))
}

func (first Interval) MergeWith(second Interval) Interval {
	return Interval{
		Begin: func(t1, t2 time.Time) time.Time {
			if t1.After(t2) {
				return t2
			} else {
				return t1
			}
		}(first.Begin, second.Begin),
		End: func(t1, t2 time.Time) time.Time {
			if t1.After(t2) {
				return t1
			} else {
				return t2
			}
		}(first.End, second.End),
	}
}
