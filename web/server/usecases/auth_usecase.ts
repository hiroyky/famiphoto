import { ApiGateway } from '~/server/drivers/api-gateway'

export class AuthUseCase {
  constructor (
    private api: ApiGateway,
  ) {
  }

  public async login (userId: string, password: string, now: Date) {
    const {
      accessToken,
      refreshToken,
      expireIn,
    } = await this.api.login({ userId, password })
    return {
      accessToken,
      refreshToken,
      expireAt: expireIn + (now.getTime() / 1000),
    }
  }

  public isExpired (expireAt: number, now:Date): boolean {
    return expireAt < now.getTime() / 1000
  }

  public async refreshToken (refreshToken: string, now: Date) {
    const res = await this.api.postOauthToken({
      grantType: 'refresh_token',
      refreshToken,
    })
    return {
      accessToken: res.accessToken,
      expireAt: res.expireIn + (now.getTime() / 1000),
    }
  }
}
