## About

This is a simple Todo application which will demonstrate how to use persistent key/value storage in Mantil APIs.

## Prerequisites

This template is created with Mantil. To download [Mantil CLI](https://github.com/mantil-io/mantil#installation) on Mac or Linux use Homebrew 
```
brew tap mantil-io/mantil
brew install mantil
```
or check [direct download links](https://github.com/mantil-io/mantil#installation).

To deploy this application you will need:
- An [AWS account](https://aws.amazon.com/premiumsupport/knowledge-center/create-and-activate-aws-account/)
- A GitHub account with a repository where you have admin rights
- A Slack account with the right to create apps

## Installation

To locally create a new project from this template run:
```
mantil new app --from todo
cd app
```

## Using the K/V store

In this example we are using a Mantil KV store to persistently store and query todos.

First we [initialize](api/todo/todo.go#L18) the store:
```
kv, _ := mantil.NewKV("todos")
```
This will create a new DynamoDB table for the stage (if it doesn't already exist) and a `todos` partition on it.

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

## Deploying the application

Note: If this is the first time you are using Mantil you will firstly need to install Mantil Node on your AWS account. For detailed instructions please follow these simple, one-step setup [instructions](https://github.com/mantil-io/mantil/blob/master/docs/getting_started.md#setup)
```
mantil aws install
```
After configuring the environment variable you can proceed with the creation of the first stage.
```
mantil stage new development
```

This command will create a new stage called `development` and deploy it to your node. 
(After configuring the environment variable you can proceed with application deployment.
```
mantil deploy
```

This command will create a new stage for your project with default name `development` and deploy it to your node.)

Now you can output the stage endpoint with `mantil env -u`. The API endpoint for your function will have the name of that function in the path, in our case that is `$(mantil env -u)/handler`.

With this URL we can now create a Github webhook which will invoke our Lambda function on each star to our Github repository.


## Modification

If you want different behavior out of your function you can make necessary changes to your code the `api` folder.

The client code is located in `client/todo`, to build it run:

```
cd client/todo
npm install
npm run build
```

This will build the static assets and copy them over to the Mantil public folder. Note that this is optional and only needed if you want to modify the client code. The project already contains prebuilt assets in `public` so you can start by deploying a new stage immediately.

After each change you have to deploy your changes with `mantil deploy`, or instruct Mantil to  automatically deploy all saved changes with `mantil watch`.

For more detailed instruction please refer to the [Mantil documentation](https://github.com/mantil-io/mantil#documentation).

## Cleanup

To remove the created stage from your AWS account destroy it with:
```
mantil stage destroy development
```

## Final thoughts

With this template you learned how to create a simple serverless todo application with AWS Lambda and Mantil's KV store backed by DynamoDB. Check out [our documentation](https://github.com/mantil-io/mantil#documentation) to find more interesting templates.

If you have any questions or comments on this concrete template or would just like to share your view on Mantil contact us at [support@mantil.com](mailto:support@mantil.com) or create an issue.