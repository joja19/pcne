package welcome

func Welcome(name string) string {
	return "Welcome, " + name + "!"
}

func DescribeExam() string {
	return ("The PCNE (Professional Cloud Network Engineer) exam by Google is " +
	"a tough one. It is 2 hours of 50 to 60 multi-choice and multi-select questions. " +
	"The passing score is 85%, so not much room for mistakes, careful! " +
	"It needs good concentration, practice, food and sleep! " +
	"In addition to the Google Skills Boost program offered to partners, " +
	"you may need to review documentation associated with each exam objective. " +
	"This binary will act as your practice companion, to improve your score " +
	"along the journey. It will show questions randomly, and will keep track of " +
	"your score in each attempt. At the end of each attempt, it will show you " +
	"the correct answers, and let you know if you got it right or not. " +
	"Finally, it will show you your score in each objective, and " +
	"Let you know which objectives to focus on more with documentation links.")
}