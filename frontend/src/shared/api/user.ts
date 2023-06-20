import Base from './baseWrapper';

export function createUser(name: string, password: string) {
    if (!name || !password) throw new Error("Missing name or password.");
    return Base.post("/users", {
        name,
        password
    });
}

export function Login(name: string, password: string) {
    if (!name || !password) throw new Error("Missing name or password.");
    return Base.post("/users/login", {
        name,
        password
    });
}

export function getUser(token: string) {
    if (!token) throw new Error("Missing token");
    return Base.get("/users", {
        headers: {
            authorization: token
        }
    });
}