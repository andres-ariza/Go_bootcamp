Ejercicio 1
¿Cuáles son las diferencias entre White Box y Black Box? 
Respuesta:
- en el blackbox no se conoce el funcionamiento interno, mientras que en el whitebox si
- el blackbox se usa para saber si la app funciona, mientras que el whitebox para saber si los metodos y el flujo de datos se comportan con respecto a lo esperado

Ejercicio 2
¿Qué es un test funcional?.
Respuesta:
- los test funcionales son los aplicados al funcionamiento adecuado de toda la aplicacion

Ejercicio 3
¿Qué es un Test de Integración?
Respuesta:
- los test de integracion comprueban la comunicacion de las diferentes capas o modulos de la aplicacion

Ejercicio 4
Indicar las dimensiones de calidad prioritarias en MELI.
Respuesta:
1. SOLID:
- Un objeto solo debería tener una única responsabilidad.
- Los objetos deben estar abiertos para su extensión, pero cerrados para su modificación.
- Los objetos deberían ser reemplazables por objetos hijos (de sus subtipos) sin alterar el correcto funcionamiento del programa.
- Muchas interfaces más específicas son mejores que una interfaz de propósito general.
- Desacoplar los objetos abstractos de sus implementaciones.
2. FIRST:
- Rápidos: Es posible tener miles de tests en tu proyecto y deben ser rápidos de correr. 
- Aislados/Independientes: Un método de test debe cumplir con los “3A” (Arrange, Act, Assert) o lo que es lo mismo: Given, when, then. Además no debe ser necesario que sean corridos en un determinado orden para funcionar.
- Repetibles: Resultados determinísticos. No deben depender de datos del ambiente mientras están corriendo (por ejemplo: la hora del sistema).
- Auto-Validados: No debe ser requerida una inspección manual para validar los resultados.
- Completos: Deben cubrir cada escenario de un caso de uso, y no sólo buscar un coverage del 100%. Probar mutaciones, edge cases, excepciones, errores, 
