import { useContext, useEffect, useState } from 'react';
import { Input } from './components/Input';
import { Todo } from './components/Todo';
import './styles/App.css';
import { Auth } from './components/Auth';
import { getUser } from './shared/api/user';
import { getAllTodos } from './shared/api/todo';
import { Toast } from './components/Toast'

interface Todo {
  id: number;
  title: string;
  done: boolean;
}

interface IToast {
  id: number;
  type: number;
  content: string;
}

function App() {
  const [active, setActive] = useState(false);
  const [name, setName] = useState("");
  const [loadingTodos, setLoadingTodos] = useState(true);
  const [loadingUser, setLoadingUser] = useState(true);
  const [todos, setTodos] = useState([]);
  const [toasts, setToasts] = useState<IToast[]>([]);

  async function fetchUser() {
    const token = localStorage.getItem("token");
    if (!token) return;
    setLoadingUser(true);

    const user = await getUser(token);
    setName(user.data.name);
    setLoadingUser(false);
  }

  async function fetchTodos() {
    const token = localStorage.getItem("token");
    if (!token) return;
    setLoadingTodos(true);

    const todos = await getAllTodos(token);
    setTodos(todos.data || []);
    setLoadingTodos(false);
  }

  async function reloadData(user = true, todos = true) {
    try {
      await Promise.all([user ? fetchUser() : {}, todos ? fetchTodos() : {}])
    } catch (err) {
      setLoadingTodos(false);
      setLoadingUser(false);
      console.log(err);
    }
  }

  function fireToast(content: string) {
    setToasts([...toasts, { id: Math.floor(Math.random() * 999999), type: 1, content }]);
  }

  function removeToast(id: number) {
    const index = toasts.findIndex((t) => t.id == id);
    const _toasts = toasts;

    if (index != -1) {
      _toasts.splice(index, 1);
    }

    setToasts(_toasts);
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
      <Auth fireToast={fireToast} reloadData={reloadData} active={active} setActive={setActive} />
      <div id="container">
        <div id="header">
          <h2>TodoList</h2>
          <div id="avatar" className={`${loadingUser || name.length > 1 ? "name" : ""}`} onClick={() => setActive(true)}>
            {loadingUser ? <img src="loader.svg" width={25} height={25} /> : name.length < 1 ? <img src="account-placeholder.png" width={45} height={45} /> : name[0].toUpperCase()}
          </div>
        </div>
        <Input fireToast={fireToast} reloadData={reloadData} />
        <ul id="items">
          {todos.map((todo: Todo) => (
            <Todo fireToast={fireToast} reloadData={reloadData} done={todo.done} id={todo.id} title={todo.title} key={todo.id} />
          ))}
        </ul>
      </div>
      <div id="toasts">
        {toasts.map((toast) => (
          <Toast key={toast.id} removeToast={removeToast} id={toast.id} content={toast.content} />
        ))}
      </div>
    </>
  );
}

export default App;
