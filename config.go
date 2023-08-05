package goalint

var Configurator ConfiguratorFunc

func Configure(f ConfiguratorFunc) struct{} {
	Configurator = f

	return struct{}{}
}
