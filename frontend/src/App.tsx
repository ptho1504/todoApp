import { useEffect, useState } from "react";

type User = {
  id: number;
  name: string;
};

const API_URL = import.meta.env.VITE_API_URL;

export default function UserPage() {
  const [users, setUsers] = useState<User[]>([]);
  const [name, setName] = useState("");
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState("");

  const fetchUsers = async () => {
    try {
      setLoading(true);
      setError("");

      const res = await fetch(`${API_URL}/users`);
      if (!res.ok) throw new Error("Failed to fetch users");

      const data = await res.json();
      setUsers(data);
    } catch (err: any) {
      setError(err.message);
    } finally {
      setLoading(false);
    }
  };

  const addUser = async () => {
    if (!name.trim()) return;

    try {
      setLoading(true);
      setError("");

      const res = await fetch(API_URL + "/users", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ name }),
      });

      if (!res.ok) throw new Error("Failed to add user");

      setName("");
      await fetchUsers();
    } catch (err: any) {
      setError(err.message);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchUsers();
  }, []);

  return (
    <div style={styles.container}>
      <h1 style={styles.title}>User Management</h1>

      <div style={styles.inputGroup}>
        <input
          style={styles.input}
          value={name}
          onChange={(e) => setName(e.target.value)}
          placeholder="Enter name..."
        />
        <button style={styles.button} onClick={addUser}>
          Add
        </button>
      </div>

      {loading && <p>Loading...</p>}
      {error && <p style={styles.error}>{error}</p>}

      <UserList users={users} />
    </div>
  );
}

/* ------------------ Components ------------------ */

function UserList({ users }: { users: User[] }) {
  if (!users.length) return <p>No users found</p>;

  return (
    <ul style={styles.list}>
      {users.map((u) => (
        <li key={u.id} style={styles.listItem}>
          {u.name}
        </li>
      ))}
    </ul>
  );
}

/* ------------------ Styles ------------------ */

const styles: Record<string, React.CSSProperties> = {
  container: {
    maxWidth: "400px",
    margin: "40px auto",
    padding: "20px",
    borderRadius: "12px",
    boxShadow: "0 4px 12px rgba(0,0,0,0.1)",
    fontFamily: "sans-serif",
  },
  title: {
    marginBottom: "16px",
  },
  inputGroup: {
    display: "flex",
    gap: "8px",
    marginBottom: "16px",
  },
  input: {
    flex: 1,
    padding: "8px",
    borderRadius: "6px",
    border: "1px solid #ccc",
  },
  button: {
    padding: "8px 12px",
    borderRadius: "6px",
    border: "none",
    backgroundColor: "#007bff",
    color: "white",
    cursor: "pointer",
  },
  list: {
    listStyle: "none",
    padding: 0,
  },
  listItem: {
    padding: "8px",
    borderBottom: "1px solid #eee",
  },
  error: {
    color: "red",
  },
};