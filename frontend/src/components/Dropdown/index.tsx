import "./style.css";

export interface IProps {
    children: React.ReactNode;
    onClick?: (...args: any[]) => any;
}

export interface IDropdownProps extends IProps {
    active: boolean;
    setActive: (value: boolean) => void;
}

export function Dropdown(props: IDropdownProps) {
    return (
        <>
            <div onClick={() => props.setActive(false)} className={`overlay transparent ${props.active ? "active" : ""}`}></div>
            <div className={`dropdown ${props.active ? "active" : ""}`}>
                {props.children}
            </div>
        </>
    )
}

export function DropdownTrigger(props: IProps) {
    return (
        <>
            <div onClick={() => props.onClick?.()} className="dropdown__trigger">
                {props.children}
            </div>
        </>
    )
}

export function DropdownContent(props: IProps) {
    return (
        <>
            <div className="dropdown__content">
                {props.children}
            </div>
        </>
    )
}