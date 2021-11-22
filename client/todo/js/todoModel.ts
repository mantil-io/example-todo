/*jshint quotmark:false */
/*jshint white:false */
/*jshint trailing:false */
/*jshint newcap:false */

/// <reference path="./interfaces.d.ts"/>

// Generic "model" object. You can use whatever
// framework you want. For this application it
// may not even be worth separating this logic
// out, but we do this to demonstrate one way to
// separate out parts of your application.
class TodoModel implements ITodoModel {

  public key : string;
  public todos : Array<ITodo>;
  public onChanges : Array<any>;

  constructor(key) {
    this.key = key;
    this.todos = [];
    this.onChanges = [];
    this.inform();
  }

  private async apiCall(method: string, body?: any) {
    const url = `/todo/${method}`;
    const opts: any = {
      method: 'POST',
    }
    if (body) {
      opts.body = JSON.stringify(body);
    }
    const rsp = await fetch(url, opts);
    try {
      return await rsp.json();
    } catch {
      return null
    }
  }

  public subscribe(onChange) {
    this.onChanges.push(onChange);
  }

  public async inform() {
    const rsp = await this.apiCall('get');
    this.todos = rsp.Todos;
    this.onChanges.forEach(function (cb) { cb(); });
  }

  public async addTodo(title : string) {
    await this.apiCall('add', { title });
    this.inform();
  }

  public async toggleAll(checked : Boolean) {
    await this.apiCall('toggleAll', { val: checked });
    this.inform();
  }

  public async toggle(todoToToggle : ITodo) {
    await this.apiCall('toggle', { id: todoToToggle.id, val: !todoToToggle.completed });
    this.inform();
  }

  public async destroy(todo : ITodo) {
    await this.apiCall('destroy', { id: todo.id });
    this.inform();
  }

  public async save(todoToSave : ITodo, text : string) {
    await this.apiCall('save', { id: todoToSave.id, title: text });
    this.inform();
  }

  public async clearCompleted() {
    await this.apiCall('clearCompleted');
    this.inform();
  }
}

export { TodoModel };
