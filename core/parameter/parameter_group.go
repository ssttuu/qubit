package parameter

type Group struct {
	ParameterType
	GroupType string
	Children  ParameterSpecs
}

func (g *Group) GetParameter(name string) *Parameter {
	for _, ps := range g.Children {
		if ps.Name == name {
			return ps.Parameter
		}
	}
	return nil
}

func (g *Group) GetBool(name string) Bool {
	return g.GetParameter(name).GetBool()
}

func (g *Group) GetInt64(name string) Int64 {
	return g.GetParameter(name).GetInt64()
}

func (g *Group) GetFloat64(name string) Float64 {
	return g.GetParameter(name).GetFloat64()
}

func (g *Group) GetString(name string) String {
	return g.GetParameter(name).GetString()
}

func (g *Group) GetEnum(name string) Enum {
	return g.GetParameter(name).GetEnum()
}

func (g *Group) GetGroup(name string) Group {
	return g.GetParameter(name).GetGroup()
}

func (g *Group) GetMulti(name string) Multi {
	return g.GetParameter(name).GetMulti()
}

func NewGroupParameter(groupType string, children ParameterSpecs) *Parameter {
	return &Parameter{
		TypeName: "group",
		Type: Group{
			GroupType: groupType,
			Children:  children,
		},
	}
}
