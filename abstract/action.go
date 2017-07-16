package abstract

/*
*	Parameter structure, used for defining parameter data
 */
type Parameter struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
	Mandatory   bool   `yaml:"mandatory"`
	HasValue    bool   `yaml:"hasvalue"`
	SampleValue string `yaml:"sample"`
}

/*
*	Action structure, used for defining action data
 */
type Action struct {
	Parameters  []Parameter `yaml:"parameters"`
	Name        string      `yaml:"name"`
	Command     string      `yaml:"command"`
	Description string      `yaml:"description"`
	Usage       string      `yaml:"usage"`
}

/*
*	ActionOps interface, defining abstract Action Operations
 */
type ActionOps interface {
	/*
	*	Init method
	* Initialise Action to work
	 */
	Init() bool
	/*
	*	Reset method
	* Reset Action for next work
	 */
	Reset() bool
	/*
	*	Execute method
	* Execute Action accordingly to parameters
	* Return execution start/complete status
	 */
	Execute(chan string) bool
	/*
	*	IsInProgress method
	* Return execution progress state
	 */
	IsInProgress() bool
	/*
	*	GetCommand method
	* Return command name
	 */
	GetCommand() string
	/*
	*	GetName method
	* Return command descriptive name
	 */
	GetName() string
	/*
	*	GetUsage method
	* Return command usage sample
	 */
	GetUsage() string
	/*
	*	AcquireValues method
	* Return flag if parameters are satisfied
	 */
	AcquireValues() bool
	/*
	*	GetExitCode method
	* Return execution exit code (0 - ok)
	 */
	GetExitCode() int
	/*
	*	GetLastMessage method
	* Return operation message or error message
	 */
	GetLastMessage() string
}

/*
*	ActionRegistry structure, used for collecting plugins
 */
type ActionRegistry struct {
	Actions []ActionOps
}

/*
*	ActionRegistry Add method, add a single action to registry
 */
func (r *ActionRegistry) Add(a ActionOps) {
	r.Actions = append(r.Actions, a)
}

/*
*	ActionRegistry AddMany method, add a multiple actions to registry
 */
func (r *ActionRegistry) AddMany(a []ActionOps) {
	r.Actions = append(r.Actions, a...)
}

/*
*	ActionRegistry GetActions method, return all actions in registry
 */
func (r *ActionRegistry) GetActions() []ActionOps {
	return r.Actions
}

/*
*	ActionRegistry Size method, return number of actions in registry
 */
func (r *ActionRegistry) Size() int {
	return len(r.Actions)
}

/*
*	ActionRegistry ElementAt method, return one action in registry at index or nil
 */
func (r *ActionRegistry) ElementAt(index int) ActionOps {
	if index >= r.Size() {
		return nil
	}
	return r.Actions[index]
}

var ActiveActionRegistry ActionRegistry = ActionRegistry{}

func RegisterAction(Action ActionOps) {
	ActiveActionRegistry.Add(Action)
}

func GetActionRegistry() *ActionRegistry {
	return &ActiveActionRegistry
}
