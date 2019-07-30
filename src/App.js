import React, {Component} from 'react';
import './App.css';
import TodoList from './TodoList';
import TodoItems from "./TodoItems";



class App extends Component {
  inputElement = React.createRef();

  /* Constructor */
  constructor() {
      super();
      this.state = {
          items: [],
          currentItem: {
            text: '',
            key: ''
          }
      }
  }
  /* Method to delete task in todolist*/
  deleteItem = key => {
    const filteredItems = this.state.items.filter(item => {
      return item.key !== key
    });
    this.setState({
      items: filteredItems
    })
  };

  handleInput = e => {
    const itemText = e.target.value;
    const currentItem = { text: itemText, key: Date.now() };
    this.setState({
      currentItem,
    })
  };

  /* Add a new task */
  addItem = e => {
    e.preventDefault();
    const newItem = this.state.currentItem;
    if(newItem.text !== '') {
      const items = [...this.state.items, newItem];
      this.setState({
        items: items,
        currentItem: { text: '', key: '' },
      })
    }
  };

  /* View of TodoList */
  render() {
    return (
        <div className="App">
          <TodoList
                    addItem={this.addItem}
                    inputElement={this.inputElement}
                    handleInput={this.handleInput}
                    currentItem={this.state.currentItem}
          />
          <TodoItems entries={this.state.items} deleteItem={this.deleteItem} />
        </div>
    );
  }


}

export default App;
