package add

// This needs to get deleted

type Step int

const (
	StepProject Step = iota
	StepProjectAdd
	StepStopwatch
	StepDescription
)

// func (s Step) String() string {
// 	switch s {
// 	case StepProject:
// 		return "project"
// 	case StepProjectAdd:
// 		return "project_add"
// 	case StepStopwatch:
// 		return "stopwatch"
// 	case StepDescription:
// 		return "description"
// 	default:
// 		return "unknown"
// 	}
// }
