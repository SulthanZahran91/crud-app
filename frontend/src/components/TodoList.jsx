import { useState } from 'react'

const buttonStyle = {
  padding: '8px 14px',
  minHeight: '44px',
  cursor: 'pointer',
  border: '1px solid #ccc',
  background: '#fff',
  borderRadius: '8px',
  fontSize: '14px',
}

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
    return <p style={{ color: '#888', padding: '16px 0' }}>No todos yet. Add one above.</p>
  }

  return (
    <ul style={{ listStyle: 'none', padding: 0, margin: 0 }}>
      {todos.map(todo => (
        <li
          key={todo.id}
          style={{
            display: 'flex',
            alignItems: 'center',
            gap: '10px',
            padding: '10px 0',
            borderBottom: '1px solid #eee',
            flexWrap: 'wrap',
          }}
        >
          <input
            type="checkbox"
            checked={todo.completed}
            onChange={e => onUpdate(todo.id, todo.title, e.target.checked)}
            style={{ width: '20px', height: '20px', flexShrink: 0 }}
            aria-label={`Mark "${todo.title}" ${todo.completed ? 'incomplete' : 'complete'}`}
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
                style={{
                  flex: '1 1 160px',
                  padding: '8px',
                  fontSize: '16px',
                  minHeight: '44px',
                  border: '1px solid #ccc',
                  borderRadius: '8px',
                  boxSizing: 'border-box',
                }}
                aria-label="Edit todo title"
                autoFocus
              />
              <button onClick={() => saveEdit(todo)} style={buttonStyle}>Save</button>
              <button onClick={cancelEdit} style={buttonStyle}>Cancel</button>
            </>
          ) : (
            <>
              <span
                style={{
                  flex: '1 1 160px',
                  textDecoration: todo.completed ? 'line-through' : 'none',
                  color: todo.completed ? '#aaa' : '#1a1a1a',
                  fontSize: '16px',
                  lineHeight: 1.4,
                }}
              >
                {todo.title}
              </span>
              <button onClick={() => startEdit(todo)} style={buttonStyle}>Edit</button>
              <button
                onClick={() => onDelete(todo.id)}
                style={{ ...buttonStyle, color: '#b3261e', borderColor: '#f5c6c2' }}
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
