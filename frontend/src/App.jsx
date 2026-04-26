import { useEffect, useState } from 'react'
import { fetchTodos, createTodo, updateTodo, deleteTodo } from './api.js'
import TodoForm from './components/TodoForm.jsx'
import TodoList from './components/TodoList.jsx'

export default function App() {
  const [todos, setTodos] = useState([])
  const [error, setError] = useState(null)

  useEffect(() => {
    load()
  }, [])

  async function load() {
    try {
      setTodos(await fetchTodos())
      setError(null)
    } catch {
      setError('Failed to load todos')
    }
  }

  async function handleAdd(title) {
    try {
      const todo = await createTodo(title)
      setTodos(prev => [...prev, todo])
      setError(null)
    } catch {
      setError('Failed to add todo')
    }
  }

  async function handleUpdate(id, title, completed) {
    try {
      const updated = await updateTodo(id, title, completed)
      setTodos(prev => prev.map(t => (t.id === id ? updated : t)))
      setError(null)
    } catch {
      setError('Failed to update todo')
    }
  }

  async function handleDelete(id) {
    try {
      await deleteTodo(id)
      setTodos(prev => prev.filter(t => t.id !== id))
      setError(null)
    } catch {
      setError('Failed to delete todo')
    }
  }

  return (
    <div style={{ maxWidth: '600px', margin: '40px auto', padding: '0 16px', fontFamily: 'sans-serif' }}>
      <h1 style={{ marginBottom: '24px' }}>Todos</h1>
      {error && (
        <p style={{ color: '#c00', marginBottom: '16px' }}>{error}</p>
      )}
      <TodoForm onAdd={handleAdd} />
      <TodoList todos={todos} onUpdate={handleUpdate} onDelete={handleDelete} />
    </div>
  )
}
