export class ApiDriver {
  constructor (private baseUrl: string) {
  }

  public async request (path: string, init?: RequestInit) {
    const res = await fetch(this.baseUrl + '/' + path, init)
    if (!res.ok) {
      const body: any = await res.json()
      throw new ApiError(res, body)
    }
    return res
  }
}

export class ApiError extends Error {
  public httpStatus: number
  public errorCode: string
  public errorMessage: string

  constructor (res: Response, body: any) {
    super()
    this.httpStatus = res.status
    this.errorCode = body.error_code
    this.errorMessage = body.error_message
  }

  public toString (): string {
    return `${this.httpStatus} ${this.errorCode}: ${this.errorMessage} ${super.toString()}`
  }
}
