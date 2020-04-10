package singleton

/*
Singleton

O Singleton é um padrão de projeto criacional que permite a você
garantir que uma classe tenha apenas uma instância, enquanto provê
um ponto de acesso global para essa instância.
Referencia: https://refactoring.guru/pt-br/design-patterns/singleton
*/
type singleton struct {
	Name string
}

var instance *singleton

func New() *singleton {
	if instance == nil {
		instance = &singleton{}
	}
	return instance
}
