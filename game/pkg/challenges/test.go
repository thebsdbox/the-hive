package challenges

import "time"

var test = Challenge{
	Name:        "This challenge is a test",
	Description: "As in a LITERAL test",
	AllowedTime: 1 * time.Hour,
}

func init() {

	Challenges = append(Challenges, test)

}
