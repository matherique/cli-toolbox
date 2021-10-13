package toolkit

type Command interface {
	Name() string
	Desc() string
	SetDesc(text string)
	Sub() []*command
	AddSub(sub Command)
	LogDesc() string
	SetLongDesc(text string)
	Handler() handler
	SetHandler(fn handler)
}

type handler = func(args []string) error

type command struct {
	subcmd   []*command
	name     string
	desc     string
	longDesc string
	handler  handler
}

func New(name string) *command {
	cmd := new(command)
	cmd.name = name

	return cmd
}

// Name return the name of the command
func (c *command) Name() string { return c.name }

// Desc return the description of the command
func (c *command) Desc() string { return c.desc }

// SetDesc add description to the command
func (c *command) SetDesc(text string) { c.desc = text }

// AddSub add a new subcommand to the command
func (c *command) AddSub(sub *command) {
	c.subcmd = append(c.subcmd, sub)
}

// Subs return all subcommands added in command
func (c *command) Sub() []*command { return c.subcmd }

// LongDesc return the long description
func (c *command) LongDesc() string { return c.longDesc }

// AddLongDesc return the long description
func (c *command) SetLongDesc(text string) { c.longDesc = text }

// Handler return the handler function that executes when the command
// subcommand is called
func (c *command) Handler() handler { return c.handler }

func (c *command) SetHandler(fn handler) { c.handler = fn }
