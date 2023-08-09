import "./style.css";
import React from 'react';

export interface IProps {
    label: string;
    onClick?: (...args: any[]) => any;
    type?: React.ButtonHTMLAttributes<HTMLButtonElement>['type'];
}

export function Button(props: IProps) {
    return (
        <>
            <button type={props.type} className="button" onClick={() => props.onClick?.()}>{props.label}</button>
        </>
    )
}