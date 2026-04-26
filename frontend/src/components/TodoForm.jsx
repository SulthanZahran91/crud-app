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
    <form onSubmit={handleSubmit} style={{ display: 'flex', gap: '8px', marginBottom: '16px' }}>
      <input
        type="text"
        value={title}
        onChange={e => setTitle(e.target.value)}
        placeholder="New todo..."
        style={{ flex: 1, padding: '8px', fontSize: '16px' }}
      />
      <button type="submit" style={{ padding: '8px 16px', cursor: 'pointer' }}>Add</button>
    </form>
  )
}
