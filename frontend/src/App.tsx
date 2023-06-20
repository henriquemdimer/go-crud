import { useEffect, useState } from 'react';
import { Input } from './components/Input';
import { Todo } from './components/Todo';
import './styles/App.css';
import { Auth } from './components/Auth';
import { getUser } from './shared/api/user';
import { getAllTodos } from './shared/api/todo';

interface Todo {
  id: number;
  title: string;
  done: boolean;
}

function App() {
  const [active, setActive] = useState(false);
  const [name, setName] = useState("");
  const [loading, setLoading] = useState(true);
  const [todos, setTodos] = useState([]);

  const fetchUser = async () => {
    const token = localStorage.getItem("token");
    if (!token) return;

    const user = await getUser(token);
    setName(user.data.name);
  }

  const fetchTodos = async () => {
    const token = localStorage.getItem("token");
    if (!token) return;

    const todos = await getAllTodos(token);
    setTodos(todos.data || []);
  }

  async function reloadData(user = true, todos = true) {
    setLoading(true);

    try {
      await Promise.all([user ? fetchUser() : {}, todos ? fetchTodos() : {}])
    } finally {
      setLoading(false);
    }
  }

  useEffect(() => {
    const initialFetch = async () => {
      await reloadData();

      setTimeout(() => {
        const items = document.getElementById("items");

        if (items) items.scroll({
          top: items.scrollHeight,
          behavior: "smooth"
        });
      }, 800)
    }

    initialFetch();
  }, []);

  return (
    <>
      <Auth reloadData={reloadData} active={active} setActive={setActive} />
      <div id="container">
        <div id="header">
          <h2>TodoList</h2>
          <div id="avatar" className={`${loading || name.length > 1 ? "name" : ""}`} onClick={() => setActive(true)}>
            {loading ? <img src="loader.svg" width={25} height={25} /> : name.length < 1 ? <img src="account-placeholder.png" width={45} height={45} /> : name[0].toUpperCase()}
          </div>
        </div>
        <Input reloadData={reloadData} />
        <ul id="items">
          {todos.map((todo: Todo) => (
            <Todo reloadData={reloadData} done={todo.done} id={todo.id} title={todo.title} key={todo.id} />
          ))}
        </ul>
      </div>
    </>
  );
}

export default App;
