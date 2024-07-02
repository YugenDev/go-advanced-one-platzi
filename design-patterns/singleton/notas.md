## Patrón de diseño Singleton

El patrón Singleton es un patrón de diseño creacional que garantiza que una clase tenga una única instancia, a la vez que proporciona un punto de acceso global a dicha instancia. 

### Propósito

* **Controlar la creación de objetos:** Se asegura que solo exista una instancia de una clase específica.
* **Acceso global:** Permite acceder a la instancia única desde cualquier parte del programa.
* **Compartir recursos:** Se pueden compartir datos y recursos entre diferentes partes del programa.

### Implementación

La implementación del patrón Singleton se basa en los siguientes principios:

1. **Variable estática privada:** Se crea una variable estática privada dentro de la clase para almacenar la instancia única.
2. **Método de instancia estático:** Se define un método estático que verifica si la instancia existe y la devuelve si ya está creada. Si no existe, la crea y la devuelve.
3. **Constructor privado:** El constructor de la clase se hace privado para evitar la creación de instancias desde fuera de la clase.

### Ejemplo en Python

```python
class Singleton:
    _instancia = None

    def __new__(cls):
        if cls._instancia is None:
            cls._instancia = super().__new__(cls)
        return cls._instancia

    def some_method(self):
        pass

# Obtener la instancia única
singleton = Singleton()
singleton.some_method()
```

### Ventajas y desventajas

**Ventajas:**

* **Control:** Se controla estrictamente el número de instancias de una clase.
* **Acceso global:** Se facilita el acceso a la instancia desde cualquier parte del programa.
* **Compartir recursos:** Se pueden compartir datos y recursos de manera eficiente.

**Desventajas:**

* **Dificultad en pruebas:** Puede ser difícil probar clases que usan el patrón Singleton.
* **Dependencia global:** La dependencia de la instancia única puede dificultar el mantenimiento del código.
* **Abuso:** El uso excesivo del patrón Singleton puede llevar a un código rígido y difícil de mantener.

### Alternativas

El patrón Singleton es útil en algunos casos, pero es importante considerar las alternativas antes de usarlo. Algunas alternativas incluyen:

* **Patrón Factory:** Permite crear objetos de diferentes tipos sin necesidad de instanciarlos directamente.
* **Dependencia de inyección:** Permite inyectar dependencias en clases, lo que facilita las pruebas y el mantenimiento del código.

### Conclusión

El patrón Singleton es un patrón de diseño útil para controlar la creación de objetos y compartir recursos. Sin embargo, es importante usarlo con moderación y considerar las alternativas antes de implementarlo.

**Recursos adicionales:**

* **Wikipedia:** [https://www.juntadeandalucia.es/servicios/madeja/contenido/recurso/202](https://www.juntadeandalucia.es/servicios/madeja/contenido/recurso/202)
* **Refactoring Guru:** [https://refactoring.guru/design-patterns/singleton](https://refactoring.guru/design-patterns/singleton)
* **IONOS:** [https://www.ionos.es/digitalguide/paginas-web/desarrollo-web/patron-singleton/](https://www.ionos.es/digitalguide/paginas-web/desarrollo-web/patron-singleton/)


## Patrón Singleton en Go

### Ejemplo en Go

En Go, no existen métodos estáticos ni variables estáticas como en otros lenguajes. Sin embargo, se puede implementar el patrón Singleton utilizando la biblioteca estándar `sync` y la programación funcional.

```Go
package main

import (
    "fmt"
    "sync"
)

// Estructura Singleton
type Singleton struct {
    Tiempo int64
}

// Variable global para la instancia Singleton
var instancia *Singleton
var once sync.Once

// Función para obtener la instancia Singleton
func GetInstancia() *Singleton {
    once.Do(func() {
        instancia = &Singleton{time.Now().Unix()}
    })
    return instancia
}

func main() {
    // Obtener la instancia Singleton y mostrar su tiempo
    singleton := GetInstancia()
    fmt.Println("Tiempo de la instancia Singleton:", singleton.Tiempo)

    // Esperar un segundo y obtener la instancia de nuevo
    time.Sleep(1 * time.Second)
    singleton2 := GetInstancia()
    fmt.Println("Tiempo de la instancia Singleton:", singleton2.Tiempo)

    // Comprobar que ambas instancias son la misma
    if singleton == singleton2 {
        fmt.Println("Ambas instancias son la misma")
    } else {
        fmt.Println("¡Error! Las instancias son diferentes")
    }
}

```

### Explicación del código

1. **Estructura Singleton:** Define la estructura `Singleton` con un campo `Tiempo` para almacenar un valor.
2. **Variable global `instancia`:** Se declara una variable global `instancia` de tipo `*Singleton` para almacenar la instancia Singleton.
3. **Variable `once` de tipo `sync.Once`:** Se crea una variable `once` de tipo `sync.Once` para garantizar que la inicialización de la instancia Singleton se realice solo una vez.
4. **Función `GetInstancia`:** Define la función `GetInstancia` que devuelve la instancia Singleton.
    - Utiliza la función `Do` del paquete `sync` para ejecutar la función de inicialización solo una vez.
    - Dentro de la función de inicialización, se crea una nueva instancia de `Singleton` y se asigna a la variable `instancia`.
5. **Función `main`:** 
    - Obtiene la instancia Singleton y muestra su tiempo.
    - Espera un segundo y obtiene la instancia Singleton de nuevo.
    - Comprueba que ambas instancias son la misma.

### Consideraciones

* **Goroutines:** En Go, el patrón Singleton debe ser utilizado con cuidado en entornos concurrentes (goroutines). Se recomienda utilizar mecanismos de sincronización adicionales para proteger el acceso a la instancia Singleton.
* **Alternativas:** Existen otras alternativas al patrón Singleton en Go, como la inyección de dependencias o la creación de un objeto global de configuración. Se recomienda evaluar las opciones disponibles antes de implementar el patrón Singleton.

**¿Tienes alguna pregunta sobre el patrón Singleton en Go o sobre algún otro tema?** 

### Recursos adicionales

* **Patrones de Diseño en Go:** [https://github.com/tmrts/go-patterns](https://github.com/tmrts/go-patterns)
* **Patrón Singleton en Go:** [https://refactoring.guru/design-patterns/singleton/go/example](https://refactoring.guru/design-patterns/singleton/go/example)
* **Variables globales en Go:** [https://linuxhint.com/golang-global-variables/](https://linuxhint.com/golang-global-variables/)