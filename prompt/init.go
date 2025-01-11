package prompt

var (
	ChatAssistant    string
	TellJoke         string
	TellJokeUser     string
	JudgeSubmit      string
	ProblemParse     string
	ProblemTranslate string
	ProblemGenerate  string
	SolutionGenerate string
	TestcaseGenerate string
)

func InitPrompt() {
	ChatAssistant = chatAssistantZh

	TellJoke = tellJokeZh
	TellJokeUser = tellJokeUserZh

	JudgeSubmit = judgeSubmitZh

	ProblemParse = problemParseZh
	ProblemGenerate = problemGenerateZh
	ProblemTranslate = problemTranslateZh

	SolutionGenerate = solutionGenerateZh

	TestcaseGenerate = testcaseGenerateZh
}
