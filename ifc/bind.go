package ifc

import "fmt"

// BoundAttribute pairs an explicit entity attribute with an instance argument.
type BoundAttribute struct {
	Attribute ExplicitAttribute
	Value     Parameter
}

// AllExplicitAttributes returns the inherited explicit attributes for entityName in declaration order.
func (s *Schema) AllExplicitAttributes(entityName string) ([]ExplicitAttribute, error) {
	if s == nil {
		return nil, ErrNilSchema
	}

	seen := map[string]bool{}
	return s.collectExplicitAttributes(entityName, seen)
}

func (s *Schema) collectExplicitAttributes(entityName string, seen map[string]bool) ([]ExplicitAttribute, error) {
	entity, ok := s.Entity(entityName)
	if !ok {
		return nil, fmt.Errorf("%w %s", ErrUnknownEntity, entityName)
	}
	if seen[entityName] {
		return nil, fmt.Errorf("%w at %s", ErrCyclicInheritance, entityName)
	}
	seen[entityName] = true
	defer delete(seen, entityName)

	attributes := []ExplicitAttribute{}
	for _, supertype := range entity.Supertypes {
		parentAttributes, err := s.collectExplicitAttributes(supertype, seen)
		if err != nil {
			return nil, err
		}
		attributes = append(attributes, parentAttributes...)
	}
	attributes = append(attributes, entity.Attributes...)
	return attributes, nil
}

// Bind aligns an instance argument list with the entity's inherited explicit attributes.
func (s *Schema) Bind(instance *Instance) ([]BoundAttribute, error) {
	if s == nil {
		return nil, ErrNilSchema
	}
	if instance == nil {
		return nil, ErrNilInstance
	}

	attributes, err := s.AllExplicitAttributes(instance.Entity)
	if err != nil {
		return nil, err
	}
	if len(attributes) != len(instance.Arguments) {
		return nil, fmt.Errorf(
			"%w: entity %s expects %d explicit attributes, got %d arguments",
			ErrArgumentCountMismatch,
			instance.Entity,
			len(attributes),
			len(instance.Arguments),
		)
	}

	bound := make([]BoundAttribute, 0, len(attributes))
	for i := range attributes {
		bound = append(bound, BoundAttribute{
			Attribute: attributes[i],
			Value:     instance.Arguments[i],
		})
	}
	return bound, nil
}
