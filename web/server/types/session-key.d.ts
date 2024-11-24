import 'express-session'

declare module 'express-session' {
  export interface SessionData {
    access_token: string
    access_token_expires: number
    refresh_token: string
  }
}
