import { useState } from 'react'

export default function TodoForm({ onAdd }) {
  const [title, setTitle] = useState('')

  function handleSubmit(e) {
    e.preventDefault()
    const trimmed = title.trim()
    if (!trimmed) return
    onAdd(trimmed)
    setTitle('')
  }

  return (
    <form
      onSubmit={handleSubmit}
      style={{
        display: 'flex',
        gap: '8px',
        marginBottom: '16px',
        flexWrap: 'wrap',
      }}
    >
      <input
        type="text"
        value={title}
        onChange={e => setTitle(e.target.value)}
        placeholder="New todo..."
        aria-label="New todo"
        style={{
          flex: '1 1 200px',
          padding: '12px',
          fontSize: '16px',
          minHeight: '44px',
          border: '1px solid #ccc',
          borderRadius: '8px',
          boxSizing: 'border-box',
        }}
      />
      <button
        type="submit"
        style={{
          padding: '12px 20px',
          minHeight: '44px',
          cursor: 'pointer',
          border: 'none',
          borderRadius: '8px',
          background: '#1a1a1a',
          color: '#fff',
          fontWeight: 600,
        }}
      >
        Add
      </button>
    </form>
  )
}
