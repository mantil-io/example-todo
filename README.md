## About

This is a simple Todo application which will demonstrate how to use persistent key/value storage in Mantil APIs.

## Getting started

To use this template, simply create a new Mantil project with the `--from=todo` flag:
```
mantil new todo-app --from=todo
cd todo-app
```

To deploy the project, run:
```
mantil stage new development
```

Now the project website will be available at the root URL for the stage which you can obtain by running:

```
mantil env -u
```

## Client

The client code is located in `client/todo`, to build it run:

```
cd client/todo
npm install
npm run build
```

This will build the static assets and copy them over to the Mantil public folder. Note that this is optional and only needed if you want to modify the client code. The project already contains prebuilt assets in `public` so you can start by deploying a new stage immediately.

## Using the K/V storage

In this example we are using a Mantil KV store to persistently store and query todos.

First we [initialize](api/todo/todo.go#L18) the store:
```
kv, _ := mantil.NewKV("todos")
```
This will create a new dynamodb table for the stage (if it doesn't already exist) and a `todos` partition on it.

Now we can use the `Put` method to [add](api/todo/add.go#L15) a new todo:
```
id := uuid.NewString()
kv.Put(id, &TodoItem{
    ID:        id,
    Title:     "Do the laundry",
    Completed: false,
})
```

To [fetch all](api/todo/get.go#L13) todos, we can use the `FindAll` method:
```
var items []TodoItem
_, err := kv.FindAll(&items)
```

For more complex queries we can use in-memory filtering (such as in the [toggleAll example](api/todo/toggleAll.go#L12)).
Alternatively, we could define a more suitable sort key. For example, instead of just using the generated `uuid` as the key as above, we could define:
```
func (t *TodoItem) kvKey() string {
    return fmt.Sprintf("completed:%v-%s", t.Completed, t.ID)
}
...
t := &TodoItem{
    ID:        uuid.NewString(),
    Title:     "Do the laundry",
    Completed: false,
}
kv.Put(t.kvKey(), t)
```

Now we can query all completed todos by calling:
```
var items []TodoItem
_, err := kv.Find(&items, mantil.FindBeginsWith, "completed:true")
```

## Cleanup

To destroy the created stage, run:
```
mantil stage destroy development
```

This will destroy all AWS resources associated with the project. Now you can
safely delete the project folder.