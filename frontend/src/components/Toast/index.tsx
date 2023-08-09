import { useEffect, useState } from "react"
import "./style.css"

export interface IProps {
    id: number;
    content: string;
    removeToast: (...args: any[]) => any;
}

export function Toast(props: IProps) {
    const [active, setActive] = useState(false);

    useEffect(() => {
        setTimeout(() => {
            setActive(true);

            setTimeout(() => {
                setActive(false)

                // bug
                // setTimeout(() => {
                //     props.removeToast(props.id);
                // }, 600);
            }, 3000);
        }, 10);

    }, []);

    return (
        <>
            <div className={`toast ${active ? "active" : ""}`}>
                <span>{props.content}</span>
            </div>
        </>
    )
}