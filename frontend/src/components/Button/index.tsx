import "./style.css";
import React from 'react';

export interface IProps extends React.ButtonHTMLAttributes<HTMLButtonElement> {
    label: string;
    onClick?: (...args: any[]) => any;
    loading?: boolean;
}

export function Button(props: IProps) {
    return (
        <>
            <button type={props.type} className="button" onClick={() => props.onClick?.()}>
                {props.loading ? <img src="loader.svg" width={15} height={15} /> : props.label}
            </button>
        </>
    )
}