import Base from "./baseWrapper";

export function ping() {
    return Base.get("/");
}