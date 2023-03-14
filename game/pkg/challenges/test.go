package challenges

import "time"

var test = Challenge{
	Name:        "This challenge is a test",
	Description: "As in a LITERAL test",
	AllowedTime: 20 * time.Second,
}

func init() {

	Challenges = append(Challenges, test)

}
