import { ApiDriver } from '../drivers/api-driver'
import { ApiGateway } from '../drivers/api-gateway'
import { AuthUseCase } from '../usecases/auth_usecase'

function newApiDriver () {
    const c = useRuntimeConfig()
    const apiBaseUrl = c.apiBaseUrl
    if (apiBaseUrl === undefined) {
        throw new Error('env: api base url is invalid')
    }
    return new ApiDriver(apiBaseUrl)
}

function newApiGateway () {
    return new ApiGateway(newApiDriver())
}

export function newAuthUseCase (): AuthUseCase {
    return new AuthUseCase(newApiGateway())
}
