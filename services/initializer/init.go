package initializer

var group []Interface

// Interface is the type to call early on system initialize call
type Interface interface {
	Initialize() func()
}

// Register a module in initializer
func Register(initializer Interface) {
	group = append(group, initializer)
}

// Initialize all modules and return the finalizer function
func Initialize() func() {
	var finalizers []func()
	for i := range group {
		if finalizer := group[i].Initialize(); finalizer != nil {
			finalizers = append(finalizers, finalizer)
		}
	}

	return func() {
		for i := range finalizers {
			finalizers[i]()
		}
	}
}
