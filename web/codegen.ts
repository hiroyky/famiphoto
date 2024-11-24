import type {CodegenConfig} from "@graphql-codegen/cli";

const config: CodegenConfig = {
    schema: 'http://localhost:3000/api/graphql',
    documents: './graphql/*.graphql',
    generates: {
        'types/api-gql.ts': {
            config: {
                scalars: {
                    Timestamp: {
                        input: 'string',
                        output: 'string'
                    }
                }
            },
            plugins: [
                "typescript",
                "typescript-operations",
                "typescript-graphql-request"
            ],
        },
    },
}

export default config