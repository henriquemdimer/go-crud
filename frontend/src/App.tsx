import { useContext, useEffect, useState } from 'react';
import { Input } from './components/Input';
import { Todo } from './components/Todo';
import './styles/App.css';
import { Auth } from './components/Auth';
import { getUser } from './shared/api/user';
import { getAllTodos } from './shared/api/todo';
import { Toast } from './components/Toast'
import { Modal } from './components/Modal';
import { Button } from './components/Button';
import { Dropdown, DropdownContent, DropdownTrigger } from './components/Dropdown';
import { ping } from './shared/api/main';

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
  const [loadingTodos, setLoadingTodos] = useState(false);
  const [loadingUser, setLoadingUser] = useState(false);
  const [todos, setTodos] = useState([]);
  const [toasts, setToasts] = useState<IToast[]>([]);
  const [alert, setAlert] = useState(false);
  const [dropdown, setDropdown] = useState(false);

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
      throw err;
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

  function Logout() {
    localStorage.removeItem("token")
    setName("");
    setTodos([]);
    setDropdown(false);
    fireToast("Você saiu da conta com sucesso!")
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

      try {
        await ping()
      } catch (_) {
        setAlert(true);
        const interval = setInterval(async () => {
          if (!alert) clearInterval(interval);

          try {
            await ping();
            setAlert(false);
            fireToast("Conexão reestabelecida!")
          } catch { }
        }, 5000);
      }
    }

    initialFetch();
  }, []);

  return (
    <>
      <Modal blockClose={true} active={alert} setActive={setAlert}>
        <div id="outage">
          <img src="outage.png" width={100} height={100} />
          <p>Parece que o servidor está offline, tente novamente dentro de alguns minutos!</p>
          <Button loading={loadingUser} disabled={loadingUser} onClick={() => ping().then(() => { setAlert(false); fireToast("Conexão reestabelecida!") }).catch(() => fireToast("Não foi possivel se comunicar com o servidor!"))} label='Tentar novamente' />
        </div>
      </Modal>
      <Auth fireToast={fireToast} reloadData={reloadData} active={active} setActive={setActive} />
      <div id="container">
        <div id="header">
          <h2>TodoList</h2>
          <Dropdown active={dropdown} setActive={setDropdown}>
            <DropdownTrigger onClick={() => name && setDropdown(true)}>
              <div id="avatar" className={`${loadingUser || name.length > 1 ? "name" : ""}`} onClick={() => name.length ? {} : setActive(true)}>
                {loadingUser ? <img src="loader.svg" width={25} height={25} /> : name.length < 1 ? <img src="account-placeholder.png" width={45} height={45} /> : name[0].toUpperCase()}
              </div>
            </DropdownTrigger>
            <DropdownContent>
              <p className='user'>{name}</p>
              <hr></hr>
              <p id="logout" onClick={() => Logout()}>Logout</p>
            </DropdownContent>
          </Dropdown>
        </div>
        <Input setModal={setActive} fireToast={fireToast} reloadData={reloadData} />
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
