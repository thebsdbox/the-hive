package challenges

import "time"

var challenge2 = Challenge{
	Name:        "The mild one",
	Description: "This one involves some spice",
	AllowedTime: 10 * time.Minute,
}

func init() {

	Challenges = append(Challenges, challenge2)

}
