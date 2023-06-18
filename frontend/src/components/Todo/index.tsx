import { useState } from "react"
import "./style.css"

export function Todo() {
    const [done, setDone] = useState(false)

    return (
        <>
            <li className={`todo ${done ? "line" : ""}`}>
                <div onClick={() => setDone(!done)}>
                    <label className="checkbox">
                        <input type="checkbox" checked={done} />
                        <span className="checkmark" />
                    </label>
                    <p>Terminar de desenvolver este site</p>
                </div>
                <div className='x'></div>
            </li>
        </>
    )
}