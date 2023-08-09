import { useEffect, useState } from "react"
import "./style.css"
import { deleteTodo, setTodoDone } from "../../shared/api/todo";

export interface IProps {
    title: string;
    id: number;
    done: boolean;
    reloadData: (...args: any[]) => Promise<void>;
    fireToast: (content: string) => void;
}

export function Todo(props: IProps) {
    const [done, setDone] = useState(props.done)
    const [active, setActive] = useState(false);

    useEffect(() => {
        setTimeout(() => {
            setActive(true);
        }, 100);
    }, []);

    async function removeTodo() {
        const token = localStorage.getItem("token");
        if (!token) return;

        try {
            await deleteTodo(props.id, token);
            setActive(false);

            setTimeout(async () => {
                await props.reloadData(false);
            }, 500);
        } catch (err) {
            props.fireToast("Houve um erro ao remover o todo, tente novamente mais tarde!");
            console.log(err);
        }
    }

    async function markDone() {
        const token = localStorage.getItem("token");
        if (!token) return;

        try {
            setDone(!done);
            await setTodoDone(props.id, !done, token);
        } catch (err) {
            console.log(err);
        }
    }

    return (
        <>
            <li className={`todo ${done ? "line" : ""} ${active ? "active" : ""}`}>
                <div onClick={() => markDone()}>
                    <label className="checkbox">
                        <input type="checkbox" onChange={() => { }} checked={done} />
                        <span className="checkmark" />
                    </label>
                    <p>{props.title}</p>
                </div>
                <div onClick={() => removeTodo()} className='x'></div>
            </li>
        </>
    )
}