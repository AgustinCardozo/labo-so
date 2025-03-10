package list

import "fmt"

// List Definir la interfaz List
type List[T any] interface {
	Add(item T)                                          // Añadir un elemento al final de la lista
	Filter(value T, predicate func(a, b T) bool) List[T] // Filtra elementos de la lista
	Get(index int) (T, error)                            // Obtener un elemento a partir de un índice dado
	Insert(index int, item T) error                      // Insertar un elemento en el índice dado
	Pop() (T, error)                                     // Remover el último elemento de la lista
	Remove(index int)                                    // Eliminar un elemento en el índice dado
	Size() int                                           // Retornar el tamaño de la lista
	Sort(less func(a, b T) bool)                         // Sort Ordena una Lista de acuerdo al criterio
}

// ArrayList implements List
type ArrayList[T any] struct {
	items []T
}

/*
* Add: Inserta un elemento al final de la lista.
* @param item: Elemento a insertar.
 */
func (list *ArrayList[T]) Add(item T) {
	list.items = append(list.items, item)
}

/*
* Filter: Filtra elementos de la lista en base a un predicado.
* @param value: Valor a comparar.
* @param predicate: Función que compara dos elementos.
 */
func (list *ArrayList[T]) Filter(value T, predicate func(a, b T) bool) List[T] {
	filteredList := &ArrayList[T]{}

	for _, item := range list.items {
		if predicate(value, item) {
			filteredList.Add(item)
		}
	}

	return filteredList
}

/*
* Get: Devuelve el elemento en el índice proporcionado.
* @param index: Índice del elemento a obtener.
 */
func (list *ArrayList[T]) Get(index int) (T, error) {
	// Validar si el índice está dentro del rango
	if index < 0 || index >= len(list.items) {
		// Get item from a List
		var zero T // Crear un valor cero del tipo genérico T
		return zero, fmt.Errorf("index out of range: %d", index)
	}
	return list.items[index], nil
}

/*
* Insert: Inserta un elemento en la lista en el índice proporcionado.
* @param index: Índice donde se va a ingresar el elemento.
* @param item:  Elemento a ingresar.
 */
func (list *ArrayList[T]) Insert(index int, item T) error {
	if index < 0 || index > len(list.items) {
		return fmt.Errorf("index out of range: %d", index)
	}
	list.items = append(list.items[:index], append([]T{item}, list.items[index:]...)...)
	return nil
}

/*
* Pop: Remueve el último elemento de la lista y lo devuelve.
 */
func (list *ArrayList[T]) Pop() (T, error) {
	if len(list.items) == 0 {
		var zero T
		return zero, fmt.Errorf("list is empty")
	}
	lastIndex := len(list.items) - 1
	item := list.items[lastIndex]
	list.items = list.items[:lastIndex]
	return item, nil
}

/*
* Remove: Remueve un elemento de la lista en base a su índice.
* @param list: lista de cualquier tipo.
* @param index: Índice del elemento a remover.
 */
func (list *ArrayList[T]) Remove(index int) {
	if index >= 0 && index < len(list.items) {
		list.items = append(list.items[:index], list.items[index+1:]...)
	}
}

/*
* Size: Devuelve el tamaño de la lista.
 */
func (list *ArrayList[T]) Size() int {
	return len(list.items)
}

/*
* Sort: Ordena una lista de acuerdo a un criterio.
* @param less: Función que compara dos elementos.
 */
func (list *ArrayList[T]) Sort(less func(a, b T) bool) {
	size := list.Size()
	for i := 0; i < size-1; i++ {
		for j := 0; j < size-i-1; j++ {
			if !less(list.items[j], list.items[j+1]) {
				// Intercambiar los elementos si están en el orden incorrecto
				list.items[j], list.items[j+1] = list.items[j+1], list.items[j]
			}
		}
	}
}
