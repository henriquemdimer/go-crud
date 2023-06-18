import { useState } from 'react';
import { Input } from './components/Input';
import { Modal } from './components/Modal';
import { Todo } from './components/Todo';
import './styles/App.css';
import { Button } from './components/Button';

function App() {
  const [active, setActive] = useState(false);
  const [method, setMethod] = useState("login")

  return (
    <>
      <Modal active={active} setActive={setActive}>
        {method == "login" ? (
          <div id="login" className="auth slideleft">
            <h3>Entrar na sua conta</h3>
            <div className="auth__inputs">
              <div>
                <label>Nome de usuario</label>
                <input />
              </div>
              <div>
                <label>Senha</label>
                <input />
              </div>
              <Button label='Logar' />
              <div className="auth__change">
                <p>Não possui uma conta? <span onClick={() => setMethod("register")}>Crie uma</span></p>
              </div>
            </div>
          </div>
        ) : (
          <div id="register" className="auth slideright">
            <h3>Crie uma conta</h3>
            <div className="auth__inputs">
              <div>
                <label>Nome de usuario</label>
                <input />
              </div>
              <div>
                <label>Senha</label>
                <input />
              </div>
              <div>
                <label>Repetir senha</label>
                <input />
              </div>
              <Button label='Logar' />
              <div className="auth__change">
                <p>Já possui uma conta? <span onClick={() => setMethod("login")}>Entre nela</span></p>
              </div>
            </div>
          </div>
        )}
      </Modal>
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
