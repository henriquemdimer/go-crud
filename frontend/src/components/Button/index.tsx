import "./style.css";

export interface IProps {
    label: string;
}

export function Button(props: IProps) {
    return (
        <>
            <button className="button">{props.label}</button>
        </>
    )
}