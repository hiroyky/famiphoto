import express from 'express'
import session from 'cookie-session'
import handler from "~/server/interfaces/http/handler";
import {SessionCookieName} from "~/server/config";

const app = express()
const nuxtConfig = useRuntimeConfig()

app.use(session({
    secret: nuxtConfig.sessionSecret,
    name: SessionCookieName,
    path: '/',
    httpOnly: true,
    maxAge: 1000 * 3600 * 24 * 90,
    secure: false,
    sameSite: 'lax',
}))

app.use(handler)

export default fromNodeMiddleware(app)
