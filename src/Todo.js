import React from 'react'

 const Todo = () => {
  return (
    <div  className="todo">
      <form className="todo-list">
        <input
        className="input"
        type="text"
        name="text"
        placeholder="add an item"
        />

        <button type="submit" name="submit" value="submit" >  Add</button>
      </form>
    </div>
  )
}

export default Todo
