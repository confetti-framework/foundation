package foundation

import (
	"github.com/lanvard/contract/inter"
	"github.com/lanvard/support"
	"reflect"
)

type Container struct {
	// The container created at boot time
	bootContainer inter.Container

	// The container's bindings.
	bindings inter.Bindings

	// The registered type aliases.
	aliases map[string]interface{}

	// The registered aliases keyed by the abstract name.
	abstractAliases map[string]map[string]interface{}

	// The container's shared instances.
	instances inter.Instances
}

func NewContainer() *Container {
	containerStruct := Container{}
	containerStruct.bindings = make(inter.Bindings)
	containerStruct.instances = make(inter.Instances)

	return &containerStruct
}

func NewContainerByBoot(bootContainer inter.Container) inter.Container {
	container := NewContainer()
	container.bootContainer = bootContainer

	return container
}

// Determine if the given abstract type has been bound.
func (c *Container) Bound(abstract string) bool {
	_, bind := c.bindings[abstract]
	_, instance := c.instances[abstract]

	if bind || instance || c.IsAlias(abstract) {
		return true
	}

	return false
}

// Determine if a given string is an alias.
func (c *Container) IsAlias(name string) bool {
	if _, ok := c.aliases[name]; ok {
		return true
	}

	return false
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

// Register an existing instance as shared in the container.
func (c *Container) Instance(abstract interface{}, concrete interface{}) interface{} {
	abstractName := support.Name(abstract)

	_, ok := c.aliases[abstractName]
	if ok {
		delete(c.aliases, abstractName)
	}

	if c.instances == nil {
		c.instances = make(inter.Instances)
	}

	c.instances[abstractName] = concrete

	return concrete
}

// Register an existing instance as shared in the container without an abstract
func (c *Container) BindStruct(concrete interface{}) interface{} {
	c.Instance(concrete, concrete)

	return concrete
}

// Get the container's bindings.
func (c *Container) Bindings() inter.Bindings {
	return c.bindings
}

// Resolve the given type from the container.
func (c *Container) Make(abstract interface{}) interface{} {
	return c.resolve(abstract)
}

// Resolve the given type from the container.
func (c *Container) resolve(abstract interface{}) interface{} {
	var concrete interface{}
	var abstractName = support.Name(abstract)

	if support.Type(abstract) == reflect.Ptr && abstract == nil {
		panic("Can't resolve interface. To resolve an interface, use the following syntax: (*interface)(nil), " +
			"use a string or use the struct itself.")
	}

	if object, present := c.instances[abstractName]; present {
		// The instance will always be returned on subsequent calls
		concrete = object

	} else if support.Type(abstract) == reflect.String && c.IsAlias(abstractName) {
		concrete = c.getConcreteAlias(concrete, abstract)

	} else if object, present := c.bindings[abstractName]; present {
		concrete = c.getConcreteBinding(concrete, object, abstractName)

	} else if c.bootContainer != nil && c.bootContainer.Bound(abstractName) {
		// Check the container that was created at boot time
		concrete = c.bootContainer.Make(abstract)

	} else if support.Type(abstract) == reflect.Struct {
		// If struct cannot be found, we simply have to use the struct itself.
		concrete = abstract
	}

	return concrete
}

func (c *Container) getConcreteBinding(concrete interface{}, object interface{}, abstractName string) interface{} {
	// If abstract is bound, use that object.
	concrete = object

	// If concrete is a callback, run it and save the result
	if support.Type(concrete) == reflect.Func {
		callback, simpleCallback := concrete.(func() interface{})
		if simpleCallback {
			concrete = callback()
		}
	}
	c.bindings[abstractName] = concrete

	return concrete
}

func (c *Container) getConcreteAlias(concrete, abstract interface{}) interface{} {
	concrete = c.aliases[abstract.(string)]

	// If concrete is a callback, run it and save the result
	if support.Type(concrete) == reflect.Func {
		callback, simpleCallback := concrete.(func() interface{})
		if simpleCallback {
			concrete = callback()
		}
	}
	c.aliases[abstract.(string)] = concrete

	return concrete
}

// "Extend" an abstract type in the container.
func (c *Container) Extend(abstract interface{}, function func(service interface{}) interface{}) {
	concrete := c.resolve(abstract)

	newConcrete := function(concrete)

	c.Instance(abstract, newConcrete)
}
