## Patrón Observer en Go

### Escenario de uso

Imaginemos una aplicación que tiene un sensor de temperatura. Cuando la temperatura cambia, se debe notificar a diferentes partes del sistema, como una interfaz gráfica de usuario o un registro de datos. El patrón Observer permite implementar este tipo de comunicación de forma flexible y desacoplada.

### Implementación

En Go, podemos implementar el patrón Observer utilizando interfaces, structs y canales.

1. **Definir las interfaces:** Creamos dos interfaces: `Observable` para el sujeto y `Observer` para los observadores.

```go
type Observable interface {
    RegisterObserver(observer Observer)
    RemoveObserver(observer Observer)
    NotifyObservers()
}

type Observer interface {
    Update(data interface{})
}
```

2. **Crear la estructura `Sensor`:** Implementamos la interfaz `Observable` para representar el sensor de temperatura.

```go
type Sensor struct {
    observers []Observer
    temperatura float64
}

func (s *Sensor) RegisterObserver(observer Observer) {
    s.observers = append(s.observers, observer)
}

func (s *Sensor) RemoveObserver(observer Observer) {
    for i, o := range s.observers {
        if o == observer {
            s.observers = append(s.observers[:i], s.observers[i+1:]...)
            break
        }
    }
}

func (s *Sensor) NotifyObservers() {
    for _, observer := range s.observers {
        observer.Update(s.temperatura)
    }
}

func (s *Sensor) MedirTemperatura() {
    // Simular la medición de la temperatura
    s.temperatura = 25.5 + rand.Float64()*5

    // Notificar a los observadores del cambio de temperatura
    s.NotifyObservers()
}
```

3. **Crear la estructura `Observador`:** Implementamos la interfaz `Observer` para representar las partes interesadas en la temperatura.

```go
type TemperaturaGUI struct {
    sensor *Sensor
}

func (t *TemperaturaGUI) Update(data interface{}) {
    temperatura := data.(float64)
    t.ActualizarInterfaz(temperatura)
}

func (t *TemperaturaGUI) ActualizarInterfaz(temperatura float64) {
    // Actualizar la interfaz gráfica de usuario con la nueva temperatura
    fmt.Println("Temperatura actual:", temperatura)
}
```

4. **Utilizar el patrón Observer:** Creamos un `Sensor` y registramos un `TemperaturaGUI` como observador.

```go
sensor := Sensor{}
gui := TemperaturaGUI{&sensor}

sensor.RegisterObserver(&gui)

// Simular la medición de temperatura y notificación de cambios
sensor.MedirTemperatura()
sensor.MedirTemperatura()
```

### Ventajas y desventajas

**Ventajas:**

* **Flexibilidad:** Permite agregar o eliminar observadores dinámicamente.
* **Desacople:** Los observadores no dependen de la implementación interna del sujeto.
* **Comunicación eficiente:** Se pueden notificar a varios observadores con una sola llamada.

**Desventajas:**

* **Complejidad:** Aumenta la complejidad del código al introducir más clases e interfaces.
* **Dependencia cíclica:** Se debe tener cuidado para evitar dependencias circulares entre sujetos y observadores.
* **Broadcast:** Todas las notificaciones se envían a todos los observadores, incluso si no les interesan.

### Alternativas

* **Patrón Chain of Responsibility:** Permite pasar las solicitudes entre objetos en una cadena.
* **Patrón Mediator:** Facilita la comunicación entre objetos sin dependencias directas.

### Conclusión

El patrón Observer es una herramienta útil para implementar sistemas de notificaciones y actualizaciones en tiempo real. Sin embargo, es importante considerar las ventajas y desventajas antes de utilizarlo y evaluar las alternativas disponibles.

**¿Tienes alguna pregunta sobre el patrón Observer o sobre algún otro tema de desarrollo de software?**

### Recursos adicionales

* **Refactoring Guru:** [https://refactoring.guru/design-patterns/observer/go/example](https://refactoring.guru/design-patterns/observer/go/example)
* **Patrones de diseño con Go:** [https://www.youtube.com/watch?v=xHhalCokHmA](https://www.youtube.com/watch?v=xHhalCokHmA)