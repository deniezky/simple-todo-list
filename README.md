# Simple TodoList

## Функция Add
```
func Add(title string) Todo
```

> Функция ***Add*** добавляет новую задачу в массив ***todos*** и возвращает объект ***Todo***.

## Функция Remove
```
func Remove(id int) Todo
```
> Функция ***Remove*** удаляет задачу с указанным идентификатором из массива ***todos*** и возвращает удаленную задачу.


## Функция Complete (Todo)

```
func (t *Todo) Complete()
```
> Меняет ***boolean*** значение для элемента ***Completed***  **(true/false)**

## Функция Edit (Todo)

```
func (t *Todo) Edit(title string)
```

> Функция ***Edit*** редактирует объект, изменяя его ***title***.


# Структура Todo

```
type Todo struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}
```
