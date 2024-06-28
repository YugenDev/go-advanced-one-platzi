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


### Diferencia entre Composición y Herencia

La herencia es un mecanismo poderoso en POO que permite crear nuevas clases (clases hijas) que heredan propiedades y comportamientos de clases existentes (clases padre). Las clases hijas pueden especializar la funcionalidad heredada o agregar nuevas características.

La composición, por otro lado, es una técnica para crear objetos que contienen instancias de otros objetos. Estos objetos contenidos se conocen como objetos miembros o partes. El objeto contenedor puede acceder a los métodos y propiedades públicos de los objetos miembros para lograr su funcionalidad.

### Ejemplo en el Código

El código implementa una fábrica de computadoras que puede producir laptops y desktops. A continuación, se detalla cómo demuestra la composición:

1. **Interfaz `IProduct`:** Esta interfaz define un conjunto de métodos que todos los tipos de productos (laptops y desktops) deben seguir. Estos métodos incluyen `setStock`, `getStock`, `setName` y `getName`. Esta interfaz promueve la reutilización del código y garantiza que la fábrica pueda tratar diferentes tipos de productos de manera similar.

2. **Clase `Computer`:** Esta clase sirve como base para `Laptop` y `Desktop`. Encapsula atributos comunes como `name` y `stock` y proporciona métodos para accederlos y modificarlos.

3. **Clases `Laptop` y `Desktop`:** Estas clases heredan de la clase `Computer`. No heredan ninguna funcionalidad o propiedad adicional. En cambio, aprovechan el enfoque de composición.

4. **Funciones `newLaptop` y `newDesktop`:** Estas funciones son responsables de crear objetos `Laptop` y `Desktop`, respectivamente. Sin embargo, no crean directamente instancias de estas clases. En cambio, crean un objeto `Computer` con valores predefinidos (nombre y stock) y lo devuelven.

### Beneficios de la Composición en este Ejemplo

- **Flexibilidad:** El código puede acomodar fácilmente nuevos tipos de productos sin modificar la jerarquía de clases existente. Al simplemente implementar la interfaz `IProduct`, se pueden integrar nuevos tipos de productos en la fábrica.
- **Mantenibilidad:** La separación de preocupaciones entre la clase base `Computer` y los tipos de productos específicos (`Laptop` y `Desktop`) facilita la comprensión y el mantenimiento del código.
- **Acoplamiento Débil:** La función `GetComputerFactory` no depende de clases concretas de laptop o desktop. Solo interactúa con la interfaz `IProduct`, lo que promueve un acoplamiento débil y hace que el código sea más adaptable a los cambios.

En esencia, el código demuestra cómo la composición puede ser un enfoque más flexible y mantenible en comparación con la herencia en este escenario. Logra la reutilización del código a través de interfaces y promueve un acoplamiento débil entre las diferentes partes del código.

### Recursos adicionales

Si deseas profundizar en este tema, te recomiendo consultar los siguientes recursos:

- **Composición sobre herencia en POO:** [https://m.youtube.com/watch?v=OyTPDFyGWRc](https://m.youtube.com/watch?v=OyTPDFyGWRc)
- **¿Por qué utilizar composición sobre herencia en POO?:** [https://www.youtube.com/watch?v=9NynVRpZzv4](https://www.youtube.com/watch?v=9NynVRpZzv4)
- **Herencia vs Composición ¿Tienes claro cuál es el rival más débil?:** [https://www.nature.com/articles/s41598-018-22531-2](https://www.nature.com/articles/s41598-018-22531-2)

Espero que esta explicación en español te haya sido útil. ¡No dudes en preguntar si tienes más dudas!