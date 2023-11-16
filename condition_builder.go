package cypher

type ConditionBuilder struct {
	conditions []ConditionalConfig
}

func NewConditionBuilder() *ConditionBuilder {
	return &ConditionBuilder{}
}

func (cb *ConditionBuilder) AddCondition(condition ConditionalConfig) *ConditionBuilder {
	cb.conditions = append(cb.conditions, condition)
	return cb
}

func (cb *ConditionBuilder) And() *ConditionBuilder {
	if len(cb.conditions) == 0 {
		return cb
	}
	cb.conditions[len(cb.conditions)-1].Condition = AND
	return cb
}
func (cb *ConditionBuilder) Or() *ConditionBuilder {
	if len(cb.conditions) == 0 {
		return cb
	}
	cb.conditions[len(cb.conditions)-1].Condition = OR
	return cb
}

func (cb *ConditionBuilder) ReleaseConditions() []ConditionalConfig {
	return cb.conditions
}
func (cb *ConditionBuilder) Clear() *ConditionBuilder {
	cb.conditions = make([]ConditionalConfig, 0)
	return cb
}
