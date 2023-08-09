import { useState } from 'react';
import { Input } from './components/Input';
import { Todo } from './components/Todo';
import './styles/App.css';
import { Auth } from './components/Auth';

function App() {
  const [active, setActive] = useState(false);

  return (
    <>
      <Auth active={active} setActive={setActive} />
      <div id="container">
        <div id="header">
          <h2>TodoList</h2>
          <div id="avatar" onClick={() => setActive(true)}>
            <img src="account-placeholder.png" width={45} height={45} />
          </div>
        </div>
        <Input />
        <ul id="items">
          <Todo />
          <Todo />
        </ul>
      </div>
    </>
  );
}

export default App;
