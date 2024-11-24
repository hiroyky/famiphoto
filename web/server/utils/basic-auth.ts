import {Base64} from "js-base64";

export function basicAuthValue (userName: string, password: string): string {
    const pass = Base64.encode(`${userName}:${password}`)
    return `Basic ${pass}`
}
