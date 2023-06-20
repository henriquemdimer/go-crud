import Base from "./baseWrapper"

export function getAllTodos(token: string) {
    if (!token) throw new Error("Missing token");

    return Base.get("/todos", {
        headers: {
            Authorization: token
        }
    });
}

export function createTodo(title: string, token: string) {
    if (!title) throw new Error("Missing title");
    if (!token) throw new Error("Missing token");

    return Base.post("/todos", { title }, {
        headers: {
            Authorization: token
        }
    });
}

export function deleteTodo(id: number, token: string) {
    if (!id) throw new Error("Missing id");
    if (!token) throw new Error("Missing token");

    return Base.delete(`/todos/${id}`, {
        headers: {
            Authorization: token
        }
    });
}

export function setTodoDone(id: number, done: boolean, token: string) {
    if (!id) throw new Error("Missing id");
    if (!token) throw new Error("Missing token");

    return Base.put(`/todos/${id}`, { done }, {
        headers: {
            Authorization: token
        }
    });
}