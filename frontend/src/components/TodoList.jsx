import { useState } from 'react'

export default function TodoList({ todos, onUpdate, onDelete }) {
  const [editId, setEditId] = useState(null)
  const [editTitle, setEditTitle] = useState('')

  function startEdit(todo) {
    setEditId(todo.id)
    setEditTitle(todo.title)
  }

  function cancelEdit() {
    setEditId(null)
    setEditTitle('')
  }

  function saveEdit(todo) {
    const trimmed = editTitle.trim()
    if (!trimmed) return
    onUpdate(todo.id, trimmed, todo.completed)
    cancelEdit()
  }

  if (todos.length === 0) {
    return <p style={{ color: '#888' }}>No todos yet. Add one above.</p>
  }

  return (
    <ul style={{ listStyle: 'none', padding: 0, margin: 0 }}>
      {todos.map(todo => (
        <li
          key={todo.id}
          style={{
            display: 'flex',
            alignItems: 'center',
            gap: '8px',
            padding: '8px 0',
            borderBottom: '1px solid #eee',
          }}
        >
          <input
            type="checkbox"
            checked={todo.completed}
            onChange={e => onUpdate(todo.id, todo.title, e.target.checked)}
          />
          {editId === todo.id ? (
            <>
              <input
                type="text"
                value={editTitle}
                onChange={e => setEditTitle(e.target.value)}
                onKeyDown={e => {
                  if (e.key === 'Enter') saveEdit(todo)
                  if (e.key === 'Escape') cancelEdit()
                }}
                style={{ flex: 1, padding: '4px', fontSize: '16px' }}
                autoFocus
              />
              <button onClick={() => saveEdit(todo)} style={{ cursor: 'pointer' }}>Save</button>
              <button onClick={cancelEdit} style={{ cursor: 'pointer' }}>Cancel</button>
            </>
          ) : (
            <>
              <span
                style={{
                  flex: 1,
                  textDecoration: todo.completed ? 'line-through' : 'none',
                  color: todo.completed ? '#aaa' : '#000',
                  fontSize: '16px',
                }}
              >
                {todo.title}
              </span>
              <button onClick={() => startEdit(todo)} style={{ cursor: 'pointer' }}>Edit</button>
              <button
                onClick={() => onDelete(todo.id)}
                style={{ cursor: 'pointer', color: '#c00' }}
              >
                Delete
              </button>
            </>
          )}
        </li>
      ))}
    </ul>
  )
}
