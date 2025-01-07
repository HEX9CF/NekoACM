package model

type Prompt struct {
	Role           string `json:"role"`           // 角色
	Goals          string `json:"goals"`          // 目标
	Constrains     string `json:"constrains"`     // 约束
	Skills         string `json:"skills"`         // 技能
	InputFormat    string `json:"input_format"`   // 输入格式
	OutputFormat   string `json:"output_format"`  // 输出格式
	Workflow       string `json:"workflow"`       // 工作流程
	Initialization string `json:"initialization"` // 初始化
}

func (p *Prompt) String() string {
	var str string
	if p.Role != "" {
		str += "Role: " + p.Role + "\n"
	}
	if p.Goals != "" {
		str += "Goals: " + p.Goals + "\n"
	}
	if p.Constrains != "" {
		str += "Constrains: " + p.Constrains + "\n"
	}
	if p.Skills != "" {
		str += "Skills: " + p.Skills + "\n"
	}
	if p.InputFormat != "" {
		str += "InputFormat: " + p.InputFormat + "\n"
	}
	if p.OutputFormat != "" {
		str += "OutputFormat: " + p.OutputFormat + "\n"
	}
	if p.Workflow != "" {
		str += "Workflow: " + p.Workflow + "\n"
	}
	if p.Initialization != "" {
		str += "Initialization: " + p.Initialization + "\n"
	}
	return str
}
