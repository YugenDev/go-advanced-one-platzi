## Patrón Adapter en Go

### Escenario de uso

Imaginemos que tenemos una librería de terceros que define una interfaz para dibujar formas geométricas, pero nuestra aplicación utiliza una interfaz diferente. El patrón Adapter nos permite adaptar la interfaz de la librería a la que utiliza nuestra aplicación, sin modificar el código fuente de la librería.

### Implementación

En Go, podemos implementar el patrón Adapter utilizando interfaces, structs y composición.

1. **Definir las interfaces:** Creamos dos interfaces: `Forma` para la librería de terceros y `FormaDibujable` para nuestra aplicación.

```go
type Forma interface {
    GetArea() float64
}

type FormaDibujable interface {
    Dibujar()
}
```

2. **Crear la estructura Adapter:** Creamos una estructura `Adapter` que implementa la interfaz `FormaDibujable` y contiene una instancia de la estructura `Forma` de la librería.

```go
type Adapter struct {
    forma Forma
}

func (a *Adapter) Dibujar() {
    // Adaptar la llamada al método de la librería
    area := a.forma.GetArea()
    fmt.Println("Dibujando forma con área:", area)
}
```

3. **Utilizar el Adapter:** Creamos una instancia de la estructura `Forma` de la librería y la envolvemos en un `Adapter` para poder usarla con nuestra aplicación.

```go
// Crear una forma de la librería de terceros
formaLib := Rectangulo{5, 3}

// Envolver la forma en un Adapter
adapter := Adapter{formaLib}

// Usar el Adapter con la interfaz de nuestra aplicación
adapter.Dibujar()
```

### Ventajas y desventajas

**Ventajas:**

* **Flexibilidad:** Permite integrar librerías con interfaces incompatibles.
* **Reutilización:** Facilita la reutilización de código de terceros.
* **Modularidad:** Mejora la modularidad al separar la lógica de adaptación.

**Desventajas:**

* **Complejidad:** Aumenta la complejidad del código al introducir una clase adicional.
* **Encapsulamiento:** Puede romper la encapsulación de la librería de terceros.
* **Rigidez:** Puede dificultar la modificación de las interfaces en el futuro.

### Alternativas

* **Patrón Bridge:** Similar al Adapter, pero separa la lógica de adaptación de la interfaz.
* **Patrón Facade:** Proporciona una interfaz simplificada para una librería compleja.

### Conclusión

El patrón Adapter es una herramienta útil para integrar librerías con interfaces incompatibles y mejorar la flexibilidad del código. Sin embargo, es importante evaluar las alternativas y considerar las desventajas antes de utilizarlo.

**¿Tienes alguna pregunta sobre el patrón Adapter o sobre algún otro tema de desarrollo de software?**

### Recursos adicionales

* **Refactoring Guru:** [https://refactoring.guru/design-patterns/adapter/go/example](https://refactoring.guru/design-patterns/adapter/go/example)
* **Patrones de diseño con Go:** [https://www.youtube.com/watch?v=xHhalCokHmA](https://www.youtube.com/watch?v=xHhalCokHmA)