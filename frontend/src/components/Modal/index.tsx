import "./style.css";

export interface IProps {
    active: boolean;
    setActive: (state: boolean) => void;
    children: React.ReactNode;
}

export function Modal(props: IProps) {
    return (
        <>
            <div onClick={() => props.setActive(false)} className={`overlay ${props.active ? "active" : ""}`}></div>
            <div className={`modal ${props.active ? "active" : ""}`}>
                {props.children}
            </div>
        </>
    )
}