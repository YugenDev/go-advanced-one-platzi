## Patrón Factory en Desarrollo de Software

Imagina que estás en una panadería y quieres un pastel. No necesitas saber cómo se hace cada tipo de pastel, solo le dices al panadero qué tipo quieres (chocolate, fresa, etc.) y él te lo prepara.

El patrón factory funciona de manera similar en software. Define una interfaz para crear objetos, pero deja que las subclases decidan qué tipo de objeto crear. Esto te permite crear objetos sin tener que saber cómo se implementan internamente.

### Beneficios del Patrón Factory

* **Mejora la modularidad**: El código de creación de objetos se separa del código que los utiliza, lo que facilita el mantenimiento y la reutilización.
* **Incrementa la flexibilidad**: Puedes agregar nuevos tipos de objetos sin modificar el código cliente.
* **Promueve la abstracción**: Oculta los detalles de implementación de los objetos.

### Ejemplo en Go

Veamos un ejemplo simple del patrón factory en Go:

```go
// Interfaz para productos
type Producto interface {
  Operar() string
}

// Producto concreto 1
type ProductoA struct{}

func (p *ProductoA) Operar() string {
  return "Soy un Producto A"
}

// Producto concreto 2
type ProductoB struct{}

func (p *ProductoB) Operar() string {
  return "Soy un Producto B"
}

// Interfaz para fábricas
type Factory interface {
  CrearProducto() Producto
}

// Fábrica concreta 1
type FabricaA struct{}

func (f *FabricaA) CrearProducto() Producto {
  return &ProductoA{}
}

// Fábrica concreta 2
type FabricaB struct{}

func (f *FabricaB) CrearProducto() Producto {
  return &ProductoB{}
}

// Código cliente
func main() {
  // Crea una fábrica según el tipo deseado
  fabrica := FabricaA{}

  // Crea un producto usando la fábrica
  producto := fabrica.CrearProducto()

  // Utiliza el producto
  resultado := producto.Operar()
  fmt.Println(resultado) // Imprime: "Soy un Producto A"
}
```

En este ejemplo, las interfaces `Producto` y `Factory` definen las estructuras generales para productos y fábricas, respectivamente. Las implementaciones concretas (`ProductoA`, `ProductoB`, `FabricaA`, `FabricaB`) definen los tipos específicos de productos y fábricas. El código cliente crea una fábrica según el tipo de producto deseado y luego usa la fábrica para crear el producto.

Este es un ejemplo básico, pero el patrón factory se puede usar para crear implementaciones más complejas con múltiples tipos de productos y fábricas.

### Recursos adicionales

* [Patrón Factory en Go - Refactoring.Guru]: [https://refactoring.guru/design-patterns/factory-method/go/example](https://refactoring.guru/design-patterns/factory-method/go/example)
* [Patrones de diseño con Go - YouTube]: [https://m.youtube.com/watch?v=qHuI_IWUkfA](https://m.youtube.com/watch?v=qHuI_IWUkfA)
* [Factory Method | Design Patterns in Go - GitBook]: [https://refactoring.guru/design-patterns/factory-method/go/example](https://refactoring.guru/design-patterns/factory-method/go/example)