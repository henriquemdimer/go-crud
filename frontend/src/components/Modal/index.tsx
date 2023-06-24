import "./style.css";

export interface IProps {
    active: boolean;
    setActive: (state: boolean) => void;
    children: React.ReactNode;
    blockClose?: boolean;
}

export function Modal(props: IProps) {
    return (
        <>
            <div onClick={() => !props.blockClose && props.setActive(false)} className={`overlay ${props.blockClose ? "" : "pointer"} ${props.active ? "active" : ""}`}></div>
            <div className={`modal ${props.active ? "active" : ""}`}>
                {props.children}
            </div>
        </>
    )
}