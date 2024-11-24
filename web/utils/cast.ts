export function toNumber(val: any): number | undefined{
    if (typeof val == 'number') {
        return val
    }
    if (typeof val !== "string") {
        return undefined
    }

    const num = parseInt(val, 10)
    if (isNaN(num)) {
        return undefined
    }
    return num
}