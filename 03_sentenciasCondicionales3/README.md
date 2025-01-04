# Ejercicios de Programación

## 1. Verificar si un estudiante está aprobado
Programa que determine si un estudiante está aprobado y muestre su nota final con medio punto a favor del alumno.

### Condiciones:
- Si el laboratorio está aprobado, la nota final se calcula con la fórmula:
  \[ NF = 0.75 \times NT + 0.25 \times NL \]
- En cualquier otro caso, se asigna la nota mínima 

- nota aprobatoria es: 13.

---

## 2. Operaciones con dos números enteros

### Condiciones:
- Si los números son iguales, multiplicarlos.
- Si el primero es mayor, restarlos.
- Si no, sumarlos.

---

## 3. Número mayor entre cuatro números
Determinar el mayor número de entre cuatro números reales sin usar `Math.max()`.

---

## 4. Utilidades por antigüedad
Determinar la utilidad anual que recibe un trabajador basada en su antigüedad.

### Tabla de utilidad:
| Tiempo | Utilidad |
|--------|----------|
| Menos de 1 año | 5% del salario |
| 1 a 2 años | 7% del salario |
| 2 a 5 años | 10% del salario |
| 5 a 10 años | 15% del salario |
| 10 o más años | 20% del salario |

---

## 5. Descuento en un supermercado
Determinar el monto final que debe pagar un cliente basándose en el color de una bolita.

### Descuentos por color:
| Color | Descuento |
|-------|-----------|
| Blanco | 0% |
| Verde | 10% |
| Amarillo | 25% |
| Azul | 50% |
| Rojo | 100% |

---

## 6. Clasificación para jubilación
Determinar en qué tipo de jubilación está adscrita una persona según las siguientes condiciones:

### Tipos de jubilación:
- **Por edad**: 60 años o más y antigüedad menor a 25 años.
- **Por antigüedad joven**: Menos de 60 años y antigüedad de 25 años o más.
- **Por antigüedad adulta**: 60 años o más y antigüedad de 25 años o más.
- **No accede a jubilación**: En cualquier otro caso.

---

## 7. Calificar una nota en palabras
Asignar una calificación en palabras a una nota según el siguiente criterio:

### Tabla de calificación:
| Rango de notas | Calificación |
|----------------|---------------|
| 0.0 a 4.9 | Reprobado, repite el semestre |
| 5.0 a 10.4 | Reprobado, pasa a subsanación |
| 10.5 a 15.9 | Aprobado |
| 16.0 a 20.0 | Aprobado con distinción máxima |

---

## 8. Ecuación cuadrática
Programa que muestre las raíces de una ecuación cuadrática y que muestre mensajes en caso de problemas. Considerar también el caso de solución lineal.

---

## 9. Determinar si un año es bisiesto
Condiciones:
- Es divisible por 4 pero no por 100, o es divisible por 400.

### Ejemplos:
| Año | Resultado |
|------|-----------|
| 1998 | No |
| 1996 | Sí |
| 1900 | No |
| 2000 | Sí |

---

## 10. Suma de fracciones
Calcular la suma de dos fracciones usando la fórmula:
\[ \frac{a}{b} + \frac{c}{d} = \frac{ad + bc}{bd} \]

### Notas:
- Validar que el denominador no sea cero.
- Si el resultado es entero, mostrarlo como fracción sin simplificar.
