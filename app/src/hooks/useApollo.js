import { useMemo } from 'react'
import {
  ApolloClient,
  createHttpLink,
} from '@apollo/client'
import { InMemoryCache } from '@apollo/client/cache'
import { setContext } from '@apollo/client/link/context'

function createApolloClient() {
  const cache = new InMemoryCache({
    typePolicies: {
      Query: {
        fields: {
          listTopicsByAccount: {
            merge: (_, incoming) => incoming,
          },
        },
      },
    },
  })

  const link = createHttpLink({
    uri: 'http://localhost:4040/graphql',
  })

  const authLink = setContext((_, { headers }) => {
    const val = document.cookie.match(`(^|;)\\s*token\\s*=\\s*([^;]+)`)
    const token = val ? val.pop() : undefined
    return {
      headers: {
        ...headers,
        authorization: token ? `Bearer ${token}` : '',
      },
    }
  })

  return new ApolloClient({
    cache,
    link: authLink.concat(link),
    name: 'roshambo',
    version: '1.0.0',
  })
}

let apolloClient

function initializeApollo(initialState = null) {
  const apollo = apolloClient ?? createApolloClient()
  if (initialState) {
    apollo.cache.restore(initialState)
  }
  return apollo
}

export function useApollo(initialState) {
  const apollo = useMemo(() => initializeApollo(initialState), [
    initialState,
  ])
  return apollo
}
