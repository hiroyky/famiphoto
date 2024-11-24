import {defineStore} from "pinia";

export  const useApiStore = defineStore('api', {
    actions: {
        async login(req: {  userId: string, password: string}) {
            const headers = useRequestHeaders(['cookie'])
            const res = await fetch('/api/auth/login', {
                method: 'POST',
                body: JSON.stringify({
                    userId: req.userId,
                    password: req.password,
                }),
                headers:{
                    'content-type': 'application/json',
                    ...headers
                }
            })
            if (!res.ok) {
                throw await ApiError.createInstanceFromRes(res)
            }
        },

        async logout() {
            const headers = useRequestHeaders(['cookie'])
            const res = await fetch('/api/auth/logout', {
                method: 'POST',
                headers: {
                    'content-type': 'application/json',
                    ...headers,
                }
            })
            if (!res.ok) {
                throw await ApiError.createInstanceFromRes(res)
            }
        },
        downloadUrl(fileId: string) {
            return `/api/download/files/${fileId}`
        },

        async uploadFile(file: File, uploadURL: string) {
            const fd=new FormData()
            fd.append("file", file)
            const res = await fetch(uploadURL, {
                method: 'POST',
                body: fd,
            })
            if (!res.ok) {
                throw await ApiError.createInstanceFromRes(res)
            }
        }
    }
})


export class ApiError extends Error {
    public httpStatus: number
    public errorCode: string
    public errorMessage: string

    private constructor(data: any) {
        super()
        this.httpStatus = data.httpStatus
        this.errorCode = data.errorCode
        this.errorMessage = data.errorMessage
    }

    public toString(): string {
        return `${this.httpStatus} ${this.errorCode}: ${this.errorMessage} ${super.toString()}`
    }
    static async createInstanceFromRes(res: Response): Promise<ApiError> {
        const body = await res.json()

        return new ApiError({
            errorCode: body && body.error_code,
            errorMessage: body && body.error_message,
            httpStatus: res.status,
        })
    }
}