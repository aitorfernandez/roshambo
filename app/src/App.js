import React from 'react'
import PropTypes from 'prop-types'
import {
  BrowserRouter as Router,
  Route,
  Switch,
} from 'react-router-dom'
import { ApolloProvider } from '@apollo/client'

import { routes } from './routes'
import { useApollo } from './hooks/useApollo'

function RenderRoute({
  component: Component,
  layout: Layout,
  ...rest
}) {
  return (
    <Route { ...rest } render={(props) => (
      <Layout>
        <Component { ...props } />
      </Layout>
    )} />
  )
}

RenderRoute.propTypes = {
  component: PropTypes.func.isRequired,
  layout: PropTypes.func.isRequired,
}

function RouteWithSubRoutes(route) {
  return (
    <RenderRoute { ...route } />
  )
}

export function App() {
  return (
    <ApolloProvider { ...{ client: useApollo() } }>
      <Router>
        <Switch>
        {
          routes.map(({ ...props }, i) => (
            <RouteWithSubRoutes key={i} { ...props } />
          ))
        }
        </Switch>
      </Router>
    </ApolloProvider>
  )
}
