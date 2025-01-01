# Ejercicios golang

### 1. Determinar si un alumno aprueba o no un curso
Escribe un programa que determine si un alumno aprueba o reprueba un curso basado en su promedio de tres calificaciones. El promedio se calcula considerando números decimales y:
- Aprueba si el promedio es mayor o igual a 13.
- Reprueba en caso contrario.

---

### 2. Imprimir dos números en forma ascendente
Escribe un programa que lea dos números enteros y los imprima en orden ascendente. No se deben usar métodos de la clase `Math`.

---

### 3. Operadores relacionales y de igualdad
Dados dos números enteros, escribe un programa que muestre todos los operadores relacionales y de igualdad válidos entre ellos. Esto incluye:
- Comparaciones como `==`, `!=`, `<`, `>`, `<=`, `>=`.

---

### 4. Cálculo de interés anual
Escribe un programa que permita a una persona consultar cuánto dinero se generará por concepto de interés anual sobre el capital que desea invertir en un banco. Consideraciones:
- La tasa de interés es anual.
- Si los intereses generados exceden $200 al año, el programa recomendará invertir e indicará el monto total al final del año.
- Si no excede $200, indicará que no es rentable invertir.

---

### 5. Evaluación de un estudiante con teoría y laboratorio
Crea un programa que determine si un estudiante aprueba o no, y calcule su nota final basada en las siguientes reglas:
- El estudiante aprueba si sus notas de Teoría y Laboratorio (enteros) son aprobatorias.
- Si aprueba, la nota final es el promedio ponderado (75% Teoría y 25% Laboratorio).
- Si no aprueba, la nota final es la mínima entre sus dos notas (sin usar `Math.min()`).
