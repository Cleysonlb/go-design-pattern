package singleton

/*
Singleton

O padrão de design criacional de Singleton restringe a instanciação de um tipo para um único objeto.
Garanta que uma classe tenha apenas uma instância e forneça um ponto de acesso global a ela.
Encapsulado "inicialização just-in-time" ou "inicialização no primeiro uso".
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
