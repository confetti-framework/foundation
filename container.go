package foundation

import (
	"github.com/lanvard/contract/inter"
	"github.com/lanvard/errors"
	"github.com/lanvard/support"
	"reflect"
)

type Container struct {
	// The container created at boot time
	bootContainer inter.Container

	// The container's bindings.
	bindings inter.Bindings
}

func NewContainer() *Container {
	containerStruct := Container{}
	containerStruct.bindings = make(inter.Bindings)

	return &containerStruct
}

func NewContainerByBoot(bootContainer inter.Container) inter.Container {
	container := NewContainer()
	container.bootContainer = bootContainer

	return container
}

// Determine if the given abstract type has been bound.
func (c *Container) Bound(abstract string) bool {
	_, bound := c.bindings[abstract]
	return bound
}

// Register a binding with the container.
func (c *Container) Bind(abstract interface{}, concrete interface{}) {
	abstractString := support.Name(abstract)

	c.bindings[abstractString] = concrete
}

// Register a shared binding in the container.
func (c *Container) Singleton(abstract interface{}, concrete interface{}) {
	c.Bind(abstract, concrete)
}

// Register an existing instance as shared in the container without an abstract
func (c *Container) Instance(concrete interface{}) interface{} {
	c.Bind(concrete, concrete)

	return concrete
}

// GetE the container's bindings.
func (c *Container) Bindings() inter.Bindings {
	return c.bindings
}

// MakeE the given type from the container.
func (c *Container) Make(abstract interface{}) interface{} {
	concrete, err := c.MakeE(abstract)
	if nil != err {
		panic(err)
	}
	return concrete
}

// MakeE the given type from the container.
func (c *Container) MakeE(abstract interface{}) (interface{}, error) {
	var concrete interface{}
	var err error = nil
	var abstractName = support.Name(abstract)

	kind := support.Kind(abstract)
	if kind == reflect.Ptr && abstract == nil {
		return nil, errors.New("can't resolve interface. To resolve an interface, " +
			"use the following syntax: (*interface)(nil), use a string or use the struct itself")
	}

	if object, present := c.bindings[abstractName]; present {
		concrete = c.getConcreteBinding(concrete, object, abstractName)

	} else if c.bootContainer != nil && c.bootContainer.Bound(abstractName) {
		// Check the container that was created at boot time
		concrete, err = c.bootContainer.MakeE(abstract)

	} else if kind == reflect.Struct {
		// If struct cannot be found, we simply have to use the struct itself.
		concrete = abstract
	} else if kind == reflect.String {
		var instances support.Map
		instances, err = support.NewMapE(c.bindings)
		if err == nil {
			var value support.Value
			if c.bootContainer != nil {
				bootBindings := c.bootContainer.Bindings()
				bootInstances, err := support.NewMapE(bootBindings)
				if err != nil {
					return nil, err
				}
				instances.Merge(bootInstances)
			}
			value, err = instances.GetE(abstract.(string))
			//goland:noinspection GoNilness
			concrete = value.Raw()
		}
	}

	if err != nil {
		err = errors.Wrap(err, "get instance from container")
	}

	return concrete, err
}

func (c *Container) getConcreteBinding(concrete interface{}, object interface{}, abstractName string) interface{} {
	// If abstract is bound, use that object.
	concrete = object

	// If concrete is a callback, run it and save the result
	if support.Kind(concrete) == reflect.Func {
		callback, simpleCallback := concrete.(func() interface{})
		if simpleCallback {
			concrete = callback()
		}
	}
	c.bindings[abstractName] = concrete

	return concrete
}

// "Extend" an abstract type in the container.
func (c *Container) Extend(abstract interface{}, function func(service interface{}) interface{}) {
	concrete := c.Make(abstract)

	newConcrete := function(concrete)

	c.Bind(abstract, newConcrete)
}
