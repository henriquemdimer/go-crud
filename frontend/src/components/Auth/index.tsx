import { FormEvent, useState } from "react";
import { Button } from "../Button";
import { Modal } from "../Modal";

import "./style.css";
import { Login, createUser } from "../../shared/api/user";

export interface IProps {
    active: boolean;
    setActive: (state: boolean) => void;
    reloadData: (...args: any[]) => Promise<void>;
}

export function Auth(props: IProps) {
    const [method, setMethod] = useState("login");
    const [loading, setLoading] = useState(false);

    async function createAccount(e: FormEvent) {
        setLoading(true);
        e.preventDefault();
        const username = document.getElementById("register__username") as HTMLInputElement;
        const password = document.getElementById("register__password") as HTMLInputElement;
        const rpassword = document.getElementById("register__rpassword") as HTMLInputElement;

        if ((!username || !password || !rpassword)) return console.log("missing input");
        if (!username.value || !password.value || !rpassword.value) return console.log("missing values");
        if (password.value !== rpassword.value) return console.log("password not match");

        try {
            const { data } = await createUser(username.value, password.value);
            localStorage.setItem("token", data.token);
            props.setActive(false);
            await props.reloadData();
        } catch (err) {
            console.log(err);
        } finally {
            setLoading(false);
        }
    }

    async function getAccount(e: FormEvent) {
        setLoading(true);
        e.preventDefault();
        const username = document.getElementById("login__username") as HTMLInputElement;
        const password = document.getElementById("login__password") as HTMLInputElement;

        if (!username || !password) return console.log("missing inputs");
        if (!username.value || !username.value) return console.log("missing values");

        try {
            const { data } = await Login(username.value, password.value);
            localStorage.setItem("token", data.token);
            props.setActive(false);
            await props.reloadData();
        } catch (err) {
            console.error(err);
        } finally {
            setLoading(false);
        }
    }

    return (
        <>
            <Modal active={props.active} setActive={props.setActive}>
                {method == "login" ? (
                    <form onSubmit={(e) => getAccount(e)} id="login" className="auth slideleft">
                        <h3>Entrar na sua conta</h3>
                        <div className="auth__inputs">
                            <div>
                                <label>Nome de usuario</label>
                                <input required id="login__username" />
                            </div>
                            <div>
                                <label>Senha</label>
                                <input required type="password" id="login__password" />
                            </div>
                            <Button loading={loading} type="submit" label='Logar' />
                            <div className="auth__change">
                                <p>Não possui uma conta? <span onClick={() => setMethod("register")}>Crie uma</span></p>
                            </div>
                        </div>
                    </form>
                ) : (
                    <form onSubmit={(e) => createAccount(e)} id="register" className="auth slideright">
                        <h3>Crie uma conta</h3>
                        <div className="auth__inputs">
                            <div>
                                <label>Nome de usuario</label>
                                <input required id="register__username" />
                            </div>
                            <div>
                                <label>Senha</label>
                                <input required type="password" id="register__password" />
                            </div>
                            <div>
                                <label>Repetir senha</label>
                                <input required type="password" id="register__rpassword" />
                            </div>
                            <Button loading={loading} type="submit" label='Criar conta' />
                            <div className="auth__change">
                                <p>Já possui uma conta? <span onClick={() => setMethod("login")}>Entre nela</span></p>
                            </div>
                        </div>
                    </form>
                )}
            </Modal>
        </>
    )
}