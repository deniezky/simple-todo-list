# Simple TodoList

На данный момент, это мой первый проект, написанный на языке **Go**

## Функция Add
```go
func Add(title string) Todo
```

> Функция ***Add*** добавляет новую задачу в массив ***todos*** и возвращает объект ***Todo***.

## Функция Remove
```go
func Remove(id int) Todo
```
> Функция ***Remove*** удаляет задачу с указанным идентификатором из массива ***todos*** и возвращает удаленную задачу.


## Функция Complete (Todo)

```go
func (t *Todo) Complete()
```
> Меняет ***boolean*** значение для элемента ***Completed***  **(true/false)**

## Функция Edit (Todo)

```go
func (t *Todo) Edit(title string)
```

> Функция ***Edit*** редактирует объект, изменяя его ***title***.


# Структура Todo

```go
type Todo struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}
```
