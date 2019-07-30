import React, {Component} from 'react';


class TodoList extends Component {
    componentDidUpdate() {
        this.props.inputElement.current.focus()
    }
    /* Real view of TODOLIST */
    render() {
        return (
            <div className="todoListMain">
                <div className="header">
                    <h1 className="title">TODO LIST</h1>
                    <form onSubmit={this.props.addItem}>
                        <input
                            placeholder="Tâches"
                            ref={this.props.inputElement}
                            value={this.props.currentItem.text}
                            onChange={this.props.handleInput}
                        />
                        <button className="button  button4" type="submit"> Ajouter </button>
                    </form>
                </div>
            </div>
        )
    }
}

export default TodoList
