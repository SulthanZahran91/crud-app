import { useEffect, useState } from 'react'
import { fetchTodos, createTodo, updateTodo, deleteTodo } from './api.js'
import TodoForm from './components/TodoForm.jsx'
import TodoList from './components/TodoList.jsx'

export default function App() {
  const [todos, setTodos] = useState([])
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState(null)

  async function load() {
    setLoading(true)
    setError(null)
    try {
      setTodos(await fetchTodos())
    } catch {
      setError('Failed to load todos')
    } finally {
      setLoading(false)
    }
  }

  useEffect(() => {
    load()
  }, [])

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
    <div
      style={{
        maxWidth: '600px',
        margin: '0 auto',
        padding: '48px 16px',
        minHeight: '100dvh',
        fontFamily:
          'system-ui, -apple-system, "Segoe UI", Roboto, sans-serif',
        color: '#1a1a1a',
      }}
    >
      <h1 style={{ marginBottom: '24px', fontSize: '28px', letterSpacing: '-0.02em' }}>
        Todos
      </h1>
      {error && (
        <div
          role="alert"
          style={{
            color: '#b3261e',
            background: '#fdecea',
            border: '1px solid #f5c6c2',
            borderRadius: '8px',
            padding: '12px 16px',
            marginBottom: '16px',
            display: 'flex',
            alignItems: 'center',
            justifyContent: 'space-between',
            gap: '12px',
            flexWrap: 'wrap',
          }}
        >
          <span>{error}</span>
          <button
            onClick={load}
            style={{
              padding: '8px 16px',
              minHeight: '44px',
              cursor: 'pointer',
              border: '1px solid #b3261e',
              background: 'transparent',
              color: '#b3261e',
              borderRadius: '8px',
              fontWeight: 600,
            }}
          >
            Retry
          </button>
        </div>
      )}
      <TodoForm onAdd={handleAdd} />
      {loading ? (
        <p style={{ color: '#888', padding: '16px 0' }}>Loading todos...</p>
      ) : (
        <TodoList todos={todos} onUpdate={handleUpdate} onDelete={handleDelete} />
      )}
    </div>
  )
}
