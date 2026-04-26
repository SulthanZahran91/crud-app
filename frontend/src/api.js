const BASE = '/todos'

export async function fetchTodos() {
  const res = await fetch(BASE)
  if (!res.ok) throw new Error('fetch failed')
  return res.json()
}

export async function createTodo(title) {
  const res = await fetch(BASE, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ title }),
  })
  if (!res.ok) throw new Error('create failed')
  return res.json()
}

export async function updateTodo(id, title, completed) {
  const res = await fetch(`${BASE}/${id}`, {
    method: 'PUT',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ title, completed }),
  })
  if (!res.ok) throw new Error('update failed')
  return res.json()
}

export async function deleteTodo(id) {
  const res = await fetch(`${BASE}/${id}`, { method: 'DELETE' })
  if (!res.ok) throw new Error('delete failed')
}
