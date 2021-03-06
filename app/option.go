package app

//Flag is a command line flag
type Flag = Option

//Flags is a set of Flag
type Flags = Options

//Arg is a command line arg
type Arg = Option

//Args is a set of Arg
type Args = Options

//Option reresents any comand line's flag or arg
type Option struct {
	//Name of the opton. If option is a flag it will be triggered using
	//"--Name" synthax.
	Name string

	//Usage contains a description of the Option
	Usage string

	//Var is the variable that will contain the Option actual value after
	//command line has been parsed.
	//Supported values are: int64, string and []strings
	//If args is of type []string, it should be the last of the accepted
	//arguments' list otherwise the parsing will panic.
	Var interface{}

	//Optional indicate sthat the arguments can be omitted. It is actually
	//working if args is the last of the accepted arguments' list otherwise the
	//parsing will panic.
	Optional bool
}

func (o *Option) value() value {
	return newValue(o.Var)
}

func (o *Option) isBool() bool {
	_, ok := o.Var.(*bool)
	return ok
}

func (o *Option) isCumulative() bool {
	_, ok := o.Var.(*[]string)
	return ok
}

//Options represents a set of flags or args
type Options []*Option

//Add adds a new option to an options set
func (o *Options) Add(opt *Option) {
	o.append(opt)
}

func (o *Options) get(name string) *Option {
	for _, opt := range *o {
		if opt.Name == name {
			return opt
		}
	}
	return nil
}

func (o *Options) append(opt *Option) {
	*o = append(*o, opt)
}
