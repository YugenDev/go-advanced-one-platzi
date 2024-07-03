## Patrón Strategy en Go

### Propósito

El patrón Strategy, también conocido como Estrategia, define una familia de algoritmos, encapsula cada uno de ellos y los hace intercambiables. Permite que un algoritmo varíe independientemente de los clientes que lo usan.

### Escenarios de uso

* **Algoritmos con diferentes comportamientos:** Se utiliza para implementar diferentes algoritmos para una misma tarea, como ordenar datos o realizar cálculos.
* **Cambio de comportamiento en tiempo de ejecución:** Permite cambiar el algoritmo utilizado sin modificar el código del cliente.
* **Reutilización de código:** Facilita la reutilización de algoritmos en diferentes partes del programa.

### Implementación

En Go, podemos implementar el patrón Strategy utilizando interfaces, structs y funciones.

1. **Definir la interfaz `Strategy`:** Creamos una interfaz que define el método común a todos los algoritmos.

```go
type Strategy interface {
    DoOperation(a int, b int) int
}
```

2. **Crear las estructuras `ConcreteStrategy`:** Implementamos la interfaz `Strategy` para cada algoritmo concreto.

```go
type SumaStrategy struct{}

func (s SumaStrategy) DoOperation(a int, b int) int {
    return a + b
}

type RestaStrategy struct{}

func (s RestaStrategy) DoOperation(a int, b int) int {
    return a - b
}
```

3. **Crear la estructura `Context`:** Implementa la lógica que utiliza el algoritmo y mantiene una referencia a la estrategia actual.

```go
type Context struct {
    strategy Strategy
}

func (c *Context) SetStrategy(strategy Strategy) {
    c.strategy = strategy
}

func (c *Context) ExecuteOperation(a int, b int) int {
    return c.strategy.DoOperation(a, b)
}
```

4. **Utilizar el patrón Strategy:** Creamos instancias de las estrategias concretas y las configuramos en el contexto.

```go
sumaStrategy := SumaStrategy{}
restaStrategy := RestaStrategy{}

context := Context{}

context.SetStrategy(sumaStrategy)
resultadoSuma := context.ExecuteOperation(5, 3)
fmt.Println("Suma:", resultadoSuma)

context.SetStrategy(restaStrategy)
resultadoResta := context.ExecuteOperation(5, 3)
fmt.Println("Resta:", resultadoResta)
```

### Ventajas y desventajas

**Ventajas:**

* **Flexibilidad:** Permite cambiar el algoritmo utilizado sin modificar el código del cliente.
* **Reutilización de código:** Facilita la reutilización de algoritmos en diferentes partes del programa.
* **Encapsulamiento:** Separa la lógica de los algoritmos de la lógica del cliente.

**Desventajas:**

* **Complejidad:** Aumenta la complejidad del código al introducir más clases e interfaces.
* **Overhead:** Se genera un pequeño overhead al tener que delegar la ejecución al algoritmo seleccionado.
* **Rigidez:** Puede dificultar la modificación de la interfaz de la estrategia en el futuro.

### Alternativas

* **Patrón Factory Method:** Permite crear objetos de diferentes tipos sin necesidad de instanciarlos directamente.
* **Dependencia de inyección:** Permite inyectar dependencias en clases, lo que facilita las pruebas y el mantenimiento del código.

### Conclusión

El patrón Strategy es una herramienta útil para implementar algoritmos intercambiables y mejorar la flexibilidad del código. Sin embargo, es importante considerar las ventajas y desventajas antes de utilizarlo y evaluar las alternativas disponibles.

**¿Tienes alguna pregunta sobre el patrón Strategy o sobre algún otro tema de desarrollo de software?**

### Recursos adicionales

* **Refactoring Guru:** [https://refactoring.guru/design-patterns/strategy/go/example](https://refactoring.guru/design-patterns/strategy/go/example)
* **Patrones de diseño con Go:** [https://www.youtube.com/watch?v=xHhalCokHmA](https://www.youtube.com/watch?v=xHhalCokHmA)