import { FormEvent, useState } from "react";
import "./style.css";
import { createTodo } from "../../shared/api/todo";

export interface IProps {
    reloadData: (...args: any[]) => Promise<void>;
}

export function Input(props: IProps) {
    const [text, setText] = useState("");
    const [loading, setLoading] = useState(false);

    async function addTodo(e: FormEvent) {
        setLoading(true);
        e.preventDefault();
        try {
            const token = localStorage.getItem("token");
            if (!token) return;

            await createTodo(text, token);
            await props.reloadData(false);

            const btn = document.getElementById("addtodo");
            if (btn) {
                btn.classList.add("active")

                setTimeout(() => {
                    setText("");
                    btn.innerText = "Todo adicionado!"

                    const parent = document.getElementById("items");
                    if (parent) {
                        parent.scroll({
                            top: parent.scrollHeight,
                            behavior: "smooth"
                        });
                    }
                }, 300);

                setTimeout(() => {
                    btn.classList.remove("active");
                    btn.innerText = "+";
                    setLoading(false);
                }, 2000);
            }
        } catch (err) {
            console.log(err);
            setLoading(false);
        }
    }

    return (
        <>
            <form onSubmit={addTodo} id="input">
                <input value={text} onChange={(e) => setText(e.target.value)} placeholder="Fazer algo..." />
                <button disabled={loading} id="addtodo" type="submit">+</button>
            </form>
        </>
    )
}