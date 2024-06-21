## Explicación del código Golang para caché

### Paquetes

- `fmt`: Sirve para imprimir texto formateado en la consola.
- `log`: Sirve para registrar mensajes (en este caso, errores).
- `time`: Sirve para medir el tiempo transcurrido.

### Estructuras de datos

#### Estructura `Memory`

- `f`: Almacena la función original (`GetFibonacci` en este ejemplo) que calcula el valor.
- `cache`: Un `map[int]FunctionResult` para almacenar los resultados en caché. La clave es el argumento de la función (`int`) y el valor es una estructura `FunctionResult`.

#### Tipo `Function`

- Un tipo de función que recibe un entero `key` como entrada y devuelve una interfaz{} que representa el resultado y un error (`error`). Esto permite que la caché funcione con varios tipos de datos.

#### Estructura `FunctionResult`

- `value`: Almacena el resultado real de la llamada a la función.
- `err`: Almacena cualquier error que haya ocurrido durante la llamada a la función.

### Funciones

#### `NewCache`

- Toma una `Function` como entrada.
- Crea y devuelve una nueva estructura `Memory` con la función proporcionada y un mapa `cache` vacío.

#### `(m *Memory) Get`

(Método en la estructura `Memory`)

- Toma un entero `key` como entrada.
- Comprueba si la `key` existe en el mapa `cache`.
    - Si existe (`exists` es verdadero):
        - Devuelve el `value` y `err` en caché del correspondiente `FunctionResult`.
    - Si no existe (`exists` es falso):
        - Llama a la función original (`m.f`) con la `key`.
        - Almacena el `value` y `err` devueltos en una nueva estructura `FunctionResult`.
        - Añade el par `key`-`FunctionResult` al mapa `cache`.
        - Devuelve el `value` y `err` del `FunctionResult` recién creado.

#### `GetFibonacci`

- Esta función implementa la lógica de cálculo de Fibonacci.
- Toma un entero `n` como entrada.
- Calcula el n-ésimo número de Fibonacci y lo devuelve como una interfaz{}.
- En un escenario real, esta función podría obtener datos de una fuente externa o realizar un cálculo complejo.

#### `Fibonacci` (recursivo)

- Esta función recursiva calcula el n-ésimo número de Fibonacci.
- Caso base: Si `n` es menor o igual que 1, devuelve `n`.
- Caso recursivo: Devuelve la suma de `Fibonacci(n-1)` y `Fibonacci(n-2)`.

### Función `main`

- Crea una nueva caché utilizando la función `GetFibonacci`.
- Define una lista `fibo` que contiene claves enteras (números de Fibonacci) para recuperar.
- Recorre la lista `fibo`:
    - Registra la hora de inicio usando `time.Now()`.
    - Llama a `cache.Get(n)` para recuperar el número de Fibonacci para cada `n`.
    - Imprime el `n`, el `value` recuperado y el tiempo transcurrido desde el inicio.
    - Cualquier error se registra utilizando `log.Println(err)`.

### En resumen

Este código:

1. Define una estructura de caché para almacenar los resultados de una función en función de sus argumentos.
2. Utiliza un mapa para almacenar los resultados en caché para una recuperación eficiente.
3. Proporciona un método `Get` para comprobar la caché en busca de una clave y recuperar el valor si está disponible.
4. Si el valor no está en caché, llama a la función original, almacena el resultado y lo devuelve.
5. Demuestra cómo se puede utilizar la caché para mejorar el rendimiento de una llamada a una función (cálculo de Fibonacci en este ejemplo).

### Puntos importantes para principiantes

- La caché **aumenta el rendimiento** al evitar recalcular resultados costosos.
- La caché **reduce el tiempo de respuesta** para las operaciones repetitivas.
- La caché **es útil** para cálculos lentos o costosos, como la recuperación de datos de una base de datos o API.
- La caché **debe usarse con cuidado** para evitar almacenar datos obsoletos.


Este código en Go calcula números de Fibonacci y usa una caché para hacer los cálculos más rápidos. Imagina la caché como una "memoria" que guarda resultados anteriores para no tener que volver a calcularlos.

¿Qué hace cada parte?

Memory: Es como una caja que guarda la función para calcular Fibonacci (f) y un diccionario (cache) para almacenar los resultados.

Function: Es simplemente un tipo de función que recibe un número y devuelve el resultado y un posible error.

FunctionResult: Es una estructura que guarda el resultado del cálculo (value) y cualquier error que haya ocurrido (err).

NewCache: Crea una nueva "caja" Memory con la función de Fibonacci y un diccionario vacío.

Get: Este método de Memory busca el resultado en la caché. Si no lo encuentra, lo calcula usando la función f, lo guarda en la caché y lo devuelve.

GetFibonacci: Esta función simplemente calcula el número de Fibonacci y devuelve el resultado (y un error nulo, ya que no esperamos errores aquí).

Fibonacci: La famosa función recursiva que calcula el número de Fibonacci.

main: La función principal del programa:

Crea una caché para Fibonacci.
Pide al usuario que ingrese números separados por espacios.
Calcula el Fibonacci de cada número usando la caché y muestra el resultado y el tiempo que tardó.
En resumen: El código calcula números de Fibonacci de manera eficiente usando una caché para evitar cálculos repetidos. ¡Es como tener una calculadora con memoria!